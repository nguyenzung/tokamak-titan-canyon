// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/tokamak-network/tokamak-thanos/op-bindings/solc"
)

const L2UsdcBridgeProxyStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/tokamak-contracts/USDC/L2/tokamak-UsdcBridge/L2UsdcBridgeProxy.sol:L2UsdcBridgeProxy\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/tokamak-contracts/USDC/L2/tokamak-UsdcBridge/L2UsdcBridgeProxy.sol:L2UsdcBridgeProxy\",\"label\":\"otherBridge\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1002,\"contract\":\"src/tokamak-contracts/USDC/L2/tokamak-UsdcBridge/L2UsdcBridgeProxy.sol:L2UsdcBridgeProxy\",\"label\":\"l1Usdc\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"src/tokamak-contracts/USDC/L2/tokamak-UsdcBridge/L2UsdcBridgeProxy.sol:L2UsdcBridgeProxy\",\"label\":\"l2Usdc\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/tokamak-contracts/USDC/L2/tokamak-UsdcBridge/L2UsdcBridgeProxy.sol:L2UsdcBridgeProxy\",\"label\":\"l2UsdcMasterMinter\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"}}}"

var L2UsdcBridgeProxyStorageLayout = new(solc.StorageLayout)

var L2UsdcBridgeProxyDeployedBin = "0x6080604052600436106100c05760003560e01c806356c3b58711610074578063a1b4bc041161004e578063a1b4bc0414610270578063c89701a21461029d578063dfd3dcb3146102ca5761012c565b806356c3b587146102195780635c60da1b146102465780638da5cb5b1461025b5761012c565b80633659cfe6116100a55780633659cfe6146101ac5780633cb747bf146101cc5780634f1ef286146101f95761012c565b806305db940f1461013657806312c594881461018c5761012c565b3661012c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f63616e6e6f74207265636569766520544f4e000000000000000000000000000060448201526064015b60405180910390fd5b6101346102ea565b005b34801561014257600080fd5b506004546101639073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34801561019857600080fd5b506101346101a7366004610d18565b6102fc565b3480156101b857600080fd5b506101346101c7366004610d7d565b610698565b3480156101d857600080fd5b506000546101639073ffffffffffffffffffffffffffffffffffffffff1681565b34801561020557600080fd5b50610134610214366004610dc7565b610750565b34801561022557600080fd5b506002546101639073ffffffffffffffffffffffffffffffffffffffff1681565b34801561025257600080fd5b506101636107fa565b34801561026757600080fd5b50610163610809565b34801561027c57600080fd5b506003546101639073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102a957600080fd5b506001546101639073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102d657600080fd5b506101346102e5366004610d7d565b610813565b6102fa6102f56108b8565b6108c2565b565b610304610809565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610398576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f6e6f74206f776e657200000000000000000000000000000000000000000000006044820152606401610123565b8473ffffffffffffffffffffffffffffffffffffffff8116610416576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f7a65726f206164647265737300000000000000000000000000000000000000006044820152606401610123565b8473ffffffffffffffffffffffffffffffffffffffff8116610494576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f7a65726f206164647265737300000000000000000000000000000000000000006044820152606401610123565b8473ffffffffffffffffffffffffffffffffffffffff8116610512576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f7a65726f206164647265737300000000000000000000000000000000000000006044820152606401610123565b8473ffffffffffffffffffffffffffffffffffffffff8116610590576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f7a65726f206164647265737300000000000000000000000000000000000000006044820152606401610123565b8473ffffffffffffffffffffffffffffffffffffffff811661060e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f7a65726f206164647265737300000000000000000000000000000000000000006044820152606401610123565b5050600080547fffffffffffffffffffffffff000000000000000000000000000000000000000090811673ffffffffffffffffffffffffffffffffffffffff9a8b1617909155600180548216988a16989098179097555050600280548616948716949094179093556003805485169286169290921790915560048054909316931692909217905550565b6106a0610809565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610734576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f6e6f74206f776e657200000000000000000000000000000000000000000000006044820152606401610123565b61074d81604051806020016040528060008152506108e6565b50565b610758610809565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f6e6f74206f776e657200000000000000000000000000000000000000000000006044820152606401610123565b6107f682826108e6565b5050565b60006108046108b8565b905090565b600061080461094e565b61081b610809565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f6e6f74206f776e657200000000000000000000000000000000000000000000006044820152606401610123565b61074d8161098e565b60006108046109ef565b3660008037600080366000845af43d6000803e8080156108e1573d6000f35b3d6000fd5b6108ef82610a17565b60405173ffffffffffffffffffffffffffffffffffffffff8316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a2805115610946576109418282610ae9565b505050565b6107f6610b6c565b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b5473ffffffffffffffffffffffffffffffffffffffff16919050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6109b761094e565b6040805173ffffffffffffffffffffffffffffffffffffffff928316815291841660208301520160405180910390a161074d81610ba4565b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc610972565b8073ffffffffffffffffffffffffffffffffffffffff163b600003610a80576040517f4c9c8ce300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff82166004820152602401610123565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b60606000808473ffffffffffffffffffffffffffffffffffffffff1684604051610b139190610ea7565b600060405180830381855af49150503d8060008114610b4e576040519150601f19603f3d011682016040523d82523d6000602084013e610b53565b606091505b5091509150610b63858383610c1b565b95945050505050565b34156102fa576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610bf4576040517f62e77ba200000000000000000000000000000000000000000000000000000000815260006004820152602401610123565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103610aa3565b606082610c3057610c2b82610cad565b610ca6565b8151158015610c54575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610ca3576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85166004820152602401610123565b50805b9392505050565b805115610cbd5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b803573ffffffffffffffffffffffffffffffffffffffff81168114610d1357600080fd5b919050565b600080600080600060a08688031215610d3057600080fd5b610d3986610cef565b9450610d4760208701610cef565b9350610d5560408701610cef565b9250610d6360608701610cef565b9150610d7160808701610cef565b90509295509295909350565b600060208284031215610d8f57600080fd5b610ca682610cef565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060408385031215610dda57600080fd5b610de383610cef565b9150602083013567ffffffffffffffff80821115610e0057600080fd5b818501915085601f830112610e1457600080fd5b813581811115610e2657610e26610d98565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610e6c57610e6c610d98565b81604052828152886020848701011115610e8557600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6000825160005b81811015610ec85760208186018101518583015201610eae565b50600092019182525091905056fea164736f6c6343000819000a"


func init() {
	if err := json.Unmarshal([]byte(L2UsdcBridgeProxyStorageLayoutJSON), L2UsdcBridgeProxyStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2UsdcBridgeProxy"] = L2UsdcBridgeProxyStorageLayout
	deployedBytecodes["L2UsdcBridgeProxy"] = L2UsdcBridgeProxyDeployedBin
	immutableReferences["L2UsdcBridgeProxy"] = false
}
