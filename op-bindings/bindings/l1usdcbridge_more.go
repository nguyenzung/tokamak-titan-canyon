// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1UsdcBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/tokamak-contracts/USDC/L1//tokamak-UsdcBridge/L1UsdcBridge.sol:L1UsdcBridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/tokamak-contracts/USDC/L1//tokamak-UsdcBridge/L1UsdcBridge.sol:L1UsdcBridge\",\"label\":\"otherBridge\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1002,\"contract\":\"src/tokamak-contracts/USDC/L1//tokamak-UsdcBridge/L1UsdcBridge.sol:L1UsdcBridge\",\"label\":\"l1Usdc\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"src/tokamak-contracts/USDC/L1//tokamak-UsdcBridge/L1UsdcBridge.sol:L1UsdcBridge\",\"label\":\"l2Usdc\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/tokamak-contracts/USDC/L1//tokamak-UsdcBridge/L1UsdcBridge.sol:L1UsdcBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L1UsdcBridgeStorageLayout = new(solc.StorageLayout)

var L1UsdcBridgeDeployedBin = "0x608060405234801561001057600080fd5b50600436106100a35760003560e01c80638f601f6611610076578063a1b4bc041161005b578063a1b4bc0414610191578063a9f9e675146101b1578063c89701a2146101c457600080fd5b80638f601f661461013a57806391c49bf81461017357600080fd5b80633cb747bf146100a857806356c3b587146100f257806358a997f614610112578063838b252014610127575b600080fd5b6000546100c89073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6002546100c89073ffffffffffffffffffffffffffffffffffffffff1681565b610125610120366004610d28565b6101e4565b005b610125610135366004610dab565b610290565b610165610148366004610e41565b600460209081526000928352604080842090915290825290205481565b6040519081526020016100e9565b60015473ffffffffffffffffffffffffffffffffffffffff166100c8565b6003546100c89073ffffffffffffffffffffffffffffffffffffffff1681565b6101256101bf366004610e7a565b6102a9565b6001546100c89073ffffffffffffffffffffffffffffffffffffffff1681565b333b15610278576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b610288868633338888888861064b565b505050505050565b6102a0878733888888888861064b565b50505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314801561037e5750600154600054604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9384169390921691636e296e45916004808201926020929091908290030181865afa158015610342573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103669190610ef3565b73ffffffffffffffffffffffffffffffffffffffff16145b610430576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a40161026f565b600254879073ffffffffffffffffffffffffffffffffffffffff8083169116146104b6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f74204c312075736463000000000000000000000000000000000000000000604482015260640161026f565b600354879073ffffffffffffffffffffffffffffffffffffffff80831691161461053c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f74204c322075736463000000000000000000000000000000000000000000604482015260640161026f565b73ffffffffffffffffffffffffffffffffffffffff808a166000908152600460209081526040808320938c168352929052205461057a908690610f3f565b73ffffffffffffffffffffffffffffffffffffffff808b166000818152600460209081526040808320948e16835293905291909120919091556105be908787610990565b8673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168a73ffffffffffffffffffffffffffffffffffffffff167f3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3898989896040516106389493929190610f9f565b60405180910390a4505050505050505050565b600254889073ffffffffffffffffffffffffffffffffffffffff8083169116146106d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f74204c312075736463000000000000000000000000000000000000000000604482015260640161026f565b600354889073ffffffffffffffffffffffffffffffffffffffff808316911614610757576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f6e6f74204c322075736463000000000000000000000000000000000000000000604482015260640161026f565b61077973ffffffffffffffffffffffffffffffffffffffff8b16893089610a16565b73ffffffffffffffffffffffffffffffffffffffff808b166000908152600460209081526040808320938d16835292905220546107b7908790610fd5565b73ffffffffffffffffffffffffffffffffffffffff808c1660009081526004602090815260408083208e851684529091528082209390935554600154925190821692633dbb202b9216907f662a633a000000000000000000000000000000000000000000000000000000009061083d908f908f908f908f908f908e908e90602401610fed565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b90921682526108d092918a90600401611076565b600060405180830381600087803b1580156108ea57600080fd5b505af11580156108fe573d6000803e3d6000fd5b505050508773ffffffffffffffffffffffffffffffffffffffff168973ffffffffffffffffffffffffffffffffffffffff168b73ffffffffffffffffffffffffffffffffffffffff167f718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d03968a8a898960405161097c9493929190610f9f565b60405180910390a450505050505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff838116602483015260448201839052610a1191859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050610a62565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff8481166024830152838116604483015260648201839052610a5c9186918216906323b872dd906084016109ca565b50505050565b6000610a8473ffffffffffffffffffffffffffffffffffffffff841683610af8565b90508051600014158015610aa9575080806020019051810190610aa791906110f2565b155b15610a11576040517f5274afe700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416600482015260240161026f565b6060610b0683836000610b0d565b9392505050565b606081471015610b4b576040517fcd78605900000000000000000000000000000000000000000000000000000000815230600482015260240161026f565b6000808573ffffffffffffffffffffffffffffffffffffffff168486604051610b749190611114565b60006040518083038185875af1925050503d8060008114610bb1576040519150601f19603f3d011682016040523d82523d6000602084013e610bb6565b606091505b5091509150610bc6868383610bd0565b9695505050505050565b606082610be557610be082610c5f565b610b06565b8151158015610c09575073ffffffffffffffffffffffffffffffffffffffff84163b155b15610c58576040517f9996b31500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161026f565b5080610b06565b805115610c6f5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b73ffffffffffffffffffffffffffffffffffffffff81168114610ca157600080fd5b803563ffffffff81168114610cda57600080fd5b919050565b60008083601f840112610cf157600080fd5b50813567ffffffffffffffff811115610d0957600080fd5b602083019150836020828501011115610d2157600080fd5b9250929050565b60008060008060008060a08789031215610d4157600080fd5b8635610d4c81610ca4565b95506020870135610d5c81610ca4565b945060408701359350610d7160608801610cc6565b9250608087013567ffffffffffffffff811115610d8d57600080fd5b610d9989828a01610cdf565b979a9699509497509295939492505050565b600080600080600080600060c0888a031215610dc657600080fd5b8735610dd181610ca4565b96506020880135610de181610ca4565b95506040880135610df181610ca4565b945060608801359350610e0660808901610cc6565b925060a088013567ffffffffffffffff811115610e2257600080fd5b610e2e8a828b01610cdf565b989b979a50959850939692959293505050565b60008060408385031215610e5457600080fd5b8235610e5f81610ca4565b91506020830135610e6f81610ca4565b809150509250929050565b600080600080600080600060c0888a031215610e9557600080fd5b8735610ea081610ca4565b96506020880135610eb081610ca4565b95506040880135610ec081610ca4565b94506060880135610ed081610ca4565b93506080880135925060a088013567ffffffffffffffff811115610e2257600080fd5b600060208284031215610f0557600080fd5b8151610b0681610ca4565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015610f5157610f51610f10565b500390565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff85168152836020820152606060408201526000610bc6606083018486610f56565b60008219821115610fe857610fe8610f10565b500190565b600073ffffffffffffffffffffffffffffffffffffffff808a1683528089166020840152808816604084015280871660608401525084608083015260c060a083015261103d60c083018486610f56565b9998505050505050505050565b60005b8381101561106557818101518382015260200161104d565b83811115610a5c5750506000910152565b73ffffffffffffffffffffffffffffffffffffffff8416815260606020820152600083518060608401526110b181608085016020880161104a565b63ffffffff93909316604083015250601f919091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160160800192915050565b60006020828403121561110457600080fd5b81518015158114610b0657600080fd5b6000825161112681846020870161104a565b919091019291505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L1UsdcBridgeStorageLayoutJSON), L1UsdcBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1UsdcBridge"] = L1UsdcBridgeStorageLayout
	deployedBytecodes["L1UsdcBridge"] = L1UsdcBridgeDeployedBin
	immutableReferences["L1UsdcBridge"] = false
}
