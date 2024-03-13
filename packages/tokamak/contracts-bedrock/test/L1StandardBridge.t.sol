// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
import { stdStorage, StdStorage } from "forge-std/Test.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { Bridge_Initializer } from "test/CommonTest.t.sol";

// Libraries
import { Predeploys } from "src/libraries/Predeploys.sol";
import { Constants } from "src/libraries/Constants.sol";

// Target contract dependencies
import { StandardBridge } from "src/universal/StandardBridge.sol";
import { L2StandardBridge } from "src/L2/L2StandardBridge.sol";
import { CrossDomainMessenger } from "src/universal/CrossDomainMessenger.sol";
import { L1CrossDomainMessenger } from "src/L1/L1CrossDomainMessenger.sol";
import { AddressAliasHelper } from "src/vendor/AddressAliasHelper.sol";

// Target contract
import { OptimismPortal } from "src/L1/OptimismPortal.sol";

contract L1StandardBridge_Getter_Test is Bridge_Initializer {
    /// @dev Test that the accessors return the correct initialized values.
    function test_getters_succeeds() external view {
        assert(L1Bridge.l2TokenBridge() == address(L2Bridge));
        assert(L1Bridge.OTHER_BRIDGE() == L2Bridge);
        assert(L1Bridge.messenger() == L1Messenger);
        assert(L1Bridge.MESSENGER() == L1Messenger);
    }
}

contract L1StandardBridge_Initialize_Test is Bridge_Initializer {
    /// @dev Test that the initialize function sets the correct values.
    function test_initialize_succeeds() external {
        assertEq(address(L1Bridge.messenger()), address(L1Messenger));
        assertEq(address(L1Bridge.OTHER_BRIDGE()), Predeploys.L2_STANDARD_BRIDGE);
        assertEq(address(L2Bridge), Predeploys.L2_STANDARD_BRIDGE);
        bytes32 slot0 = vm.load(address(L1Bridge), bytes32(uint256(0)));
        assertEq(slot0, bytes32(uint256(Constants.INITIALIZER)));
    }
}

contract L1StandardBridge_Initialize_TestFail is Bridge_Initializer { }

contract L1StandardBridge_Receive_Test is Bridge_Initializer {
    /// @dev Tests receive bridges ETH successfully.
    function test_receive_succeeds() external {
        assertEq(address(L1Bridge).balance, 0);

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHBridgeInitiated(alice, alice, 100, hex"");

        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(
                CrossDomainMessenger.sendMessage.selector,
                address(L2Bridge),
                abi.encodeWithSelector(
                    StandardBridge.finalizeBridgeERC20.selector, Predeploys.ETH, address(0), alice, alice, 100, hex""
                ),
                200_000
            )
        );

        vm.prank(alice, alice);
        (bool success,) = address(L1Bridge).call{ value: 100 }(hex"");
        assertEq(success, true);
        assertEq(address(L1Bridge).balance, 100);
    }
}

contract L1StandardBridge_Receive_TestFail { }

contract PreBridgeETH is Bridge_Initializer {
    /// @dev Asserts the expected calls and events for bridging ETH depending
    ///      on whether the bridge call is legacy or not.
    function _preBridgeETH(bool isLegacy) internal {
        assertEq(address(L1Bridge).balance, 0);
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, Predeploys.ETH, address(0), alice, alice, 500, hex"dead"
        );

        if (isLegacy) {
            vm.expectCall(
                address(L1Bridge), 500, abi.encodeWithSelector(L1Bridge.depositETH.selector, 50000, hex"dead")
            );
        } else {
            vm.expectCall(address(L1Bridge), 500, abi.encodeWithSelector(L1Bridge.bridgeETH.selector, 50000, hex"dead"));
        }
        vm.expectCall(
            address(L1Messenger),
            0,
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(L2Bridge), message, 50000)
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(L1Bridge), address(L2Bridge), 0, 50000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 50000);
        vm.expectCall(
            address(op),
            0,
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(L2Messenger), 0, baseGas, innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(0), uint256(0), baseGas, false, innerMessage);

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHBridgeInitiated(alice, alice, 500, hex"dead");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(L2Bridge), address(L1Bridge), message, nonce, 50000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(L1Bridge), 0);

        vm.prank(alice, alice);
    }
}

contract L1StandardBridge_DepositETH_Test is PreBridgeETH {
    /// @dev Tests that depositing ETH succeeds.
    ///      Emits ETHDepositInitiated and ETHBridgeInitiated events.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Only EOA can call depositETH.
    ///      ETH ends up in the optimismPortal.
    function test_depositETH_succeeds() external {
        _preBridgeETH({ isLegacy: true });
        L1Bridge.depositETH{ value: 500 }(50000, hex"dead");
        assertEq(address(L1Bridge).balance, 500);
    }
}

contract L1StandardBridge_BridgeETH_Test is PreBridgeETH {
    /// @dev Tests that bridging ETH succeeds.
    ///      Emits ETHDepositInitiated and ETHBridgeInitiated events.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Only EOA can call bridgeETH.
    ///      ETH ends up in the optimismPortal.
    function test_bridgeETH_succeeds() external {
        _preBridgeETH({ isLegacy: false });
        L1Bridge.bridgeETH{ value: 500 }(50000, hex"dead");
        assertEq(address(L1Bridge).balance, 500);
    }
}

contract L1StandardBridge_DepositETH_TestFail is Bridge_Initializer {
    /// @dev Tests that depositing ETH reverts if the call is not from an EOA.
    function test_depositETH_notEoa_reverts() external {
        vm.etch(alice, address(L1Token).code);
        vm.expectRevert("StandardBridge: function can only be called from an EOA");
        vm.prank(alice);
        L1Bridge.depositETH{ value: 1 }(300, hex"");
    }
}

contract PreBridgeETHTo is Bridge_Initializer {
    /// @dev Asserts the expected calls and events for bridging ETH to a different
    ///      address depending on whether the bridge call is legacy or not.
    function _preBridgeETHTo(bool isLegacy) internal {
        assertEq(address(op).balance, 0);
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));

        if (isLegacy) {
            vm.expectCall(
                address(L1Bridge), 600, abi.encodeWithSelector(L1Bridge.depositETHTo.selector, bob, 60000, hex"dead")
            );
        } else {
            vm.expectCall(
                address(L1Bridge), 600, abi.encodeWithSelector(L1Bridge.bridgeETHTo.selector, bob, 60000, hex"dead")
            );
        }

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, Predeploys.ETH, address(0), alice, bob, 600, hex"dead"
        );

        // the L1 bridge should call
        // L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(L2Bridge), message, 60000)
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(L1Bridge), address(L2Bridge), 0, 60000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 60000);
        vm.expectCall(
            address(op),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(L2Messenger), 0, baseGas, innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(0), uint256(0), baseGas, false, innerMessage);

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHBridgeInitiated(alice, bob, 600, hex"dead");

        // // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(L2Bridge), address(L1Bridge), message, nonce, 60000);

        // // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(L1Bridge), 0);

        // deposit eth to bob
        vm.prank(alice, alice);
    }
}

contract L1StandardBridge_DepositETHTo_Test is PreBridgeETHTo {
    /// @dev Tests that depositing ETH to a different address succeeds.
    ///      Emits ETHDepositInitiated event.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      EOA or contract can call depositETHTo.
    ///      ETH ends up in the optimismPortal.
    function test_depositETHTo_succeeds() external {
        _preBridgeETHTo({ isLegacy: true });
        L1Bridge.depositETHTo{ value: 600 }(bob, 60000, hex"dead");
        assertEq(address(L1Bridge).balance, 600);
    }
}

contract L1StandardBridge_BridgeETHTo_Test is PreBridgeETHTo {
    /// @dev Tests that bridging ETH to a different address succeeds.
    ///      Emits ETHDepositInitiated and ETHBridgeInitiated events.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Only EOA can call bridgeETHTo.
    ///      ETH ends up in the optimismPortal.
    function test_bridgeETHTo_succeeds() external {
        _preBridgeETHTo({ isLegacy: false });
        L1Bridge.bridgeETHTo{ value: 600 }(bob, 60000, hex"dead");
        assertEq(address(L1Bridge).balance, 600);
    }
}

contract L1StandardBridge_DepositETHTo_TestFail is Bridge_Initializer { }

contract L1StandardBridge_DepositERC20_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    // depositERC20
    // - updates bridge.deposits
    // - emits ERC20DepositInitiated
    // - calls optimismPortal.depositTransaction
    // - only callable by EOA

    /// @dev Tests that depositing ERC20 to the bridge succeeds.
    ///      Bridge deposits are updated.
    ///      Emits ERC20DepositInitiated event.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Only EOA can call depositERC20.
    function test_depositERC20_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));

        // Deal Alice's ERC20 State
        deal(address(L1Token), alice, 100000, true);
        vm.prank(alice);
        L1Token.approve(address(L1Bridge), type(uint256).max);

        // The L1Bridge should transfer alice's tokens to itself
        vm.expectCall(
            address(L1Token), abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(L1Bridge), 100)
        );

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(L2Token), address(L1Token), alice, alice, 100, hex""
        );

        // the L1 bridge should call L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(L2Bridge), message, 10000)
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(L1Bridge), address(L2Bridge), 0, 10000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 10000);
        vm.expectCall(
            address(op),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(L2Messenger), 0, baseGas, innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(0), uint256(0), baseGas, false, innerMessage);

        // Should emit both the bedrock and legacy events
        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20DepositInitiated(address(L1Token), address(L2Token), alice, alice, 100, hex"");

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20BridgeInitiated(address(L1Token), address(L2Token), alice, alice, 100, hex"");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(L2Bridge), address(L1Bridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(L1Bridge), 0);

        vm.prank(alice);
        L1Bridge.depositERC20(address(L1Token), address(L2Token), 100, 10000, hex"");
        assertEq(L1Bridge.deposits(address(L1Token), address(L2Token)), 100);
    }
}

contract L1StandardBridge_DepositERC20_TestFail is Bridge_Initializer {
    /// @dev Tests that depositing an ERC20 to the bridge reverts
    ///      if the caller is not an EOA.
    function test_depositERC20_notEoa_reverts() external {
        // turn alice into a contract
        vm.etch(alice, hex"ffff");

        vm.expectRevert("StandardBridge: function can only be called from an EOA");
        vm.prank(alice, alice);
        L1Bridge.depositERC20(address(0), address(0), 100, 100, hex"");
    }
}

contract L1StandardBridge_DepositERC20To_Test is Bridge_Initializer {
    /// @dev Tests that depositing ERC20 to the bridge succeeds when
    ///      sent to a different address.
    ///      Bridge deposits are updated.
    ///      Emits ERC20DepositInitiated event.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Contracts can call depositERC20.
    function test_depositERC20To_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));

        bytes memory message = abi.encodeWithSelector(
            StandardBridge.finalizeBridgeERC20.selector, address(L2Token), address(L1Token), alice, bob, 1000, hex""
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(L1Bridge), address(L2Bridge), 0, 10000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 10000);
        bytes memory opaqueData = abi.encodePacked(uint256(0), uint256(0), baseGas, false, innerMessage);

        deal(address(L1Token), alice, 100000, true);

        vm.prank(alice);
        L1Token.approve(address(L1Bridge), type(uint256).max);

        // Should emit both the bedrock and legacy events
        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20DepositInitiated(address(L1Token), address(L2Token), alice, bob, 1000, hex"");

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20BridgeInitiated(address(L1Token), address(L2Token), alice, bob, 1000, hex"");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(L2Bridge), address(L1Bridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(L1Bridge), 0);

        // the L1 bridge should call L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(CrossDomainMessenger.sendMessage.selector, address(L2Bridge), message, 10000)
        );
        // The L1 XDM should call OptimismPortal.depositTransaction
        vm.expectCall(
            address(op),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(L2Messenger), 0, baseGas, innerMessage
            )
        );
        vm.expectCall(
            address(L1Token), abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(L1Bridge), 1000)
        );

        vm.prank(alice);
        L1Bridge.depositERC20To(address(L1Token), address(L2Token), bob, 1000, 10000, hex"");

        assertEq(L1Bridge.deposits(address(L1Token), address(L2Token)), 1000);
    }
}

contract L1StandardBridge_DepositNativeToken_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    // deposit to L2's native tokens
    // - updates bridge.deposits
    // - emits ERC20DepositInitiated
    // - calls optimismPortal.depositTransaction
    // - only callable by EOA

    /// @dev Tests that depositing L2's native token to the bridge succeeds.
    ///      Bridge deposits are updated.
    ///      Emits ERC20DepositInitiated event.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Only EOA can call depositNativeToken.
    function test_depositNativeToken_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));
        ERC20 nativeToken = ERC20(L1Bridge.nativeTokenAddress());
        // Deal Alice's ERC20 State
        deal(address(L1Bridge.nativeTokenAddress()), alice, 100000, true);
        vm.prank(alice, alice);
        nativeToken.approve(address(L1Bridge), type(uint256).max);

        // The L1Bridge should transfer alice's tokens to itself
        vm.expectCall(
            L1Bridge.nativeTokenAddress(),
            abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(L1Bridge), 100)
        );

        bytes memory message =
            abi.encodeWithSelector(StandardBridge.finalizeBridgeETH.selector, alice, alice, 100, hex"");

        // the L1 bridge should call L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(
                L1CrossDomainMessenger.sendNativeTokenMessage.selector, address(L2Bridge), 100, message, 10000
            )
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(L1Bridge), address(L2Bridge), 100, 10000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 10000);
        vm.expectCall(
            address(op),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(L2Messenger), 100, baseGas, innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(100), uint256(100), baseGas, false, innerMessage);

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20DepositInitiated(address(nativeToken), Predeploys.LEGACY_ERC20_ETH, alice, alice, 100, hex"");

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20BridgeInitiated(address(nativeToken), Predeploys.LEGACY_ERC20_ETH, alice, alice, 100, hex"");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(L2Bridge), address(L1Bridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(L1Bridge), 100);

        vm.prank(alice, alice);
        L1Bridge.depositNativeToken(100, 10000, hex"");
    }
}

contract L1StandardBridge_DepositNativeToken_TestFail is Bridge_Initializer {
    /// @dev Tests that depositing native token to the bridge reverts
    ///      if the caller is not an EOA.
    function test_depositNativeToken_notEoa_reverts() external {
        // turn alice into a contract
        vm.etch(alice, hex"ffff");

        vm.expectRevert("StandardBridge: function can only be called from an EOA");
        vm.prank(alice, alice);
        L1Bridge.depositNativeToken(100, 100, hex"");
    }
}

contract L1StandardBridge_DepositNativeTokenTo_Test is Bridge_Initializer {
    /// @dev Tests that depositing Nativetoken to the bridge succeeds when
    ///      sent to a different address.
    ///      Bridge deposits are updated.
    ///      Emits ERC20DepositInitiated event.
    ///      Calls depositTransaction on the OptimismPortal.
    ///      Contracts can call depositNativeTokenTo.
    function test_depositNativeTokenTo_succeeds() external {
        uint256 nonce = L1Messenger.messageNonce();
        uint256 version = 0; // Internal constant in the OptimismPortal: DEPOSIT_VERSION
        address l1MessengerAliased = AddressAliasHelper.applyL1ToL2Alias(address(L1Messenger));
        ERC20 nativeToken = ERC20(L1Bridge.nativeTokenAddress());
        // Deal Alice's ERC20 State
        deal(address(L1Bridge.nativeTokenAddress()), alice, 100000, true);
        vm.prank(alice, alice);
        nativeToken.approve(address(L1Bridge), type(uint256).max);

        // The L1Bridge should transfer alice's tokens to itself
        vm.expectCall(
            L1Bridge.nativeTokenAddress(),
            abi.encodeWithSelector(ERC20.transferFrom.selector, alice, address(L1Bridge), 100)
        );

        bytes memory message = abi.encodeWithSelector(StandardBridge.finalizeBridgeETH.selector, alice, bob, 100, hex"");

        // the L1 bridge should call L1CrossDomainMessenger.sendMessage
        vm.expectCall(
            address(L1Messenger),
            abi.encodeWithSelector(
                L1CrossDomainMessenger.sendNativeTokenMessage.selector, address(L2Bridge), 100, message, 10000
            )
        );

        bytes memory innerMessage = abi.encodeWithSelector(
            CrossDomainMessenger.relayMessage.selector, nonce, address(L1Bridge), address(L2Bridge), 100, 10000, message
        );

        uint64 baseGas = L1Messenger.baseGas(message, 10000);
        vm.expectCall(
            address(op),
            abi.encodeWithSelector(
                OptimismPortal.depositTransaction.selector, address(L2Messenger), 100, baseGas, innerMessage
            )
        );

        bytes memory opaqueData = abi.encodePacked(uint256(100), uint256(100), baseGas, false, innerMessage);

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20DepositInitiated(address(nativeToken), Predeploys.LEGACY_ERC20_ETH, alice, bob, 100, hex"");

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20BridgeInitiated(address(nativeToken), Predeploys.LEGACY_ERC20_ETH, alice, bob, 100, hex"");

        // OptimismPortal emits a TransactionDeposited event on `depositTransaction` call
        vm.expectEmit(true, true, true, true, address(op));
        emit TransactionDeposited(l1MessengerAliased, address(L2Messenger), version, opaqueData);

        // SentMessage event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessage(address(L2Bridge), address(L1Bridge), message, nonce, 10000);

        // SentMessageExtension1 event emitted by the CrossDomainMessenger
        vm.expectEmit(true, true, true, true, address(L1Messenger));
        emit SentMessageExtension1(address(L1Bridge), 100);

        vm.prank(alice, alice);
        L1Bridge.depositNativeTokenTo(bob, 100, 10000, hex"");
    }
}

contract L1StandardBridge_FinalizeETHWithdrawal_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    /// @dev Tests that finalizing an ETH withdrawal succeeds.
    ///      Emits ETHWithdrawalFinalized event.
    ///      Only callable by the L2 bridge.
    function test_finalizeETHWithdrawal_succeeds() external {
        uint256 aliceBalance = alice.balance;

        uint256 depositSlotIndex = 2;
        bytes32 innerSlotIndex = keccak256(abi.encodePacked(uint256(0), depositSlotIndex));
        bytes32 depositAmountSlotIndex =
            keccak256(abi.encodePacked(uint256(uint160(Predeploys.ETH)), uint256(innerSlotIndex)));

        vm.store(address(L1Bridge), bytes32(depositAmountSlotIndex), bytes32(uint256(100)));

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHWithdrawalFinalized(alice, alice, 100, hex"");

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ETHBridgeFinalized(alice, alice, 100, hex"");

        vm.expectCall(alice, hex"");

        vm.mockCall(
            address(L1Bridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        // ensure that the messenger has ETH to call with
        vm.deal(address(L1Bridge), 100);
        vm.prank(address(L1Bridge.messenger()));
        L1Bridge.finalizeBridgeERC20(address(0), Predeploys.ETH, alice, alice, 100, hex"");

        assertEq(aliceBalance + 100, alice.balance);
    }

    /// @dev Tests that finalizing an ETH withdrawal revert: not messenger
    function test_finalizeETHWithdrawal_notMessenger_reverts() external {
        vm.mockCall(
            address(L1Bridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.prank(address(28));
        vm.expectRevert("StandardBridge: function can only be called from the other bridge");
        L1Bridge.finalizeBridgeERC20(address(0), Predeploys.ETH, alice, alice, 100, hex"");
    }

    /// @dev Tests that finalizing an ETH withdrawal revert: not other bridge
    function test_finalizeETHWithdrawal_notOtherBridge_reverts() external {
        vm.mockCall(
            address(L1Bridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(0))
        );
        vm.prank(address(L1Bridge.messenger()));
        vm.expectRevert("StandardBridge: function can only be called from the other bridge");
        L1Bridge.finalizeBridgeERC20(address(0), Predeploys.ETH, alice, alice, 100, hex"");
    }
}

contract L1StandardBridge_FinalizeETHWithdrawal_TestFail is Bridge_Initializer { }

contract L1StandardBridge_FinalizeERC20Withdrawal_Test is Bridge_Initializer {
    using stdStorage for StdStorage;

    /// @dev Tests that finalizing an ERC20 withdrawal succeeds.
    ///      Bridge deposits are updated.
    ///      Emits ERC20WithdrawalFinalized event.
    ///      Only callable by the L2 bridge.
    function test_finalizeERC20Withdrawal_succeeds() external {
        deal(address(L1Token), address(L1Bridge), 100, true);

        uint256 slot = stdstore.target(address(L1Bridge)).sig("deposits(address,address)").with_key(address(L1Token))
            .with_key(address(L2Token)).find();

        // Give the L1 bridge some ERC20 tokens
        vm.store(address(L1Bridge), bytes32(slot), bytes32(uint256(100)));
        assertEq(L1Bridge.deposits(address(L1Token), address(L2Token)), 100);

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20WithdrawalFinalized(address(L1Token), address(L2Token), alice, alice, 100, hex"");

        vm.expectEmit(true, true, true, true, address(L1Bridge));
        emit ERC20BridgeFinalized(address(L1Token), address(L2Token), alice, alice, 100, hex"");

        vm.expectCall(address(L1Token), abi.encodeWithSelector(ERC20.transfer.selector, alice, 100));

        vm.mockCall(
            address(L1Bridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.prank(address(L1Bridge.messenger()));
        L1Bridge.finalizeERC20Withdrawal(address(L1Token), address(L2Token), alice, alice, 100, hex"");

        assertEq(L1Token.balanceOf(address(L1Bridge)), 0);
        assertEq(L1Token.balanceOf(address(alice)), 100);
    }
}

contract L1StandardBridge_FinalizeERC20Withdrawal_TestFail is Bridge_Initializer {
    /// @dev Tests that finalizing an ERC20 withdrawal reverts if the caller is not the L2 bridge.
    function test_finalizeERC20Withdrawal_notMessenger_reverts() external {
        vm.mockCall(
            address(L1Bridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.prank(address(28));
        vm.expectRevert("StandardBridge: function can only be called from the other bridge");
        L1Bridge.finalizeERC20Withdrawal(address(L1Token), address(L2Token), alice, alice, 100, hex"");
    }

    /// @dev Tests that finalizing an ERC20 withdrawal reverts if the caller is not the L2 bridge.
    function test_finalizeERC20Withdrawal_notOtherBridge_reverts() external {
        vm.mockCall(
            address(L1Bridge.messenger()),
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(address(0)))
        );
        vm.prank(address(L1Bridge.messenger()));
        vm.expectRevert("StandardBridge: function can only be called from the other bridge");
        L1Bridge.finalizeERC20Withdrawal(address(L1Token), address(L2Token), alice, alice, 100, hex"");
    }
}

contract L1StandardBridge_FinalizeBridgeTON_Test is Bridge_Initializer {
    /// @dev Tests that finalizing bridged TON succeeds.
    function test_finalizeBridgeTON_succeeds() external {
        ERC20 nativeToken = ERC20(L1Bridge.nativeTokenAddress());
        address messenger = address(L1Bridge.messenger());
        uint256 beforeBalance = nativeToken.balanceOf(address(alice));

        uint256 depositSlotIndex = 2;
        bytes32 innerSlotIndex = keccak256(abi.encodePacked(uint256(uint160(address(nativeToken))), depositSlotIndex));
        bytes32 depositAmountSlotIndex =
            keccak256(abi.encodePacked(uint256(uint160(Predeploys.LEGACY_ERC20_ETH)), uint256(innerSlotIndex)));

        vm.store(address(L1Bridge), bytes32(depositAmountSlotIndex), bytes32(uint256(100)));
        deal(L1Bridge.nativeTokenAddress(), messenger, 100, true);
        vm.prank(messenger);
        nativeToken.approve(address(L1Bridge), 100);

        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.prank(messenger);
        L1Bridge.finalizeBridgeETH(alice, alice, 100, hex"");
        uint256 afterBalance = nativeToken.balanceOf(address(alice));
        assertEq(beforeBalance + 100, afterBalance);
    }
}

contract L1StandardBridge_FinalizeBridgeTON_TestFail is Bridge_Initializer {
    /// @dev Tests that finalizing bridged TON reverts if the amount is incorrect.
    function test_finalizeBridgeTON_incorrectValue_reverts() external {
        ERC20 nativeToken = ERC20(L1Bridge.nativeTokenAddress());
        address messenger = address(L1Bridge.messenger());

        uint256 depositSlotIndex = 2;
        bytes32 innerSlotIndex = keccak256(abi.encodePacked(uint256(uint160(address(nativeToken))), depositSlotIndex));
        bytes32 depositAmountSlotIndex =
            keccak256(abi.encodePacked(uint256(uint160(Predeploys.LEGACY_ERC20_ETH)), uint256(innerSlotIndex)));

        vm.store(address(L1Bridge), bytes32(depositAmountSlotIndex), bytes32(uint256(100)));
        deal(L1Bridge.nativeTokenAddress(), messenger, 100, true);
        vm.prank(messenger);
        nativeToken.approve(address(L1Bridge), 50);

        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.prank(messenger);
        vm.expectRevert("ERC20: insufficient allowance");
        L1Bridge.finalizeBridgeETH(alice, alice, 100, hex"");
    }

    /// @dev Tests that finalizing bridged ETH reverts if the destination is the L1 bridge.
    function test_finalizeBridgeTON_sendToSelf_reverts() external {
        address messenger = address(L1Bridge.messenger());

        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.prank(messenger);
        vm.expectRevert("StandardBridge: cannot send to self");
        L1Bridge.finalizeBridgeETH(alice, address(L1Bridge), 100, hex"");
    }

    /// @dev Tests that finalizing bridged ETH reverts if the destination is the messenger.
    function test_finalizeBridgeTON_sendToMessenger_reverts() external {
        address messenger = address(L1Bridge.messenger());

        vm.mockCall(
            messenger,
            abi.encodeWithSelector(CrossDomainMessenger.xDomainMessageSender.selector),
            abi.encode(address(L1Bridge.OTHER_BRIDGE()))
        );
        vm.prank(messenger);
        vm.expectRevert("StandardBridge: cannot send to messenger");
        L1Bridge.finalizeBridgeETH(alice, messenger, 100, hex"");
    }
}
