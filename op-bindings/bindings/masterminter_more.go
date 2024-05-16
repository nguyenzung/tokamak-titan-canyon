// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const MasterMinterStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/tokamak-contracts/USDC/L2/tokamak-USDC/minting/MasterMinter.sol:MasterMinter\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/tokamak-contracts/USDC/L2/tokamak-USDC/minting/MasterMinter.sol:MasterMinter\",\"label\":\"controllers\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_address)\"},{\"astId\":1002,\"contract\":\"src/tokamak-contracts/USDC/L2/tokamak-USDC/minting/MasterMinter.sol:MasterMinter\",\"label\":\"minterManager\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_contract(MinterManagementInterface)1003\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_contract(MinterManagementInterface)1003\":{\"encoding\":\"inplace\",\"label\":\"contract MinterManagementInterface\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_address)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e address)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_address\"}}}"

var MasterMinterStorageLayout = new(solc.StorageLayout)

var MasterMinterDeployedBin = "0x608060405234801561001057600080fd5b50600436106100c95760003560e01c8063c011b1c311610081578063ea7215691161005b578063ea72156914610215578063f2fde38b1461021d578063f6a74ed714610250576100c9565b8063c011b1c31461018a578063c4faf7df146101bd578063cbf2b8bf146101f8576100c9565b80637c6b8ef5116100b25780637c6b8ef5146101345780638da5cb5b146101515780639398608b14610182576100c9565b806333db2ad2146100ce578063542fef91146100ff575b600080fd5b6100eb600480360360208110156100e457600080fd5b5035610283565b604080519115158252519081900360200190f35b6101326004803603602081101561011557600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610573565b005b6100eb6004803603602081101561014a57600080fd5b5035610687565b61015961098b565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101596109a7565b610159600480360360208110156101a057600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166109c3565b610132600480360360408110156101d357600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160200135166109ee565b6100eb6004803603602081101561020e57600080fd5b5035610bc8565b6100eb610cb5565b6101326004803603602081101561023357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610e2b565b6101326004803603602081101561026657600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610f7e565b3360009081526001602052604081205473ffffffffffffffffffffffffffffffffffffffff166102fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260358152602001806114586035913960400191505060405180910390fd5b60008211610357576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a81526020018061148d602a913960400191505060405180910390fd5b336000908152600160209081526040918290205460025483517faa271e1a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff928316600482018190529451929091169263aa271e1a92602480840193829003018186803b1580156103d957600080fd5b505afa1580156103ed573d6000803e3d6000fd5b505050506040513d602081101561040357600080fd5b505161045a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260398152602001806114b76039913960400191505060405180910390fd5b600254604080517f8a6db9c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015291516000939290921691638a6db9c391602480820192602092909190829003018186803b1580156104d157600080fd5b505afa1580156104e5573d6000803e3d6000fd5b505050506040513d60208110156104fb57600080fd5b50519050600061050b8286611161565b6040805187815260208101839052815192935073ffffffffffffffffffffffffffffffffffffffff86169233927f3703d23abba1e61f32acc0682fc062ea5c710672c7d100af5ecd08485e983ad0928290030190a361056a83826111d5565b95945050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146105f957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60025460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f9992ea32e96992be98be5c833cd5b9fd77314819d2146b6f06ab9cfef957af1290600090a3600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b3360009081526001602052604081205473ffffffffffffffffffffffffffffffffffffffff16610702576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260358152602001806114586035913960400191505060405180910390fd5b6000821161075b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001806113c2602a913960400191505060405180910390fd5b336000908152600160209081526040918290205460025483517faa271e1a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff928316600482018190529451929091169263aa271e1a92602480840193829003018186803b1580156107dd57600080fd5b505afa1580156107f1573d6000803e3d6000fd5b505050506040513d602081101561080757600080fd5b505161085e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260398152602001806114f06039913960400191505060405180910390fd5b600254604080517f8a6db9c300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015291516000939290921691638a6db9c391602480820192602092909190829003018186803b1580156108d557600080fd5b505afa1580156108e9573d6000803e3d6000fd5b505050506040513d60208110156108ff57600080fd5b5051905060008482116109125781610914565b845b905060006109228383611287565b6040805184815260208101839052815192935073ffffffffffffffffffffffffffffffffffffffff87169233927f3cc75d3bf58b0100659088c03539964108d5d06342e1bd8085ee43ad8ff6f69a928290030190a361098184826111d5565b9695505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff1690565b60025473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff9081166000908152600160205260409020541690565b60005473ffffffffffffffffffffffffffffffffffffffff163314610a7457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8216610ae0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806113ec6025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8116610b4c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806114376021913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82811660008181526001602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169486169485179055517fa56687ff5096e83f6e2c673cda0b677f56bbfcdf5fe0555d5830c407ede193cb9190a35050565b3360009081526001602052604081205473ffffffffffffffffffffffffffffffffffffffff16610c43576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260358152602001806114586035913960400191505060405180910390fd5b33600081815260016020908152604091829020548251868152925173ffffffffffffffffffffffffffffffffffffffff90911693849390927f5b0b60a4f757b33d9dcb8bd021b6aa371bb2e6f134086797aefcd8c0afab538c92918290030190a3610cae81846111d5565b9392505050565b3360009081526001602052604081205473ffffffffffffffffffffffffffffffffffffffff16610d30576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260358152602001806114586035913960400191505060405180910390fd5b3360008181526001602052604080822054905173ffffffffffffffffffffffffffffffffffffffff90911692839290917f4b5ef9a786cf64a7d82ebcf2d5132667edc9faef4ac36260d9a9e52c526b62329190a3600254604080517f3092afd500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff848116600483015291519190921691633092afd59160248083019260209291908290030181600087803b158015610df957600080fd5b505af1158015610e0d573d6000803e3d6000fd5b505050506040513d6020811015610e2357600080fd5b505191505090565b60005473ffffffffffffffffffffffffffffffffffffffff163314610eb157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8116610f1d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806114116026913960400191505060405180910390fd5b6000546040805173ffffffffffffffffffffffffffffffffffffffff9283168152918316602083015280517f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09281900390910190a1610f7b816112c9565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461100457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8116611070576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806113ec6025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff818116600090815260016020526040902054166110ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806114376021913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff811660008181526001602052604080822080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055517f33d83959be2573f5453b12eb9d43b3499bc57d96bd2f067ba44803c859e811139190a250565b600082820183811015610cae57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b600254604080517f4e44d95600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301526024820185905291516000939290921691634e44d9569160448082019260209290919082900301818787803b15801561125457600080fd5b505af1158015611268573d6000803e3d6000fd5b505050506040513d602081101561127e57600080fd5b50519392505050565b6000610cae83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611310565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600081848411156113b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561137e578181015183820152602001611366565b50505050905090810190601f1680156113ab5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056fe416c6c6f77616e63652064656372656d656e74206d7573742062652067726561746572207468616e2030436f6e74726f6c6c6572206d7573742062652061206e6f6e2d7a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373576f726b6572206d7573742062652061206e6f6e2d7a65726f20616464726573735468652076616c7565206f6620636f6e74726f6c6c6572735b6d73672e73656e6465725d206d757374206265206e6f6e2d7a65726f416c6c6f77616e636520696e6372656d656e74206d7573742062652067726561746572207468616e203043616e206f6e6c7920696e6372656d656e7420616c6c6f77616e636520666f72206d696e7465727320696e206d696e7465724d616e6167657243616e206f6e6c792064656372656d656e7420616c6c6f77616e636520666f72206d696e7465727320696e206d696e7465724d616e61676572a164736f6c634300060c000a"


func init() {
	if err := json.Unmarshal([]byte(MasterMinterStorageLayoutJSON), MasterMinterStorageLayout); err != nil {
		panic(err)
	}

	layouts["MasterMinter"] = MasterMinterStorageLayout
	deployedBytecodes["MasterMinter"] = MasterMinterDeployedBin
	immutableReferences["MasterMinter"] = false
}
