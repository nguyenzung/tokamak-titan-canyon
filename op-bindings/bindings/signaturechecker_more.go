// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SignatureCheckerStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var SignatureCheckerStorageLayout = new(solc.StorageLayout)

var SignatureCheckerDeployedBin = "0x73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c80636ccea6521461003a575b600080fd5b6101026004803603606081101561005057600080fd5b73ffffffffffffffffffffffffffffffffffffffff8235169160208101359181019060608101604082013564010000000081111561008d57600080fd5b82018360208201111561009f57600080fd5b803590602001918460018302840111640100000000831117156100c157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610116945050505050565b604080519115158252519081900360200190f35b600061012184610179565b610164578373ffffffffffffffffffffffffffffffffffffffff16610146848461017f565b73ffffffffffffffffffffffffffffffffffffffff16149050610172565b61016f848484610203565b90505b9392505050565b3b151590565b600081516041146101db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806106296023913960400191505060405180910390fd5b60208201516040830151606084015160001a6101f98682858561042d565b9695505050505050565b60008060608573ffffffffffffffffffffffffffffffffffffffff16631626ba7e60e01b86866040516024018083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561026f578181015183820152602001610257565b50505050905090810190601f16801561029c5780820380516001836020036101000a031916815260200191505b50604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009098169790971787525181519196909550859450925090508083835b6020831061036957805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0909201916020918201910161032c565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d80600081146103c9576040519150601f19603f3d011682016040523d82523d6000602084013e6103ce565b606091505b50915091508180156103e257506020815110155b80156101f9575080517f1626ba7e00000000000000000000000000000000000000000000000000000000906020808401919081101561042057600080fd5b5051149695505050505050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156104a8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806106726026913960400191505060405180910390fd5b8360ff16601b141580156104c057508360ff16601c14155b15610516576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018061064c6026913960400191505060405180910390fd5b600060018686868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015610572573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661061f57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f45435265636f7665723a20696e76616c6964207369676e617475726500000000604482015290519081900360640190fd5b9594505050505056fe45435265636f7665723a20696e76616c6964207369676e6174757265206c656e67746845435265636f7665723a20696e76616c6964207369676e6174757265202776272076616c756545435265636f7665723a20696e76616c6964207369676e6174757265202773272076616c7565a164736f6c634300060c000a"

func init() {
	if err := json.Unmarshal([]byte(SignatureCheckerStorageLayoutJSON), SignatureCheckerStorageLayout); err != nil {
		panic(err)
	}

	layouts["SignatureChecker"] = SignatureCheckerStorageLayout
	deployedBytecodes["SignatureChecker"] = SignatureCheckerDeployedBin
}
