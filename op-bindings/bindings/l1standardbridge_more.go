// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"spacer_0_2_20\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1005,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_contract(CrossDomainMessenger)1008\"},{\"astId\":1006,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_array(t_uint256)46_storage\"},{\"astId\":1007,\"contract\":\"src/L1/L1StandardBridge.sol:L1StandardBridge\",\"label\":\"nativeTokenAddress\",\"offset\":0,\"slot\":\"50\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)46_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(CrossDomainMessenger)1008\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1StandardBridgeStorageLayout = new(solc.StorageLayout)

var L1StandardBridgeDeployedBin = "0x6080604052600436106101bb5760003560e01c806387087623116100ec578063a9f9e6751161008a578063db6341cc11610064578063db6341cc1461064a578063e11013dd1461066a578063fa78ad2c1461064a578063fcebb7eb1461055857600080fd5b8063a9f9e67514610617578063b1a1a88214610637578063c89701a21461057857600080fd5b806391c49bf8116100c657806391c49bf814610578578063927ede2d146105ab57806392a162cf146105d65780639a2ac6d51461060457600080fd5b806387087623146104f25780638f601f66146105125780638f9b352c1461055857600080fd5b8063485cc9551161015957806354fd4d501161013357806354fd4d501461042857806358a997f61461047e5780637f46ddb21461049e578063838b2520146104d257600080fd5b8063485cc955146103bb5780634d0047ee146103db578063540abf731461040857600080fd5b80631532ec34116101955780631532ec34146103235780631635f5fd146103365780633cb747bf146103495780634273ca161461039b57600080fd5b80630166a07a1461027957806301ffc9a71461029957806309fc88431461031057600080fd5b3661027457333b15610254576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b610272333362030d406040518060200160405280600081525061067d565b005b600080fd5b34801561028557600080fd5b50610272610294366004612bc6565b610690565b3480156102a557600080fd5b506102fb6102b4366004612c5e565b7fffffffff00000000000000000000000000000000000000000000000000000000167f4273ca16000000000000000000000000000000000000000000000000000000001490565b60405190151581526020015b60405180910390f35b61027261031e366004612cb9565b610b82565b610272610331366004612d0c565b610c59565b610272610344366004612d0c565b610c6d565b34801561035557600080fd5b506003546103769073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610307565b3480156103a757600080fd5b506102fb6103b6366004612e59565b61106a565b3480156103c757600080fd5b506102726103d6366004612ec5565b611118565b3480156103e757600080fd5b506032546103769073ffffffffffffffffffffffffffffffffffffffff1681565b34801561041457600080fd5b50610272610423366004612efe565b611392565b34801561043457600080fd5b506104716040518060400160405280600581526020017f312e342e3100000000000000000000000000000000000000000000000000000081525081565b6040516103079190612feb565b34801561048a57600080fd5b50610272610499366004612ffe565b6113d7565b3480156104aa57600080fd5b506103767f000000000000000000000000000000000000000000000000000000000000000081565b3480156104de57600080fd5b506102726104ed366004612efe565b6114b3565b3480156104fe57600080fd5b5061027261050d366004612ffe565b6114f8565b34801561051e57600080fd5b5061054a61052d366004612ec5565b600260209081526000928352604080842090915290825290205481565b604051908152602001610307565b34801561056457600080fd5b50610272610573366004613081565b6115cc565b34801561058457600080fd5b507f0000000000000000000000000000000000000000000000000000000000000000610376565b3480156105b757600080fd5b5060035473ffffffffffffffffffffffffffffffffffffffff16610376565b3480156105e257600080fd5b506105f66105f13660046130d5565b61160f565b60405161030792919061310a565b610272610612366004613129565b611647565b34801561062357600080fd5b50610272610632366004612bc6565b611689565b610272610645366004612cb9565b611698565b34801561065657600080fd5b5061027261066536600461318c565b611769565b610272610678366004613129565b61183b565b61068a848434858561187e565b50505050565b60035473ffffffffffffffffffffffffffffffffffffffff163314801561077f5750600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa158015610743573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076791906131b2565b73ffffffffffffffffffffffffffffffffffffffff16145b610831576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a40161024b565b61083a87611a7b565b8015610861575060325473ffffffffffffffffffffffffffffffffffffffff888116911614155b156109af576108708787611add565b610922576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a40161024b565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b15801561099257600080fd5b505af11580156109a6573d6000803e3d6000fd5b50505050610b35565b73ffffffffffffffffffffffffffffffffffffffff87161580156109fc575073ffffffffffffffffffffffffffffffffffffffff8616734200000000000000000000000000000000000486145b15610ab3576000610a1e855a8660405180602001604052806000815250611bfd565b905080610aad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c65640000000000000000000000000000000000000000000000000000000000606482015260840161024b565b50610b35565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a1683529290522054610af19084906131fe565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c1683529390529190912091909155610b35908585611c17565b610b79878787878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611ceb92505050565b50505050505050565b333b15610c11576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161024b565b610c543333348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061187e92505050565b505050565b610c668585858585610c6d565b5050505050565b60035473ffffffffffffffffffffffffffffffffffffffff1633148015610d5c5750600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa158015610d20573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d4491906131b2565b73ffffffffffffffffffffffffffffffffffffffff16145b610e0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a40161024b565b3415610e9c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206d757374206e6f74207265636569766560448201527f204574686572206576656e20696620746869732069732070617961626c650000606482015260840161024b565b3073ffffffffffffffffffffffffffffffffffffffff851603610f41576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c660000000000000000000000000000000000000000000000000000000000606482015260840161024b565b60035473ffffffffffffffffffffffffffffffffffffffff90811690851603610fec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e676572000000000000000000000000000000000000000000000000606482015260840161024b565b61102e85858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611d7992505050565b603254610c669073ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead00008787878787610690565b60325460009073ffffffffffffffffffffffffffffffffffffffff1633146110ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f6f6e6c792061636365707420544f4e20617070726f76652063616c6c6261636b604482015260640161024b565b6000806110fa8461160f565b9150915061110b8788878585611dec565b5060019695505050505050565b600054600390610100900460ff1615801561113a575060005460ff8083169116105b6111c6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161024b565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff83161761010017905561120083611fd6565b603280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8481169190911790915583161580159061126c575060325473ffffffffffffffffffffffffffffffffffffffff1615155b1561132f576032546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60248301529091169063095ea7b3906044016020604051808303816000875af1158015611309573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061132d9190613215565b505b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b610b7987873388888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506120b492505050565b333b15611466576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161024b565b6114ab86863333888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506120f792505050565b505050505050565b610b7987873388888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506120f792505050565b333b15611587576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161024b565b6114ab86863333888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506120b492505050565b610c663386868686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611dec92505050565b6000606060048351101561163457505060408051602081019091526000815262030d40905b5050602081015160e01c91602490910190565b61068a33858585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061067d92505050565b610b7987878787878787610690565b333b15611727576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161024b565b610c5433338585858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061067d92505050565b333b156117f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f41000000000000000000606482015260840161024b565b61068a3333868686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611dec92505050565b61068a3385348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061187e92505050565b82341461190d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c75650000606482015260840161024b565b61191985858584612106565b60035460405173ffffffffffffffffffffffffffffffffffffffff90911690633dbb202b907f0000000000000000000000000000000000000000000000000000000000000000907f0166a07a00000000000000000000000000000000000000000000000000000000906119af90734200000000000000000000000000000000000486906000908c908c908c908b90602401613237565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b9092168252611a4292918790600401613292565b600060405180830381600087803b158015611a5c57600080fd5b505af1158015611a70573d6000803e3d6000fd5b505050505050505050565b6000611aa7827f1d1d8b6300000000000000000000000000000000000000000000000000000000612173565b80611ad75750611ad7827fec4fc8e300000000000000000000000000000000000000000000000000000000612173565b92915050565b6000611b09837f1d1d8b6300000000000000000000000000000000000000000000000000000000612173565b15611bb2578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b59573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b7d91906131b2565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050611ad7565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b59573d6000803e3d6000fd5b600080600080845160208601878a8af19695505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610c549084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152612196565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3868686604051611d63939291906132d7565b60405180910390a46114ab8686868686866122a2565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e6318484604051611dd8929190613315565b60405180910390a361068a8484848461232a565b603254611e119073ffffffffffffffffffffffffffffffffffffffff16863086612389565b60325473ffffffffffffffffffffffffffffffffffffffff16600090815260026020908152604080832073deaddeaddeaddeaddeaddeaddeaddeaddead00008452909152902054611e6390849061332e565b6032805473ffffffffffffffffffffffffffffffffffffffff908116600090815260026020908152604080832073deaddeaddeaddeaddeaddeaddeaddeaddead00008085529252909120939093559054611ec392911690878787866123e7565b60035460405173ffffffffffffffffffffffffffffffffffffffff9091169063cb8398bf907f00000000000000000000000000000000000000000000000000000000000000009086907f1635f5fd0000000000000000000000000000000000000000000000000000000090611f42908b908b9085908a90602401613346565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b9092168252611a42939291889060040161338f565b600054610100900460ff1661206d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161024b565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60325473ffffffffffffffffffffffffffffffffffffffff8089169116036120e8576120e38585858585611dec565b610b79565b610b7987878787878787612475565b610b79878787878787876120b4565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af58484604051612165929190613315565b60405180910390a350505050565b600061217e836127bd565b801561218f575061218f8383612821565b9392505050565b60006121f8826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166128f09092919063ffffffff16565b805190915015610c5457808060200190518101906122169190613215565b610c54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f74207375636365656400000000000000000000000000000000000000000000606482015260840161024b565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd86868660405161231a939291906132d7565b60405180910390a4505050505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d8484604051612165929190613315565b60405173ffffffffffffffffffffffffffffffffffffffff8085166024830152831660448201526064810182905261068a9085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611c69565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d039686868660405161245f939291906132d7565b60405180910390a46114ab868686868686612907565b61247e87611a7b565b156125cc5761248d8787611add565b61253f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a40161024b565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b1580156125af57600080fd5b505af11580156125c3573d6000803e3d6000fd5b50505050612660565b6125ee73ffffffffffffffffffffffffffffffffffffffff8816863086612389565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a168352929052205461262c90849061332e565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b16835292905220555b61266e8787878787866123e7565b60035460405173ffffffffffffffffffffffffffffffffffffffff90911690633dbb202b907f0000000000000000000000000000000000000000000000000000000000000000907f0166a07a00000000000000000000000000000000000000000000000000000000906126ef908b908d908c908c908c908b90602401613237565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b909216825261278292918790600401613292565b600060405180830381600087803b15801561279c57600080fd5b505af11580156127b0573d6000803e3d6000fd5b5050505050505050505050565b60006127e9827f01ffc9a700000000000000000000000000000000000000000000000000000000612821565b8015611ad7575061281a827fffffffff00000000000000000000000000000000000000000000000000000000612821565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d915060005190508280156128d9575060208210155b80156128e55750600081115b979650505050505050565b60606128ff848460008561297f565b949350505050565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf86868660405161231a939291906132d7565b606082471015612a11576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c0000000000000000000000000000000000000000000000000000606482015260840161024b565b73ffffffffffffffffffffffffffffffffffffffff85163b612a8f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161024b565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051612ab891906133db565b60006040518083038185875af1925050503d8060008114612af5576040519150601f19603f3d011682016040523d82523d6000602084013e612afa565b606091505b50915091506128e582828660608315612b1457508161218f565b825115612b245782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024b9190612feb565b73ffffffffffffffffffffffffffffffffffffffff81168114612b7a57600080fd5b50565b60008083601f840112612b8f57600080fd5b50813567ffffffffffffffff811115612ba757600080fd5b602083019150836020828501011115612bbf57600080fd5b9250929050565b600080600080600080600060c0888a031215612be157600080fd5b8735612bec81612b58565b96506020880135612bfc81612b58565b95506040880135612c0c81612b58565b94506060880135612c1c81612b58565b93506080880135925060a088013567ffffffffffffffff811115612c3f57600080fd5b612c4b8a828b01612b7d565b989b979a50959850939692959293505050565b600060208284031215612c7057600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461218f57600080fd5b803563ffffffff81168114612cb457600080fd5b919050565b600080600060408486031215612cce57600080fd5b612cd784612ca0565b9250602084013567ffffffffffffffff811115612cf357600080fd5b612cff86828701612b7d565b9497909650939450505050565b600080600080600060808688031215612d2457600080fd5b8535612d2f81612b58565b94506020860135612d3f81612b58565b935060408601359250606086013567ffffffffffffffff811115612d6257600080fd5b612d6e88828901612b7d565b969995985093965092949392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112612dbf57600080fd5b813567ffffffffffffffff80821115612dda57612dda612d7f565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715612e2057612e20612d7f565b81604052838152866020858801011115612e3957600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060008060808587031215612e6f57600080fd5b8435612e7a81612b58565b93506020850135612e8a81612b58565b925060408501359150606085013567ffffffffffffffff811115612ead57600080fd5b612eb987828801612dae565b91505092959194509250565b60008060408385031215612ed857600080fd5b8235612ee381612b58565b91506020830135612ef381612b58565b809150509250929050565b600080600080600080600060c0888a031215612f1957600080fd5b8735612f2481612b58565b96506020880135612f3481612b58565b95506040880135612f4481612b58565b945060608801359350612f5960808901612ca0565b925060a088013567ffffffffffffffff811115612c3f57600080fd5b60005b83811015612f90578181015183820152602001612f78565b8381111561068a5750506000910152565b60008151808452612fb9816020860160208601612f75565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061218f6020830184612fa1565b60008060008060008060a0878903121561301757600080fd5b863561302281612b58565b9550602087013561303281612b58565b94506040870135935061304760608801612ca0565b9250608087013567ffffffffffffffff81111561306357600080fd5b61306f89828a01612b7d565b979a9699509497509295939492505050565b60008060008060006080868803121561309957600080fd5b85356130a481612b58565b9450602086013593506130b960408701612ca0565b9250606086013567ffffffffffffffff811115612d6257600080fd5b6000602082840312156130e757600080fd5b813567ffffffffffffffff8111156130fe57600080fd5b6128ff84828501612dae565b63ffffffff831681526040602082015260006128ff6040830184612fa1565b6000806000806060858703121561313f57600080fd5b843561314a81612b58565b935061315860208601612ca0565b9250604085013567ffffffffffffffff81111561317457600080fd5b61318087828801612b7d565b95989497509550505050565b600080600080606085870312156131a257600080fd5b8435935061315860208601612ca0565b6000602082840312156131c457600080fd5b815161218f81612b58565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015613210576132106131cf565b500390565b60006020828403121561322757600080fd5b8151801515811461218f57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a083015261328660c0830184612fa1565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006132c16060830185612fa1565b905063ffffffff83166040830152949350505050565b73ffffffffffffffffffffffffffffffffffffffff8416815282602082015260606040820152600061330c6060830184612fa1565b95945050505050565b8281526040602082015260006128ff6040830184612fa1565b60008219821115613341576133416131cf565b500190565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250836040830152608060608301526133856080830184612fa1565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff851681528360208201526080604082015260006133c46080830185612fa1565b905063ffffffff8316606083015295945050505050565b600082516133ed818460208701612f75565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1StandardBridgeStorageLayoutJSON), L1StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1StandardBridge"] = L1StandardBridgeStorageLayout
	deployedBytecodes["L1StandardBridge"] = L1StandardBridgeDeployedBin
}
