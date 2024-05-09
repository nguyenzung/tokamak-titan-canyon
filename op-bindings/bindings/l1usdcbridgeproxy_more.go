// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1UsdcBridgeProxyStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/tokamak-contracts/USDC/L1/tokamak-UsdcBridge/L1UsdcBridgeProxy.sol:L1UsdcBridgeProxy\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/tokamak-contracts/USDC/L1/tokamak-UsdcBridge/L1UsdcBridgeProxy.sol:L1UsdcBridgeProxy\",\"label\":\"otherBridge\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1002,\"contract\":\"src/tokamak-contracts/USDC/L1/tokamak-UsdcBridge/L1UsdcBridgeProxy.sol:L1UsdcBridgeProxy\",\"label\":\"l1Usdc\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"src/tokamak-contracts/USDC/L1/tokamak-UsdcBridge/L1UsdcBridgeProxy.sol:L1UsdcBridgeProxy\",\"label\":\"l2Usdc\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/tokamak-contracts/USDC/L1/tokamak-UsdcBridge/L1UsdcBridgeProxy.sol:L1UsdcBridgeProxy\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L1UsdcBridgeProxyStorageLayout = new(solc.StorageLayout)

var L1UsdcBridgeProxyDeployedBin = "0x6080604052600436106100c05760003560e01c80638da5cb5b11610074578063a1b4bc041161004e578063a1b4bc041461028a578063c89701a2146102b7578063dfd3dcb3146102e45761012c565b80638da5cb5b1461020f5780638f601f66146102245780639608088c1461026a5761012c565b80634f1ef286116100a55780634f1ef286146101ad57806356c3b587146101cd5780635c60da1b146101fa5761012c565b80633659cfe6146101365780633cb747bf146101565761012c565b3661012c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f63616e6e6f74207265636569766520457468657200000000000000000000000060448201526064015b60405180910390fd5b610134610304565b005b34801561014257600080fd5b50610134610151366004610ca6565b610316565b34801561016257600080fd5b506000546101839073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101b957600080fd5b506101346101c8366004610cf0565b6103ce565b3480156101d957600080fd5b506002546101839073ffffffffffffffffffffffffffffffffffffffff1681565b34801561020657600080fd5b50610183610478565b34801561021b57600080fd5b50610183610487565b34801561023057600080fd5b5061025c61023f366004610dd0565b600460209081526000928352604080842090915290825290205481565b6040519081526020016101a4565b34801561027657600080fd5b50610134610285366004610e03565b610491565b34801561029657600080fd5b506003546101839073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102c357600080fd5b506001546101839073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102f057600080fd5b506101346102ff366004610ca6565b61079e565b61031461030f610992565b61099c565b565b61031e610487565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f6e6f74206f776e657200000000000000000000000000000000000000000000006044820152606401610123565b6103cb8160405180602001604052806000815250610843565b50565b6103d6610487565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461046a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f6e6f74206f776e657200000000000000000000000000000000000000000000006044820152606401610123565b6104748282610843565b5050565b6000610482610992565b905090565b60006104826109c0565b610499610487565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461052d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f6e6f74206f776e657200000000000000000000000000000000000000000000006044820152606401610123565b8373ffffffffffffffffffffffffffffffffffffffff81166105ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f7a65726f206164647265737300000000000000000000000000000000000000006044820152606401610123565b8373ffffffffffffffffffffffffffffffffffffffff8116610629576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f7a65726f206164647265737300000000000000000000000000000000000000006044820152606401610123565b8373ffffffffffffffffffffffffffffffffffffffff81166106a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f7a65726f206164647265737300000000000000000000000000000000000000006044820152606401610123565b8373ffffffffffffffffffffffffffffffffffffffff8116610725576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f7a65726f206164647265737300000000000000000000000000000000000000006044820152606401610123565b50506000805473ffffffffffffffffffffffffffffffffffffffff9788167fffffffffffffffffffffffff00000000000000000000000000000000000000009182161790915560018054968816968216969096179095555050600280549285169284169290921790915560038054919093169116179055565b6107a6610487565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461083a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f6e6f74206f776e657200000000000000000000000000000000000000000000006044820152606401610123565b6103cb816108ab565b61084c82610a00565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a28051156108a35761089e828261090c565b505050565b610474610ad2565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6108d46109c0565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a16103cb81610b0a565b60606000808473ffffffffffffffffffffffffffffffffffffffff16846040516109369190610e57565b600060405180830381855af49150503d8060008114610971576040519150601f19603f3d011682016040523d82523d6000602084013e610976565b606091505b5091509150610986858383610b81565b95945050505050565b90565b6000610482610c13565b3660008037600080366000845af43d6000803e8080156109bb573d6000f35b3d6000fd5b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b8073ffffffffffffffffffffffffffffffffffffffff163b600003610a69576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610123565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b3415610314576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b5a576040517f62e77ba200000000000000000000000000000000000000000000000000000000815260006004820152602401610123565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103610a8c565b606082610b9657610b9182610c3b565b610c0c565b8151158015610bba575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610c09576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610123565b50805b9392505050565b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc6109e4565b805115610c4b5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b803573ffffffffffffffffffffffffffffffffffffffff81168114610ca157600080fd5b919050565b600060208284031215610cb857600080fd5b610c0c82610c7d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060408385031215610d0357600080fd5b610d0c83610c7d565b9150602083013567ffffffffffffffff80821115610d2957600080fd5b818501915085601f830112610d3d57600080fd5b813581811115610d4f57610d4f610cc1565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610d9557610d95610cc1565b81604052828152886020848701011115610dae57600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b60008060408385031215610de357600080fd5b610dec83610c7d565b9150610dfa60208401610c7d565b90509250929050565b60008060008060808587031215610e1957600080fd5b610e2285610c7d565b9350610e3060208601610c7d565b9250610e3e60408601610c7d565b9150610e4c60608601610c7d565b905092959194509250565b6000825160005b81811015610e785760208186018101518583015201610e5e565b81811115610e87576000828501525b50919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1UsdcBridgeProxyStorageLayoutJSON), L1UsdcBridgeProxyStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1UsdcBridgeProxy"] = L1UsdcBridgeProxyStorageLayout
	deployedBytecodes["L1UsdcBridgeProxy"] = L1UsdcBridgeProxyDeployedBin
}
