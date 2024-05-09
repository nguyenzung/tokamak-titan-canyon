// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1004,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1006,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1008,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_array(t_uint256)42_storage\"},{\"astId\":1017,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"PORTAL\",\"offset\":0,\"slot\":\"249\",\"type\":\"t_contract(OptimismPortal)1019\"},{\"astId\":1018,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"nativeTokenAddress\",\"offset\":0,\"slot\":\"250\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)42_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[42]\",\"numberOfBytes\":\"1344\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(OptimismPortal)1019\":{\"encoding\":\"inplace\",\"label\":\"contract OptimismPortal\",\"numberOfBytes\":\"20\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x6080604052600436106101965760003560e01c80635644cfdf116100e15780639fce812c1161008a578063b28ade2511610064578063b28ade25146104c3578063d764ad0b146104e3578063e0e593c5146104f6578063ecc704281461051657600080fd5b80639fce812c1461042f578063a4e7f8bd14610463578063b1b1b2091461049357600080fd5b806383a74074116100bb57806383a74074146103e65780638cbeeef2146102f757806392a162cf146103fd57600080fd5b80635644cfdf146103905780636425666b146103a65780636e296e45146103d157600080fd5b80633f827a5a116101435780634c1d6a691161011d5780634c1d6a69146102f75780634d0047ee1461030d57806354fd4d501461033a57600080fd5b80633f827a5a1461028f5780634273ca16146102b7578063485cc955146102d757600080fd5b80630ff754ea116101745780630ff754ea146102135780632828d7e8146102655780633dbb202b1461027a57600080fd5b806301ffc9a71461019b578063028f85f7146101d05780630c568498146101fe575b600080fd5b3480156101a757600080fd5b506101bb6101b6366004612309565b61057b565b60405190151581526020015b60405180910390f35b3480156101dc57600080fd5b506101e5601081565b60405167ffffffffffffffff90911681526020016101c7565b34801561020a57600080fd5b506101e5603f81565b34801561021f57600080fd5b5060f9546102409073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c7565b34801561027157600080fd5b506101e5604081565b61028d6102883660046123d2565b610614565b005b34801561029b57600080fd5b506102a4600181565b60405161ffff90911681526020016101c7565b3480156102c357600080fd5b506101bb6102d2366004612439565b610878565b3480156102e357600080fd5b5061028d6102f23660046124ac565b610a12565b34801561030357600080fd5b506101e5619c4081565b34801561031957600080fd5b5060fa546102409073ffffffffffffffffffffffffffffffffffffffff1681565b34801561034657600080fd5b506103836040518060400160405280600581526020017f312e372e3100000000000000000000000000000000000000000000000000000081525081565b6040516101c7919061255b565b34801561039c57600080fd5b506101e561138881565b3480156103b257600080fd5b5060f95473ffffffffffffffffffffffffffffffffffffffff16610240565b3480156103dd57600080fd5b50610240610c2c565b3480156103f257600080fd5b506101e562030d4081565b34801561040957600080fd5b5061041d61041836600461256e565b610d13565b6040516101c7969594939291906125f9565b34801561043b57600080fd5b506102407f000000000000000000000000000000000000000000000000000000000000000081565b34801561046f57600080fd5b506101bb61047e366004612651565b60ce6020526000908152604090205460ff1681565b34801561049f57600080fd5b506101bb6104ae366004612651565b60cb6020526000908152604090205460ff1681565b3480156104cf57600080fd5b506101e56104de36600461266a565b610dfc565b61028d6104f13660046126be565b610e6c565b34801561050257600080fd5b5061028d610511366004612744565b611854565b34801561052257600080fd5b5061056d60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016101c7565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f4273ca1600000000000000000000000000000000000000000000000000000000148061060e57507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b61074d7f0000000000000000000000000000000000000000000000000000000000000000610643858585610dfc565b347fd764ad0b000000000000000000000000000000000000000000000000000000006106af60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016106cb97969594939291906127b5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261186e565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a3385856107d260cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b866040516107e4959493929190612814565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60fa5460009073ffffffffffffffffffffffffffffffffffffffff163314610927576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f6f6e6c7920616363657074206e617469766520746f6b656e20617070726f766560448201527f2063616c6c6261636b000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60008060008036600061093a8989610d13565b9550955095509550955095508573ffffffffffffffffffffffffffffffffffffffff168c73ffffffffffffffffffffffffffffffffffffffff161480156109805750838a145b801561098c5750600084115b6109f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f696e76616c6964206f6e417070726f7665206461746100000000000000000000604482015260640161091e565b610a0086868686868661196c565b5060019b9a5050505050505050505050565b6000546003907501000000000000000000000000000000000000000000900460ff16158015610a60575060005460ff8083167401000000000000000000000000000000000000000090920416105b610aec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161091e565b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff60ff84167401000000000000000000000000000000000000000002167fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff90911617750100000000000000000000000000000000000000000017905560f9805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560fa805492851692909116919091179055610bc9611c37565b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610cf6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f74207365740000000000000000000000606482015260840161091e565b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b60008080803681604c871015610dab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f496e76616c6964206f6e417070726f7665206461746120666f72204c3143726f60448201527f7373446f6d61696e4d657373656e676572000000000000000000000000000000606482015260840161091e565b5050508435606090811c96601487013590911c95602881013595604882013560e01c9550604c90910193507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb4019150565b6000611388619c4080603f610e18604063ffffffff8816612891565b610e2291906128c1565b610e2d601088612891565b610e3a9062030d4061290f565b610e44919061290f565b610e4e919061290f565b610e58919061290f565b610e62919061290f565b90505b9392505050565b3415610efa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f000000000000000000000000000000000000000000000000606482015260840161091e565b60f087901c60028110610fb5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a40161091e565b8061ffff166000036110aa576000611006878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f9250611d10915050565b600081815260cb602052604090205490915060ff16156110a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c61796564000000000000000000606482015260840161091e565b505b60006110f0898989898989898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611d2f92505050565b90506110fa611d52565b1561115457600081815260ce602052604090205460ff161561111e5761111e61293b565b851561114f5760f95460fa5461114f9173ffffffffffffffffffffffffffffffffffffffff91821691163089611e48565b6111f2565b600081815260ce602052604090205460ff166111f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c6179656400000000000000000000000000000000606482015260840161091e565b6111fb87611ee3565b156112ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a40161091e565b600081815260cb602052604090205460ff161561134d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c6179656400000000000000000000606482015260840161091e565b61136e8561135f611388619c4061290f565b67ffffffffffffffff16611f26565b1580611394575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b156114ad57600081815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016114a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161091e565b505061184b565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a16179055600186156115955760fa546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018a90529091169063095ea7b3906044016020604051808303816000875af115801561156e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611592919061296a565b90505b60006115e789619c405a6115a9919061298c565b600089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611f4492505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905590508715801590611622575080155b156116c75760fa546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b81166004830152600060248301529091169063095ea7b3906044016020604051808303816000875af11580156116a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116c4919061296a565b91505b8080156116d15750815b1561173957600083815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a2611846565b600083815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555184917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3201611846576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d65737361676500000000000000000000000000000000000000606482015260840161091e565b505050505b50505050505050565b61186233868684878761196c565b5050505050565b905090565b34156118d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f44656e79206465706f736974696e672045544800000000000000000000000000604482015260640161091e565b60f9546040517f9a5fc9f700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690639a5fc9f79061193490879086908190899088906004016129a3565b600060405180830381600087803b15801561194e57600080fd5b505af1158015611962573d6000803e3d6000fd5b5050505050505050565b8315611a385760fa546119979073ffffffffffffffffffffffffffffffffffffffff16873087611e48565b60fa5460f9546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201526024810187905291169063095ea7b3906044016020604051808303816000875af1158015611a12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a36919061296a565b505b611aef7f0000000000000000000000000000000000000000000000000000000000000000611a67848487610dfc565b867fd764ad0b00000000000000000000000000000000000000000000000000000000611ad360cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b8b8b8b8b8b8b6040516024016106cb97969594939291906127b5565b8473ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a878484611b7460cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b88604051611b86959493929190612814565b60405180910390a28573ffffffffffffffffffffffffffffffffffffffff167f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d54685604051611bd691815260200190565b60405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff00000000000000000000000000000000000000000000000000000000000090911617905550505050565b6000547501000000000000000000000000000000000000000000900460ff16611ce2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161091e565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b6000611d1e85858585611f5e565b805190602001209050949350505050565b6000611d3f878787878787611ff7565b8051906020012090509695505050505050565b60f95460009073ffffffffffffffffffffffffffffffffffffffff1633148015611869575060f954604080517f9bf62d82000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691639bf62d829160048083019260209291908290030181865afa158015611e08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e2c91906129ee565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052611edd908590612096565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff821630148061060e57505060f95473ffffffffffffffffffffffffffffffffffffffff90811691161490565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b606084848484604051602401611f779493929190612a0b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b606086868686868660405160240161201496959493929190612a55565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b60006120f8826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166121a79092919063ffffffff16565b8051909150156121a25780806020019051810190612116919061296a565b6121a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161091e565b505050565b6060610e6284846000858573ffffffffffffffffffffffffffffffffffffffff85163b612230576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161091e565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516122599190612aa0565b60006040518083038185875af1925050503d8060008114612296576040519150601f19603f3d011682016040523d82523d6000602084013e61229b565b606091505b50915091506122ab8282866122b6565b979650505050505050565b606083156122c5575081610e65565b8251156122d55782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091e919061255b565b60006020828403121561231b57600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610e6557600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116811461236d57600080fd5b50565b60008083601f84011261238257600080fd5b50813567ffffffffffffffff81111561239a57600080fd5b6020830191508360208285010111156123b257600080fd5b9250929050565b803563ffffffff811681146123cd57600080fd5b919050565b600080600080606085870312156123e857600080fd5b84356123f38161234b565b9350602085013567ffffffffffffffff81111561240f57600080fd5b61241b87828801612370565b909450925061242e9050604086016123b9565b905092959194509250565b60008060008060006080868803121561245157600080fd5b853561245c8161234b565b9450602086013561246c8161234b565b935060408601359250606086013567ffffffffffffffff81111561248f57600080fd5b61249b88828901612370565b969995985093965092949392505050565b600080604083850312156124bf57600080fd5b82356124ca8161234b565b915060208301356124da8161234b565b809150509250929050565b60005b838110156125005781810151838201526020016124e8565b83811115611edd5750506000910152565b600081518084526125298160208601602086016124e5565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610e656020830184612511565b6000806020838503121561258157600080fd5b823567ffffffffffffffff81111561259857600080fd5b6125a485828601612370565b90969095509350505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600073ffffffffffffffffffffffffffffffffffffffff808916835280881660208401525085604083015263ffffffff8516606083015260a0608083015261264560a0830184866125b0565b98975050505050505050565b60006020828403121561266357600080fd5b5035919050565b60008060006040848603121561267f57600080fd5b833567ffffffffffffffff81111561269657600080fd5b6126a286828701612370565b90945092506126b59050602085016123b9565b90509250925092565b600080600080600080600060c0888a0312156126d957600080fd5b8735965060208801356126eb8161234b565b955060408801356126fb8161234b565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561272557600080fd5b6127318a828b01612370565b989b979a50959850939692959293505050565b60008060008060006080868803121561275c57600080fd5b85356127678161234b565b945060208601359350604086013567ffffffffffffffff81111561278a57600080fd5b61279688828901612370565b90945092506127a99050606087016123b9565b90509295509295909350565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a083015261280760c0830184866125b0565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff861681526080602082015260006128446080830186886125b0565b905083604083015263ffffffff831660608301529695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff808316818516818304811182151516156128b8576128b8612862565b02949350505050565b600067ffffffffffffffff80841680612903577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff80831681851680830382111561293257612932612862565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60006020828403121561297c57600080fd5b81518015158114610e6557600080fd5b60008282101561299e5761299e612862565b500390565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015283604082015267ffffffffffffffff8316606082015260a0608082015260006122ab60a0830184612511565b600060208284031215612a0057600080fd5b8151610e658161234b565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152612a446080830185612511565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a083015261264560c0830184612511565b60008251612ab28184602087016124e5565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
}
