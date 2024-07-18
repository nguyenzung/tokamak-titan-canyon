// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/tokamak-network/tokamak-thanos/op-bindings/solc"
)

const L2UsdcBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/usdc-bridge/L2/tokamak-UsdcBridge/L2UsdcBridge.sol:L2UsdcBridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/usdc-bridge/L2/tokamak-UsdcBridge/L2UsdcBridge.sol:L2UsdcBridge\",\"label\":\"otherBridge\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1002,\"contract\":\"src/usdc-bridge/L2/tokamak-UsdcBridge/L2UsdcBridge.sol:L2UsdcBridge\",\"label\":\"l1Usdc\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"src/usdc-bridge/L2/tokamak-UsdcBridge/L2UsdcBridge.sol:L2UsdcBridge\",\"label\":\"l2Usdc\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/usdc-bridge/L2/tokamak-UsdcBridge/L2UsdcBridge.sol:L2UsdcBridge\",\"label\":\"l2UsdcMasterMinter\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"}}}"

var L2UsdcBridgeStorageLayout = new(solc.StorageLayout)

var L2UsdcBridgeDeployedBin = "0x6080604052600436106100705760003560e01c806356c3b5871161004e57806356c3b5871461010d578063662a633a1461013a578063a1b4bc041461015a578063c89701a21461018757600080fd5b806305db940f1461007557806332b7006d146100cb5780633cb747bf146100e0575b600080fd5b34801561008157600080fd5b506004546100a29073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b6100de6100d9366004610ebf565b6101b4565b005b3480156100ec57600080fd5b506000546100a29073ffffffffffffffffffffffffffffffffffffffff1681565b34801561011957600080fd5b506002546100a29073ffffffffffffffffffffffffffffffffffffffff1681565b34801561014657600080fd5b506100de610155366004610f3b565b61025e565b34801561016657600080fd5b506003546100a29073ffffffffffffffffffffffffffffffffffffffff1681565b34801561019357600080fd5b506001546100a29073ffffffffffffffffffffffffffffffffffffffff1681565b333b15610248576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b61025785333387878787610759565b5050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331480156103335750600154600054604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9384169390921691636e296e45916004808201926020929091908290030181865afa1580156102f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031b9190610fd3565b73ffffffffffffffffffffffffffffffffffffffff16145b6103e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a40161023f565b600254879073ffffffffffffffffffffffffffffffffffffffff80831691161461046b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f74204c312075736463000000000000000000000000000000000000000000604482015260640161023f565b600354879073ffffffffffffffffffffffffffffffffffffffff8083169116146104f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f74204c322075736463000000000000000000000000000000000000000000604482015260640161023f565b6040517f8a6db9c300000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8a1690638a6db9c390602401602060405180830381865afa15801561055e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105829190610ff0565b90508581101561064357600480546040517fcbf2b8bf0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9281019290925273ffffffffffffffffffffffffffffffffffffffff169063cbf2b8bf906024016020604051808303816000875af115801561061d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106419190611009565b505b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152602482018890528a16906340c10f1990604401600060405180830381600087803b1580156106b357600080fd5b505af11580156106c7573d6000803e3d6000fd5b505050508773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd898a8a8a8a6040516107459493929190611074565b60405180910390a450505050505050505050565b600354879073ffffffffffffffffffffffffffffffffffffffff8083169116146107df576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f74204c322075736463000000000000000000000000000000000000000000604482015260640161023f565b6040517f8a6db9c300000000000000000000000000000000000000000000000000000000815230600482015260009073ffffffffffffffffffffffffffffffffffffffff8a1690638a6db9c390602401602060405180830381865afa15801561084c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108709190610ff0565b90508581101561093157600480546040517fcbf2b8bf0000000000000000000000000000000000000000000000000000000081527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9281019290925273ffffffffffffffffffffffffffffffffffffffff169063cbf2b8bf906024016020604051808303816000875af115801561090b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061092f9190611009565b505b61095373ffffffffffffffffffffffffffffffffffffffff8a16893089610b72565b6040517f42966c680000000000000000000000000000000000000000000000000000000081526004810187905273ffffffffffffffffffffffffffffffffffffffff8a16906342966c6890602401600060405180830381600087803b1580156109bb57600080fd5b505af11580156109cf573d6000803e3d6000fd5b505060005460015460025460405173ffffffffffffffffffffffffffffffffffffffff9384169550633dbb202b9450918316927fa9f9e6750000000000000000000000000000000000000000000000000000000092610a4092909116908f908f908f908f908e908e906024016110aa565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b9092168252610ad392918a9060040161112b565b600060405180830381600087803b158015610aed57600080fd5b505af1158015610b01573d6000803e3d6000fd5b505060025460405173ffffffffffffffffffffffffffffffffffffffff8c811694508d81169350909116907f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e90610b5f908c908c908b908b90611074565b60405180910390a4505050505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052610c07908590610c0d565b50505050565b6000610c2f73ffffffffffffffffffffffffffffffffffffffff841683610ca8565b90508051600014158015610c54575080806020019051810190610c529190611009565b155b15610ca3576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161023f565b505050565b6060610cb683836000610cbd565b9392505050565b606081471015610cfb576040517fcd78605900000000000000000000000000000000000000000000000000000000815230600482015260240161023f565b6000808573ffffffffffffffffffffffffffffffffffffffff168486604051610d2491906111a7565b60006040518083038185875af1925050503d8060008114610d61576040519150601f19603f3d011682016040523d82523d6000602084013e610d66565b606091505b5091509150610d76868383610d80565b9695505050505050565b606082610d9557610d9082610e0f565b610cb6565b8151158015610db9575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610e08576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161023f565b5080610cb6565b805115610e1f5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b73ffffffffffffffffffffffffffffffffffffffff81168114610e5157600080fd5b60008083601f840112610e8857600080fd5b50813567ffffffffffffffff811115610ea057600080fd5b602083019150836020828501011115610eb857600080fd5b9250929050565b600080600080600060808688031215610ed757600080fd5b8535610ee281610e54565b945060208601359350604086013563ffffffff81168114610f0257600080fd5b9250606086013567ffffffffffffffff811115610f1e57600080fd5b610f2a88828901610e76565b969995985093965092949392505050565b600080600080600080600060c0888a031215610f5657600080fd5b8735610f6181610e54565b96506020880135610f7181610e54565b95506040880135610f8181610e54565b94506060880135610f9181610e54565b93506080880135925060a088013567ffffffffffffffff811115610fb457600080fd5b610fc08a828b01610e76565b989b979a50959850939692959293505050565b600060208284031215610fe557600080fd5b8151610cb681610e54565b60006020828403121561100257600080fd5b5051919050565b60006020828403121561101b57600080fd5b81518015158114610cb657600080fd5b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152606060408201526000610d7660608301848661102b565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a08301526110fa60c08301848661102b565b9998505050505050505050565b60005b8381101561112257818101518382015260200161110a565b50506000910152565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260008351806060840152611166816080850160208801611107565b63ffffffff93909316604083015250601f919091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160160800192915050565b600082516111b9818460208701611107565b919091019291505056fea164736f6c6343000814000a"


func init() {
	if err := json.Unmarshal([]byte(L2UsdcBridgeStorageLayoutJSON), L2UsdcBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2UsdcBridge"] = L2UsdcBridgeStorageLayout
	deployedBytecodes["L2UsdcBridge"] = L2UsdcBridgeDeployedBin
	immutableReferences["L2UsdcBridge"] = false
}
