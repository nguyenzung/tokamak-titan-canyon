import argparse
import logging
import os
import subprocess
import json
import socket
import datetime
import time
import shutil
import http.client
import gzip
import glob
from multiprocessing import Process, Queue
import concurrent.futures
from collections import namedtuple

import devnet.log_setup

pjoin = os.path.join

parser = argparse.ArgumentParser(description='Bedrock devnet launcher')
parser.add_argument('--monorepo-dir', help='Directory of the monorepo', default=os.getcwd())
parser.add_argument('--allocs', help='Only create the allocs and exit', type=bool, action=argparse.BooleanOptionalAction)
parser.add_argument('--test', help='Tests the deployment, must already be deployed', type=bool, action=argparse.BooleanOptionalAction)
parser.add_argument('--fork-public-network',
                    help='Fork the public network',
                    type=bool,
                    default=os.environ.get('FORK_PUBLIC_NETWORK').lower() == 'true' if os.environ.get('FORK_PUBLIC_NETWORK') else False)
parser.add_argument('--l1-rpc-url', help='Public L1 RPC URL', type=str, default=os.environ.get('L1_RPC'))
parser.add_argument('--from-block-number', help='From block number', type=int, default=os.environ.get('FROM_BLOCK_NUMBER'))
parser.add_argument('--l2-native-token', help='L2 native token', type=str, default=os.environ.get('L2_NATIVE_TOKEN'))
parser.add_argument('--legacy', help='Legacy(using eth) or not(apply native token)', type=bool, action=argparse.BooleanOptionalAction)
parser.add_argument('--admin-key', help='The admin private key for upgrade contracts', type=str, default=os.environ.get('DEVNET_ADMIN_PRIVATE_KEY'))

log = logging.getLogger()

class Bunch:
    def __init__(self, **kwds):
        self.__dict__.update(kwds)

class ChildProcess:
    def __init__(self, func, *args):
        self.errq = Queue()
        self.process = Process(target=self._func, args=(func, args))

    def _func(self, func, args):
        try:
            func(*args)
        except Exception as e:
            self.errq.put(str(e))

    def start(self):
        self.process.start()

    def join(self):
        self.process.join()

    def get_error(self):
        return self.errq.get() if not self.errq.empty() else None


def main():
    args = parser.parse_args()

    validate_fork_public_network(args)

    monorepo_dir = os.path.abspath(args.monorepo_dir)
    devnet_dir = pjoin(monorepo_dir, '.devnet')
    contracts_bedrock_dir = pjoin(monorepo_dir, 'packages', 'tokamak', 'contracts-bedrock') if not args.legacy else pjoin(monorepo_dir, 'packages', 'contracts-bedrock')
    deployment_dir = pjoin(contracts_bedrock_dir, 'deployments', 'devnetL1')
    op_node_dir = pjoin(args.monorepo_dir, 'op-node')
    ops_bedrock_dir = pjoin(monorepo_dir, 'ops-bedrock')
    deploy_config_dir = pjoin(contracts_bedrock_dir, 'deploy-config')
    devnet_config_path = pjoin(deploy_config_dir, 'devnetL1.json')
    devnet_config_template_path = pjoin(deploy_config_dir, 'fork.devnetL1-template.json') if args.fork_public_network else pjoin(deploy_config_dir, 'devnetL1-template.json')
    ops_chain_ops = pjoin(monorepo_dir, 'op-chain-ops')
    sdk_dir = pjoin(monorepo_dir, 'packages', 'tokamak', 'sdk') if not args.legacy else pjoin(monorepo_dir, 'packages', 'sdk')
    bedrock_devnet_dir = pjoin(monorepo_dir, 'bedrock-devnet')

    paths = Bunch(
      mono_repo_dir=monorepo_dir,
      devnet_dir=devnet_dir,
      contracts_bedrock_dir=contracts_bedrock_dir,
      deployment_dir=deployment_dir,
      l1_deployments_path=pjoin(deployment_dir, '.deploy'),
      deploy_config_dir=deploy_config_dir,
      devnet_config_path=devnet_config_path,
      devnet_config_template_path=devnet_config_template_path,
      op_node_dir=op_node_dir,
      ops_bedrock_dir=ops_bedrock_dir,
      ops_chain_ops=ops_chain_ops,
      sdk_dir=sdk_dir,
      genesis_l1_path=pjoin(devnet_dir, 'genesis-l1.json'),
      genesis_l2_path=pjoin(devnet_dir, 'genesis-l2.json'),
      allocs_path=pjoin(devnet_dir, 'allocs-l1.json'),
      addresses_json_path=pjoin(devnet_dir, 'addresses.json'),
      sdk_addresses_json_path=pjoin(devnet_dir, 'sdk-addresses.json'),
      rollup_config_path=pjoin(devnet_dir, 'rollup.json'),
      fork_public_network = args.fork_public_network,
      l1_rpc_url = args.l1_rpc_url,
      l2_native_token = args.l2_native_token,
      from_block_number = args.from_block_number,
      bedrock_devnet_path=bedrock_devnet_dir,
      admin_key=args.admin_key
    )

    if args.test:
      log.info('Testing deployed devnet')
      devnet_test(paths)
      return

    os.makedirs(devnet_dir, exist_ok=True)

    if args.allocs:
      devnet_l1_genesis(paths)
      return

    git_commit = subprocess.run(['git', 'rev-parse', 'HEAD'], capture_output=True, text=True).stdout.strip()
    git_date = subprocess.run(['git', 'show', '-s', "--format=%ct"], capture_output=True, text=True).stdout.strip()

    # CI loads the images from workspace, and does not otherwise know the images are good as-is
    if os.getenv('DEVNET_NO_BUILD') == "true":
        log.info('Skipping docker images build')
    else:
        log.info(f'Building docker images for git commit {git_commit} ({git_date})')
        if args.fork_public_network:
          run_command(['docker', 'compose', '-f', 'fork.docker-compose.yml', 'build', '--progress', 'plain',
                     '--build-arg', f'GIT_COMMIT={git_commit}', '--build-arg', f'GIT_DATE={git_date}'],
                    cwd=paths.ops_bedrock_dir, env={
            'PWD': paths.ops_bedrock_dir,
            'DOCKER_BUILDKIT': '1', # (should be available by default in later versions, but explicitly enable it anyway)
            'COMPOSE_DOCKER_CLI_BUILD': '1',  # use the docker cache
            'L1_RPC': paths.l1_rpc_url,
            'FROM_BLOCK_NUMBER': str(paths.from_block_number)
          })
        else:
          run_command(['docker', 'compose', 'build', '--progress', 'plain',
                     '--build-arg', f'GIT_COMMIT={git_commit}', '--build-arg', f'GIT_DATE={git_date}'],
                    cwd=paths.ops_bedrock_dir, env={
            'PWD': paths.ops_bedrock_dir,
            'DOCKER_BUILDKIT': '1', # (should be available by default in later versions, but explicitly enable it anyway)
            'COMPOSE_DOCKER_CLI_BUILD': '1'  # use the docker cache
          })

    log.info('Devnet starting')
    devnet_deploy(paths)

def deploy_contracts(paths):
    wait_up(8545)
    wait_for_rpc_server('127.0.0.1:8545')
    res = eth_accounts('127.0.0.1:8545')

    response = json.loads(res)
    account = response['result'][0 if paths.admin_key is None else 1]
    log.info(f'Deploying with {account}')

    # Proxy exists on the fork public network(used by anvil)
    # We don't need to deploy the proxy contract
    # https://book.getfoundry.sh/tutorials/create2-tutorial
    if not paths.fork_public_network:
        # send some ether to the create2 deployer account
        cmd = [
            'cast', 'send',
            '--rpc-url', 'http://127.0.0.1:8545',
            '--value', '1ether', '0x3fAB184622Dc19b6109349B94811493BF2a45362'
        ]
        cmd.extend(['--from', account, '--unlocked']) if paths.admin_key is None else cmd.extend(['--password', '1234'])
        run_command(cmd, env={}, cwd=paths.contracts_bedrock_dir)

        # deploy the create2 deployer
        run_command([
          'cast', 'publish', '--rpc-url', 'http://127.0.0.1:8545',
          '0xf8a58085174876e800830186a08080b853604580600e600039806000f350fe7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe03601600081602082378035828234f58015156039578182fd5b8082525050506014600cf31ba02222222222222222222222222222222222222222222222222222222222222222a02222222222222222222222222222222222222222222222222222222222222222'
        ], env={}, cwd=paths.contracts_bedrock_dir)

    fqn = 'scripts/Deploy.s.sol:Deploy'
    cmd = [
        'forge', 'script', fqn,
        '--rpc-url', 'http://127.0.0.1:8545', '--broadcast'
    ]
    cmd.extend(['--sender', account, '--unlocked']) if paths.admin_key is None else cmd.extend(['--private-key', paths.admin_key])
    run_command(cmd, env={}, cwd=paths.contracts_bedrock_dir)

    shutil.copy(paths.l1_deployments_path, paths.addresses_json_path)

    log.info('Syncing contracts.')
    run_command([
        'forge', 'script', fqn, '--sig', 'sync()',
        '--memory-limit', '268435456',
        '--rpc-url', 'http://127.0.0.1:8545'
    ], env={}, cwd=paths.contracts_bedrock_dir)

def init_devnet_l1_deploy_config(paths, update_timestamp=False, temp=True):
    deploy_config = read_json(paths.devnet_config_template_path) if temp else read_json(paths.devnet_config_path)
    if update_timestamp:
        deploy_config['l1GenesisBlockTimestamp'] = '{:#x}'.format(int(time.time()))
    write_json(paths.devnet_config_path, deploy_config)

def init_admin_geth(paths):
    deploy_config = read_json(paths.devnet_config_template_path)
    admin_address = deploy_config['finalSystemOwner']

    f = open(pjoin(paths.bedrock_devnet_path, 'genesis.json'), "w+")
    run_command([
        'geth', '--dev', 'dumpgenesis'
    ], cwd=paths.bedrock_devnet_path, stdout=f)
    f.close()

    genesis = read_json(pjoin(paths.bedrock_devnet_path, 'genesis.json'))

    genesis["config"]["chainId"] = 900

    alloc = genesis['alloc']
    alloc[admin_address] = {
        "balance": "10000000000000000000"
    }
    genesis['alloc'] = alloc

    write_file(pjoin(paths.bedrock_devnet_path, 'keystore'), paths.admin_key[2:])
    write_file(pjoin(paths.bedrock_devnet_path, 'password'), '1234')

    write_json(pjoin(paths.bedrock_devnet_path, 'genesis.json'), genesis)
    geth_init(paths)
    os.environ.setdefault(
        key = 'GETH_DATADIR',
        value = 'data',
    )

    key_store = pjoin(paths.bedrock_devnet_path, 'data', 'keystore')
    key_file = glob.glob(f"{key_store}/*{admin_address[2:].lower()}*")
    os.environ.setdefault(
        key = 'ETH_KEYSTORE',
        value = pjoin(paths.bedrock_devnet_path, 'data', 'keystore', key_file[0])
    )


def geth_init(paths):
    run_command([
        'geth', '--dev', '--datadir', 'data', 'init', 'genesis.json'
    ], cwd=paths.bedrock_devnet_path)

    run_command([
        'geth', '--datadir', 'data', '--password', 'password', 'account', 'import', 'keystore'
    ], cwd=paths.bedrock_devnet_path)
    delete_file(pjoin(paths.bedrock_devnet_path, 'keystore'))
    delete_file(pjoin(paths.bedrock_devnet_path, 'password'))
    delete_file(pjoin(paths.bedrock_devnet_path, 'genesis.json'))

def devnet_l1_genesis(paths):
    log.info('Generating L1 genesis state')
    init_devnet_l1_deploy_config(paths)

    if paths.fork_public_network:
        log.info('Start to fork the public network. Wait to warm up the fork public network.')
        geth = subprocess.Popen([
            'anvil', '--fork-url', paths.l1_rpc_url, '--fork-block-number', str(paths.from_block_number),
            '--chain-id', '900', '--disable-block-gas-limit',
            '--gas-price', '0', '--base-fee', '1', '--block-time', '1'
        ])
        time.sleep(30)
    else:
        if paths.admin_key is not None and not os.path.exists(pjoin(paths.bedrock_devnet_path, 'data')) :
            init_admin_geth(paths)

        geth = subprocess.Popen([
            'geth', '--dev', '--dev.period', '2', '--http', '--http.api', 'eth,debug',
            '--verbosity', '4', '--gcmode', 'archive', '--dev.gaslimit', '30000000',
            '--rpc.allow-unprotected-txs'
        ], cwd=pjoin(paths.mono_repo_dir, 'bedrock-devnet'))

    try:
        forge = ChildProcess(deploy_contracts, paths)
        forge.start()
        forge.join()
        err = forge.get_error()
        if err:
            raise Exception(f"Exception occurred in child process: {err}")

        if paths.fork_public_network:
          res = anvil_dumpState('127.0.0.1:8545')
          allocs = convert_anvil_dump(res)
        else:
          res = debug_dumpBlock('127.0.0.1:8545')
          response = json.loads(res)
          allocs = response['result']
        write_json(paths.allocs_path, allocs)
    finally:
        geth.terminate()


# Bring up the devnet where the contracts are deployed to L1
def devnet_deploy(paths):
    if os.path.exists(paths.genesis_l1_path):
        log.info('L1 genesis already generated.')
    else:
        log.info('Generating L1 genesis.')
        if os.path.exists(paths.allocs_path) == False:
            devnet_l1_genesis(paths)

        # It's odd that we want to regenerate the devnetL1.json file with
        # an updated timestamp different than the one used in the devnet_l1_genesis
        # function.  But, without it, CI flakes on this test rather consistently.
        # If someone reads this comment and understands why this is being done, please
        # update this comment to explain.
        init_devnet_l1_deploy_config(paths, update_timestamp=True, temp=False)
        outfile_l1 = pjoin(paths.devnet_dir, 'genesis-l1.json')
        run_command([
            'go', 'run', 'cmd/main.go', 'genesis', 'l1',
            '--deploy-config', paths.devnet_config_path,
            '--l1-allocs', paths.allocs_path,
            '--l1-deployments', paths.addresses_json_path,
            '--outfile.l1', outfile_l1,
        ], cwd=paths.op_node_dir)

    log.info('Starting L1.')
    if paths.fork_public_network:
      run_command(['docker', 'compose', '-f', 'fork.docker-compose.yml', 'up', '-d', 'fork-l1'], cwd=paths.ops_bedrock_dir, env={
        'PWD': paths.ops_bedrock_dir
      })
    else:
      run_command(['docker', 'compose', '-f', 'docker-compose.yml', 'up', '-d', 'l1'], cwd=paths.ops_bedrock_dir, env={
        'PWD': paths.ops_bedrock_dir
      })
    wait_up(8545)
    wait_for_rpc_server('127.0.0.1:8545')

    if os.path.exists(paths.genesis_l2_path):
        log.info('L2 genesis and rollup configs already generated.')
    else:
        log.info('Generating L2 genesis and rollup configs.')
        run_command([
            'go', 'run', 'cmd/main.go', 'genesis', 'l2',
            '--l1-rpc', 'http://localhost:8545',
            '--deploy-config', paths.devnet_config_path,
            '--deployment-dir', paths.deployment_dir,
            '--outfile.l2', pjoin(paths.devnet_dir, 'genesis-l2.json'),
            '--outfile.rollup', pjoin(paths.devnet_dir, 'rollup.json')
        ], cwd=paths.op_node_dir)

    rollup_config = read_json(paths.rollup_config_path)
    addresses = read_json(paths.addresses_json_path)

    log.info('Bringing up L2.')
    run_command(['docker', 'compose', 'up', '-d', 'l2'], cwd=paths.ops_bedrock_dir, env={
        'PWD': paths.ops_bedrock_dir
    })
    wait_up(9545)
    wait_for_rpc_server('127.0.0.1:9545')

    l2_output_oracle = addresses['L2OutputOracleProxy']
    log.info(f'Using L2OutputOracle {l2_output_oracle}')
    batch_inbox_address = rollup_config['batch_inbox_address']
    log.info(f'Using batch inbox {batch_inbox_address}')

    log.info('Bringing up `op-node`, `op-proposer` and `op-batcher`.')
    run_command(['docker', 'compose', 'up', '-d', 'op-node', 'op-proposer', 'op-batcher'], cwd=paths.ops_bedrock_dir, env={
        'PWD': paths.ops_bedrock_dir,
        'L2OO_ADDRESS': l2_output_oracle,
        'SEQUENCER_BATCH_INBOX_ADDRESS': batch_inbox_address
    })

    log.info('Bringing up `artifact-server`')
    run_command(['docker', 'compose', 'up', '-d', 'artifact-server'], cwd=paths.ops_bedrock_dir, env={
        'PWD': paths.ops_bedrock_dir
    })

    shutil.rmtree(pjoin(paths.bedrock_devnet_path, 'data'), ignore_errors=True)

    log.info('Devnet ready.')


def eth_accounts(url):
    log.info(f'Fetch eth_accounts {url}')
    conn = http.client.HTTPConnection(url)
    headers = {'Content-type': 'application/json'}
    body = '{"id":2, "jsonrpc":"2.0", "method": "eth_accounts", "params":[]}'
    conn.request('POST', '/', body, headers)
    response = conn.getresponse()
    data = response.read().decode()
    conn.close()
    return data

def debug_dumpBlock(url):
    log.info(f'Fetch debug_dumpBlock {url}')
    conn = http.client.HTTPConnection(url)
    headers = {'Content-type': 'application/json'}
    body = '{"id":3, "jsonrpc":"2.0", "method": "debug_dumpBlock", "params":["latest"]}'
    conn.request('POST', '/', body, headers)
    response = conn.getresponse()
    data = response.read().decode()
    conn.close()
    return data


def wait_for_rpc_server(url):
    log.info(f'Waiting for RPC server at {url}')

    conn = http.client.HTTPConnection(url)
    headers = {'Content-type': 'application/json'}
    body = '{"id":1, "jsonrpc":"2.0", "method": "eth_chainId", "params":[]}'

    while True:
        try:
            conn.request('POST', '/', body, headers)
            response = conn.getresponse()
            conn.close()
            if response.status < 300:
                log.info(f'RPC server at {url} ready')
                return
        except Exception as e:
            log.info(f'Waiting for RPC server at {url}')
            time.sleep(1)


CommandPreset = namedtuple('Command', ['name', 'args', 'cwd', 'timeout'])


def devnet_test(paths):
    # Check the L2 config
    run_command(
        ['go', 'run', 'cmd/check-l2/main.go', '--l2-rpc-url', 'http://localhost:9545', '--l1-rpc-url', 'http://localhost:8545'],
        cwd=paths.ops_chain_ops,
    )

    # Run the two commands with different signers, so the ethereum nonce management does not conflict
    # And do not use devnet system addresses, to avoid breaking fee-estimation or nonce values.
    run_commands([
        CommandPreset('erc20-test',
          ['npx', 'hardhat',  'deposit-erc20', '--network',  'devnetL1',
           '--l1-contracts-json-path', paths.addresses_json_path, '--signer-index', '14'],
          cwd=paths.sdk_dir, timeout=8*60),
        CommandPreset('eth-test',
          ['npx', 'hardhat',  'deposit-eth', '--network',  'devnetL1',
           '--l1-contracts-json-path', paths.addresses_json_path, '--signer-index', '15'],
          cwd=paths.sdk_dir, timeout=8*60)
    ], max_workers=2)


def run_commands(commands: list[CommandPreset], max_workers=2):
    with concurrent.futures.ThreadPoolExecutor(max_workers=max_workers) as executor:
        futures = [executor.submit(run_command_preset, cmd) for cmd in commands]

        for future in concurrent.futures.as_completed(futures):
            result = future.result()
            if result:
                print(result.stdout)


def run_command_preset(command: CommandPreset):
    with subprocess.Popen(command.args, cwd=command.cwd,
                          stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True) as proc:
        try:
            # Live output processing
            for line in proc.stdout:
                # Annotate and print the line with timestamp and command name
                timestamp = datetime.datetime.utcnow().strftime('%H:%M:%S.%f')
                # Annotate and print the line with the timestamp
                print(f"[{timestamp}][{command.name}] {line}", end='')

            stdout, stderr = proc.communicate(timeout=command.timeout)

            if proc.returncode != 0:
                raise RuntimeError(f"Command '{' '.join(command.args)}' failed with return code {proc.returncode}: {stderr}")

        except subprocess.TimeoutExpired:
            raise RuntimeError(f"Command '{' '.join(command.args)}' timed out!")

        except Exception as e:
            raise RuntimeError(f"Error executing '{' '.join(command.args)}': {e}")

        finally:
            # Ensure process is terminated
            proc.kill()
    return proc.returncode


def run_command(args, check=True, shell=False, cwd=None, env=None, timeout=None, stdout=None):
    env = env if env else {}
    return subprocess.run(
        args,
        check=check,
        shell=shell,
        env={
            **os.environ,
            **env
        },
        cwd=cwd,
        timeout=timeout,
        stdout=stdout
    )


def wait_up(port, retries=10, wait_secs=1):
    for i in range(0, retries):
        log.info(f'Trying 127.0.0.1:{port}')
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        try:
            s.connect(('127.0.0.1', int(port)))
            s.shutdown(2)
            log.info(f'Connected 127.0.0.1:{port}')
            return True
        except Exception:
            time.sleep(wait_secs)

    raise Exception(f'Timed out waiting for port {port}.')


def write_json(path, data):
    with open(path, 'w+') as f:
        json.dump(data, f, indent='  ')


def read_json(path):
    with open(path, 'r') as f:
        return json.load(f)


def validate_fork_public_network(args):
    fork_public_network = args.fork_public_network
    l1_rpc_url = args.l1_rpc_url
    l2_native_token = args.l2_native_token
    from_block_number = args.from_block_number
    # If fork the public network, validate the required params related to
    if fork_public_network:
      if not l1_rpc_url:
        raise Exception("Please provide the L1_RPC URL for the forked network.")

      if not l2_native_token:
        raise Exception("Please provide the L2_NATIVE_TOKEN for the forked network.")


      if not from_block_number:
        raise Exception("Please provide the FROM_BLOCK_NUMBER for the forked network.")

      if from_block_number <= 0:
        raise Exception("Please provide the FROM_BLOCK_NUMBER is bigger than zero.")
      log.info(f'Fork from RPC URL: {l1_rpc_url}, from block number: {from_block_number}, l2 native token: {l2_native_token}')


def anvil_dumpState(url):
    log.info(f'Fetch debug_dumpBlock {url}')
    conn = http.client.HTTPConnection(url)
    headers = {'Content-type': 'application/json'}
    body = '{"id":3, "jsonrpc":"2.0", "method": "anvil_dumpState", "params":[]}'
    conn.request('POST', '/', body, headers)
    data = conn.getresponse().read()
    # Anvil returns a JSON-RPC response with a hex-encoded "result" field
    result = json.loads(data.decode('utf-8'))['result']
    result_bytes = bytes.fromhex(result[2:])
    uncompressed = gzip.decompress(result_bytes).decode()
    return json.loads(uncompressed)

def convert_anvil_dump(dump):
    accounts = dump['accounts']

    for account in accounts.values():
      bal = account['balance']
      account['balance'] = str(int(bal, 16))

      if 'storage' in account:
        storage = account['storage']
        storage_keys = list(storage.keys())
        for key in storage_keys:
          value = storage[key]
          del storage[key]
          storage[pad_hex(key)] = pad_hex(value)
    return dump

def pad_hex(input):
    return '0x' + input.replace('0x', '').zfill(64)

def write_file(path, data):
    f = open(path, 'w+')
    f.write(data)
    f.close()

def delete_file(path):
    os.remove(path)
