// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1004,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1006,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1008,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"src/L2/L2CrossDomainMessenger.sol:L2CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_array(t_uint256)42_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)42_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[42]\",\"numberOfBytes\":\"1344\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L2CrossDomainMessengerDeployedBin = "0x60806040526004361061015f5760003560e01c806383a74074116100c0578063b1b1b20911610074578063cb8398bf11610059578063cb8398bf146103dd578063d764ad0b146103fd578063ecc704281461041057600080fd5b8063b1b1b2091461038d578063b28ade25146103bd57600080fd5b80639fce812c116100a55780639fce812c146102e6578063a4e7f8bd1461031a578063a71198691461035a57600080fd5b806383a74074146102cf5780638cbeeef2146101fe57600080fd5b80634c1d6a69116101175780635644cfdf116100fc5780635644cfdf1461026a5780636e296e45146102805780638129fc1c146102ba57600080fd5b80634c1d6a69146101fe57806354fd4d501461021457600080fd5b80632828d7e8116101485780632828d7e8146101ac5780633dbb202b146101c15780633f827a5a146101d657600080fd5b8063028f85f7146101645780630c56849814610197575b600080fd5b34801561017057600080fd5b50610179601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101a357600080fd5b50610179603f81565b3480156101b857600080fd5b50610179604081565b6101d46101cf3660046119b4565b610475565b005b3480156101e257600080fd5b506101eb600181565b60405161ffff909116815260200161018e565b34801561020a57600080fd5b50610179619c4081565b34801561022057600080fd5b5061025d6040518060400160405280600581526020017f312e372e3000000000000000000000000000000000000000000000000000000081525081565b60405161018e9190611a84565b34801561027657600080fd5b5061017961138881565b34801561028c57600080fd5b506102956106d9565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018e565b3480156102c657600080fd5b506101d46107c5565b3480156102db57600080fd5b5061017962030d4081565b3480156102f257600080fd5b506102957f000000000000000000000000000000000000000000000000000000000000000081565b34801561032657600080fd5b5061034a610335366004611a9e565b60ce6020526000908152604090205460ff1681565b604051901515815260200161018e565b34801561036657600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610295565b34801561039957600080fd5b5061034a6103a8366004611a9e565b60cb6020526000908152604090205460ff1681565b3480156103c957600080fd5b506101796103d8366004611ab7565b610988565b3480156103e957600080fd5b506101d46103f8366004611b0b565b6109f6565b6101d461040b366004611b7a565b610c5b565b34801561041c57600080fd5b5061046760cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b60405190815260200161018e565b6105ae7f00000000000000000000000000000000000000000000000000000000000000006104a4858585610988565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061051060cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c60405160240161052c9796959493929190611c45565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611538565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561063360cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610645959493929190611ca4565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2153016107a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6000546003907501000000000000000000000000000000000000000000900460ff16158015610813575060005460ff8083167401000000000000000000000000000000000000000090920416105b61089f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161079f565b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff60ff84167401000000000000000000000000000000000000000002167fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff9091161775010000000000000000000000000000000000000000001790556109276115c6565b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b6000611388619c4080603f6109a4604063ffffffff8816611d21565b6109ae9190611d51565b6109b9601088611d21565b6109c69062030d40611d9f565b6109d09190611d9f565b6109da9190611d9f565b6109e49190611d9f565b6109ee9190611d9f565b949350505050565b610b2f7f0000000000000000000000000000000000000000000000000000000000000000610a25858585610988565b867fd764ad0b00000000000000000000000000000000000000000000000000000000610a9160cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338b8b898c8c604051602401610aad9796959493929190611c45565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261169f565b8473ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a338585610bb460cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610bc6959493929190611ca4565b60405180910390a260405184815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff000000000000000000000000000000000000000000000000000000000000909116179055505050565b60f087901c60028110610d16576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a40161079f565b8061ffff16600003610e0b576000610d67878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f9250611727915050565b600081815260cb602052604090205490915060ff1615610e09576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c61796564000000000000000000606482015260840161079f565b505b6000610e51898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061174692505050565b905073ffffffffffffffffffffffffffffffffffffffff7fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330181167f000000000000000000000000000000000000000000000000000000000000000090911603610ee957853414610ec557610ec5611dcb565b600081815260ce602052604090205460ff1615610ee457610ee4611dcb565b61103b565b3415610f9d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a40161079f565b600081815260ce602052604090205460ff1661103b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c6179656400000000000000000000000000000000606482015260840161079f565b61104487611769565b156110f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a40161079f565b600081815260cb602052604090205460ff1615611196576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c6179656400000000000000000000606482015260840161079f565b6111b7856111a8611388619c40611d9f565b67ffffffffffffffff166117be565b15806111dd575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b156112f657600081815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016112ef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161079f565b505061152f565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a16179055600061138788619c405a61134a9190611dfa565b8988888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506117dc92505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790559050801561141e57600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a261152b565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff320161152b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161079f565b5050505b50505050505050565b6040517fc2b3e5ac0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000169063c2b3e5ac90849061158e90889088908790600401611e11565b6000604051808303818588803b1580156115a757600080fd5b505af11580156115bb573d6000803e3d6000fd5b505050505050505050565b6000547501000000000000000000000000000000000000000000900460ff16611671576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161079f565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43616e6e6f742063616c6c20746869732066756e6374696f6e2065786365707460448201527f206974206973206f766572726964656400000000000000000000000000000000606482015260840161079f565b6000611735858585856117f6565b805190602001209050949350505050565b600061175687878787878761188f565b8051906020012090509695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff82163014806117b8575073ffffffffffffffffffffffffffffffffffffffff8216734200000000000000000000000000000000000016145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b60608484848460405160240161180f9493929190611e59565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b60608686868686866040516024016118ac96959493929190611ea3565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461195257600080fd5b919050565b60008083601f84011261196957600080fd5b50813567ffffffffffffffff81111561198157600080fd5b60208301915083602082850101111561199957600080fd5b9250929050565b803563ffffffff8116811461195257600080fd5b600080600080606085870312156119ca57600080fd5b6119d38561192e565b9350602085013567ffffffffffffffff8111156119ef57600080fd5b6119fb87828801611957565b9094509250611a0e9050604086016119a0565b905092959194509250565b6000815180845260005b81811015611a3f57602081850181015186830182015201611a23565b81811115611a51576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611a976020830184611a19565b9392505050565b600060208284031215611ab057600080fd5b5035919050565b600080600060408486031215611acc57600080fd5b833567ffffffffffffffff811115611ae357600080fd5b611aef86828701611957565b9094509250611b029050602085016119a0565b90509250925092565b600080600080600060808688031215611b2357600080fd5b611b2c8661192e565b945060208601359350604086013567ffffffffffffffff811115611b4f57600080fd5b611b5b88828901611957565b9094509250611b6e9050606087016119a0565b90509295509295909350565b600080600080600080600060c0888a031215611b9557600080fd5b87359650611ba56020890161192e565b9550611bb36040890161192e565b9450606088013593506080880135925060a088013567ffffffffffffffff811115611bdd57600080fd5b611be98a828b01611957565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611c9760c083018486611bfc565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000611cd4608083018688611bfc565b905083604083015263ffffffff831660608301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611d4857611d48611cf2565b02949350505050565b600067ffffffffffffffff80841680611d93577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611dc257611dc2611cf2565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082821015611e0c57611e0c611cf2565b500390565b73ffffffffffffffffffffffffffffffffffffffff8416815267ffffffffffffffff83166020820152606060408201526000611e506060830184611a19565b95945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152611e926080830185611a19565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152611eee60c0830184611a19565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2CrossDomainMessengerStorageLayoutJSON), L2CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2CrossDomainMessenger"] = L2CrossDomainMessengerStorageLayout
	deployedBytecodes["L2CrossDomainMessenger"] = L2CrossDomainMessengerDeployedBin
}
