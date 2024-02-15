// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1004,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1006,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1008,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_array(t_uint256)42_storage\"},{\"astId\":1017,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"PORTAL\",\"offset\":0,\"slot\":\"249\",\"type\":\"t_contract(OptimismPortal)1019\"},{\"astId\":1018,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"nativeTokenAddress\",\"offset\":0,\"slot\":\"250\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)42_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[42]\",\"numberOfBytes\":\"1344\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(OptimismPortal)1019\":{\"encoding\":\"inplace\",\"label\":\"contract OptimismPortal\",\"numberOfBytes\":\"20\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x6080604052600436106101965760003560e01c80635644cfdf116100e15780639fce812c1161008a578063b28ade2511610064578063b28ade2514610501578063cb8398bf14610521578063d764ad0b14610541578063ecc704281461055457600080fd5b80639fce812c1461046d578063a4e7f8bd146104a1578063b1b1b209146104d157600080fd5b806383a74074116100bb57806383a74074146104285780638cbeeef21461033957806392a162cf1461043f57600080fd5b80635644cfdf146103d25780636425666b146103e85780636e296e451461041357600080fd5b80633f827a5a116101435780634c1d6a691161011d5780634c1d6a69146103395780634d0047ee1461034f57806354fd4d501461037c57600080fd5b80633f827a5a146102d15780634273ca16146102f9578063485cc9551461031957600080fd5b80630ff754ea116101745780630ff754ea146102555780632828d7e8146102a75780633dbb202b146102bc57600080fd5b806301ffc9a71461019b578063028f85f7146102125780630c56849814610240575b600080fd5b3480156101a757600080fd5b506101fd6101b636600461214c565b7fffffffff00000000000000000000000000000000000000000000000000000000167f4273ca16000000000000000000000000000000000000000000000000000000001490565b60405190151581526020015b60405180910390f35b34801561021e57600080fd5b50610227601081565b60405167ffffffffffffffff9091168152602001610209565b34801561024c57600080fd5b50610227603f81565b34801561026157600080fd5b5060f9546102829073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610209565b3480156102b357600080fd5b50610227604081565b6102cf6102ca366004612215565b6105b9565b005b3480156102dd57600080fd5b506102e6600181565b60405161ffff9091168152602001610209565b34801561030557600080fd5b506101fd610314366004612356565b61081d565b34801561032557600080fd5b506102cf6103343660046123c2565b610980565b34801561034557600080fd5b50610227619c4081565b34801561035b57600080fd5b5060fa546102829073ffffffffffffffffffffffffffffffffffffffff1681565b34801561038857600080fd5b506103c56040518060400160405280600581526020017f312e372e3100000000000000000000000000000000000000000000000000000081525081565b6040516102099190612471565b3480156103de57600080fd5b5061022761138881565b3480156103f457600080fd5b5060f95473ffffffffffffffffffffffffffffffffffffffff16610282565b34801561041f57600080fd5b50610282610c9d565b34801561043457600080fd5b5061022762030d4081565b34801561044b57600080fd5b5061045f61045a366004612484565b610d84565b6040516102099291906124c1565b34801561047957600080fd5b506102827f000000000000000000000000000000000000000000000000000000000000000081565b3480156104ad57600080fd5b506101fd6104bc3660046124e0565b60ce6020526000908152604090205460ff1681565b3480156104dd57600080fd5b506101fd6104ec3660046124e0565b60cb6020526000908152604090205460ff1681565b34801561050d57600080fd5b5061022761051c3660046124f9565b610dbc565b34801561052d57600080fd5b506102cf61053c36600461254d565b610e2c565b6102cf61054f3660046125be565b61100f565b34801561056057600080fd5b506105ab60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b604051908152602001610209565b6106f27f00000000000000000000000000000000000000000000000000000000000000006105e8858585610dbc565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061065460cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c604051602401610670979695949392919061268d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261194d565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561077760cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b866040516107899594939291906126ec565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60fa5460009073ffffffffffffffffffffffffffffffffffffffff1633146108a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f6f6e6c792061636365707420544f4e20617070726f76652063616c6c6261636b60448201526064015b60405180910390fd5b6000806108b284610d84565b909250905084156108e25760fa546108e29073ffffffffffffffffffffffffffffffffffffffff16883088611a77565b60f9546040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e9e05c4290610941908a9089908790600090889060040161273a565b600060405180830381600087803b15801561095b57600080fd5b505af115801561096f573d6000803e3d6000fd5b5060019a9950505050505050505050565b6000546003907501000000000000000000000000000000000000000000900460ff161580156109ce575060005460ff8083167401000000000000000000000000000000000000000090920416105b610a5a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161089d565b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff60ff84167401000000000000000000000000000000000000000002167fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff90911617750100000000000000000000000000000000000000000017905560f9805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560fa805492851692909116919091179055610b37611b12565b60f95473ffffffffffffffffffffffffffffffffffffffff1615801590610b75575060fa5473ffffffffffffffffffffffffffffffffffffffff1615155b15610c3a5760fa5460f9546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602482015291169063095ea7b3906044016020604051808303816000875af1158015610c14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c389190612783565b505b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610d67576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f74207365740000000000000000000000606482015260840161089d565b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b60006060600483511015610da957505060408051602081019091526000815262030d40905b5050602081015160e01c91602490910190565b6000611388619c4080603f610dd8604063ffffffff88166127d4565b610de29190612804565b610ded6010886127d4565b610dfa9062030d40612852565b610e049190612852565b610e0e9190612852565b610e189190612852565b610e229190612852565b90505b9392505050565b610ee37f0000000000000000000000000000000000000000000000000000000000000000610e5b858585610dbc565b867fd764ad0b00000000000000000000000000000000000000000000000000000000610ec760cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338b8b898c8c604051602401610670979695949392919061268d565b8473ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a338585610f6860cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610f7a9594939291906126ec565b60405180910390a260405184815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff000000000000000000000000000000000000000000000000000000000000909116179055505050565b60f087901c600281106110ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a40161089d565b8061ffff166000036111bf57600061111b878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f9250611beb915050565b600081815260cb602052604090205490915060ff16156111bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c61796564000000000000000000606482015260840161089d565b505b6000611205898989898989898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611c0a92505050565b905061120f611c2d565b156112465734156112225761122261287e565b600081815260ce602052604090205460ff16156112415761124161287e565b611398565b34156112fa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a40161089d565b600081815260ce602052604090205460ff16611398576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c6179656400000000000000000000000000000000606482015260840161089d565b6113a187611d23565b15611454576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a40161089d565b600081815260cb602052604090205460ff16156114f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c6179656400000000000000000000606482015260840161089d565b61151485611505611388619c40612852565b67ffffffffffffffff16611d69565b158061153a575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b1561165357600081815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff320161164c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161089d565b505061193f565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a161790556001861561173b5760fa546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018a90529091169063a9059cbb906044016020604051808303816000875af1158015611714573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117389190612783565b90505b600061178d89619c405a61174f91906128ad565b600089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611d8792505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905590508080156117c55750815b1561182d57600083815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a261193a565b600083815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff320161193a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161089d565b505050505b50505050505050565b905090565b34156119b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f44656e79206465706f736974696e672045544800000000000000000000000000604482015260640161089d565b81156119e05760fa546119e09073ffffffffffffffffffffffffffffffffffffffff16333085611a77565b60f9546040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e9e05c4290611a3f9087908690889060009088906004016128c4565b600060405180830381600087803b158015611a5957600080fd5b505af1158015611a6d573d6000803e3d6000fd5b5050505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052611b0c908590611da1565b50505050565b6000547501000000000000000000000000000000000000000000900460ff16611bbd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161089d565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b6000611bf985858585611eb2565b805190602001209050949350505050565b6000611c1a878787878787611f4b565b8051906020012090509695505050505050565b60f95460009073ffffffffffffffffffffffffffffffffffffffff1633148015611948575060f954604080517f9bf62d82000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691639bf62d829160048083019260209291908290030181865afa158015611ce3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d079190612911565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff8216301480611d63575060f95473ffffffffffffffffffffffffffffffffffffffff8381169116145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b6000611e03826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611fea9092919063ffffffff16565b805190915015611ead5780806020019051810190611e219190612783565b611ead576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161089d565b505050565b606084848484604051602401611ecb949392919061292e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b6060868686868686604051602401611f6896959493929190612978565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b6060610e2284846000858573ffffffffffffffffffffffffffffffffffffffff85163b612073576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161089d565b6000808673ffffffffffffffffffffffffffffffffffffffff16858760405161209c91906129cf565b60006040518083038185875af1925050503d80600081146120d9576040519150601f19603f3d011682016040523d82523d6000602084013e6120de565b606091505b50915091506120ee8282866120f9565b979650505050505050565b60608315612108575081610e25565b8251156121185782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161089d9190612471565b60006020828403121561215e57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610e2557600080fd5b73ffffffffffffffffffffffffffffffffffffffff811681146121b057600080fd5b50565b60008083601f8401126121c557600080fd5b50813567ffffffffffffffff8111156121dd57600080fd5b6020830191508360208285010111156121f557600080fd5b9250929050565b803563ffffffff8116811461221057600080fd5b919050565b6000806000806060858703121561222b57600080fd5b84356122368161218e565b9350602085013567ffffffffffffffff81111561225257600080fd5b61225e878288016121b3565b90945092506122719050604086016121fc565b905092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126122bc57600080fd5b813567ffffffffffffffff808211156122d7576122d761227c565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561231d5761231d61227c565b8160405283815286602085880101111561233657600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000806080858703121561236c57600080fd5b84356123778161218e565b935060208501356123878161218e565b925060408501359150606085013567ffffffffffffffff8111156123aa57600080fd5b6123b6878288016122ab565b91505092959194509250565b600080604083850312156123d557600080fd5b82356123e08161218e565b915060208301356123f08161218e565b809150509250929050565b60005b838110156124165781810151838201526020016123fe565b83811115611b0c5750506000910152565b6000815180845261243f8160208601602086016123fb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610e256020830184612427565b60006020828403121561249657600080fd5b813567ffffffffffffffff8111156124ad57600080fd5b6124b9848285016122ab565b949350505050565b63ffffffff83168152604060208201526000610e226040830184612427565b6000602082840312156124f257600080fd5b5035919050565b60008060006040848603121561250e57600080fd5b833567ffffffffffffffff81111561252557600080fd5b612531868287016121b3565b90945092506125449050602085016121fc565b90509250925092565b60008060008060006080868803121561256557600080fd5b85356125708161218e565b945060208601359350604086013567ffffffffffffffff81111561259357600080fd5b61259f888289016121b3565b90945092506125b29050606087016121fc565b90509295509295909350565b600080600080600080600060c0888a0312156125d957600080fd5b8735965060208801356125eb8161218e565b955060408801356125fb8161218e565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561262557600080fd5b6126318a828b016121b3565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a08301526126df60c083018486612644565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815260806020820152600061271c608083018688612644565b905083604083015263ffffffff831660608301529695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015263ffffffff84166040820152821515606082015260a0608082015260006120ee60a0830184612427565b60006020828403121561279557600080fd5b81518015158114610e2557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff808316818516818304811182151516156127fb576127fb6127a5565b02949350505050565b600067ffffffffffffffff80841680612846577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115612875576128756127a5565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6000828210156128bf576128bf6127a5565b500390565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015267ffffffffffffffff84166040820152821515606082015260a0608082015260006120ee60a0830184612427565b60006020828403121561292357600080fd5b8151610e258161218e565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250608060408301526129676080830185612427565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a08301526129c360c0830184612427565b98975050505050505050565b600082516129e18184602087016123fb565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
}
