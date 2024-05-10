package op_e2e

import (
	"context"
	"math"
	"math/big"
	"testing"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/transactions"
	"github.com/ethereum-optimism/optimism/op-e2e/e2eutils/wait"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/testlog"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

// TestCannotWithdrawTokenWithEmptyMessage: cannot withdraw token on L1
// Success on L2
// Cannot get token on L1
func TestCannotWithdrawTokenWithEmptyMessage(t *testing.T) {
	InitParallel(t)

	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start(t)
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.Clients["l1"]
	l2Client := sys.Clients["sequencer"]

	opts, err := bind.NewKeyedTransactorWithChainID(sys.cfg.Secrets.Alice, cfg.L1ChainIDBig())
	require.Nil(t, err)

	var depositedAmount = big.NewInt(2000000000000000000)

	nativeTokenContract, err := bindings.NewL2NativeToken(cfg.L1Deployments.L2NativeToken, l1Client)
	require.NoError(t, err)

	// faucet NativeToken
	tx, err := nativeTokenContract.Faucet(opts, depositedAmount)
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l1Client, tx.Hash())
	require.NoError(t, err)

	// Approve NativeToken with the CDM
	tx, err = nativeTokenContract.Approve(opts, cfg.L1Deployments.L1CrossDomainMessengerProxy, new(big.Int).SetUint64(math.MaxUint64))
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l1Client, tx.Hash())
	require.NoError(t, err)

	l1CDM, err := bindings.NewL1CrossDomainMessenger(cfg.L1Deployments.L1CrossDomainMessengerProxy, l1Client)
	require.NoError(t, err)

	optimismPortal, err := bindings.NewOptimismPortal(cfg.L1Deployments.OptimismPortalProxy, l1Client)
	require.NoError(t, err)

	l2NativeTokenBalanceBeforeDeposit, err := l2Client.BalanceAt(context.Background(), opts.From, nil)
	require.NoError(t, err)

	l1NativeTokenBalanceBeforeDeposit, err := nativeTokenContract.BalanceOf(&bind.CallOpts{}, opts.From)
	require.NoError(t, err)
	// deposit native token transaction
	depositedTx, err := transactions.PadGasEstimate(opts, 1.1, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return l1CDM.SendNativeTokenMessage(opts, opts.From, depositedAmount, []byte{}, 200000)
	})
	require.NoError(t, err)

	depositDeployedReceipt, err := wait.ForReceiptOK(context.Background(), l1Client, depositedTx.Hash())
	require.NoError(t, err)

	l1NativeTokenBalanceAfterDeposit, err := nativeTokenContract.BalanceOf(&bind.CallOpts{}, opts.From)
	require.NoError(t, err)

	require.Equal(t, l1NativeTokenBalanceBeforeDeposit, l1NativeTokenBalanceAfterDeposit.Add(l1NativeTokenBalanceAfterDeposit, depositedAmount))

	depIt, err := optimismPortal.FilterTransactionDeposited(&bind.FilterOpts{Start: depositDeployedReceipt.BlockNumber.Uint64()}, nil, nil, nil)
	require.NoError(t, err)
	var depositEvent *bindings.OptimismPortalTransactionDeposited
	for depIt.Next() {
		depositEvent = depIt.Event
	}
	require.NotNil(t, depositEvent)

	// Calculate relayed depositTx
	relayedDepositTx, err := derive.UnmarshalDepositLogEvent(&depositEvent.Raw)
	require.NoError(t, err)
	relayedTxReceipt, err := wait.ForReceiptOK(context.Background(), l2Client, types.NewTx(relayedDepositTx).Hash())
	require.NoError(t, err)
	require.Equal(t, relayedTxReceipt.Status, types.ReceiptStatusSuccessful)

	l2NativeTokenBalanceAfterDeposit, err := l2Client.BalanceAt(context.Background(), opts.From, nil)
	require.NoError(t, err)

	require.Equal(t, l2NativeTokenBalanceAfterDeposit, l2NativeTokenBalanceBeforeDeposit.Add(l2NativeTokenBalanceBeforeDeposit, depositedAmount))

	// Start withdraw token by using L2CrossDomainMessenger
	l2CDM, err := bindings.NewL2CrossDomainMessenger(predeploys.L2CrossDomainMessengerAddr, l2Client)
	require.NoError(t, err)

	l2Opts, err := bind.NewKeyedTransactorWithChainID(sys.cfg.Secrets.Alice, cfg.L2ChainIDBig())
	require.NoError(t, err)
	l2Opts.Value = depositedAmount

	withdrawalTxL2, err := transactions.PadGasEstimate(l2Opts, 1.1, func(l2Opts *bind.TransactOpts) (*types.Transaction, error) {
		return l2CDM.SendMessage(l2Opts, opts.From, []byte{}, 200000)
	})
	require.NoError(t, err)

	withdrawalReceipt, err := wait.ForReceiptOK(context.Background(), l2Client, withdrawalTxL2.Hash())
	require.NoError(t, err)
	require.Equal(t, withdrawalReceipt.Status, types.ReceiptStatusSuccessful)

	proveReceipt, finalizedReceipt := ProveAndFinalizeWithdrawal(t, cfg, l1Client, sys.EthInstances["sequencer"], cfg.Secrets.Alice, withdrawalReceipt)
	require.Equal(t, types.ReceiptStatusSuccessful, proveReceipt.Status)
	require.Equal(t, types.ReceiptStatusSuccessful, finalizedReceipt.Status)

	allowance, err := nativeTokenContract.Allowance(&bind.CallOpts{}, cfg.L1Deployments.L1CrossDomainMessenger, opts.From)
	require.NoError(t, err)
	require.Equal(t, uint64(0), allowance.Uint64())
}

// TestDepositWithdrawalSendMessageSuccess: test create a message for withdraw
func TestDepositWithdrawalSendMessageSuccess(t *testing.T) {
	InitParallel(t)

	cfg := DefaultSystemConfig(t)

	sys, err := cfg.Start(t)
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	log := testlog.Logger(t, log.LvlInfo)
	log.Info("genesis", "l2", sys.RollupConfig.Genesis.L2, "l1", sys.RollupConfig.Genesis.L1, "l2_time", sys.RollupConfig.Genesis.L2Time)

	l1Client := sys.Clients["l1"]
	l2Client := sys.Clients["sequencer"]

	opts, err := bind.NewKeyedTransactorWithChainID(sys.cfg.Secrets.Alice, cfg.L1ChainIDBig())
	require.Nil(t, err)

	var amount = big.NewInt(2000)

	nativeTokenContract, err := bindings.NewL2NativeToken(cfg.L1Deployments.L2NativeToken, l1Client)
	require.NoError(t, err)

	// faucet NativeToken
	tx, err := nativeTokenContract.Faucet(opts, amount.Mul(amount, big.NewInt(2)))
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l1Client, tx.Hash())
	require.NoError(t, err)

	// Approve NativeToken with the CDM
	tx, err = nativeTokenContract.Approve(opts, cfg.L1Deployments.L1CrossDomainMessengerProxy, new(big.Int).SetUint64(math.MaxUint64))
	require.NoError(t, err)
	_, err = wait.ForReceiptOK(context.Background(), l1Client, tx.Hash())
	require.NoError(t, err)

	l1CDM, err := bindings.NewL1CrossDomainMessenger(cfg.L1Deployments.L1CrossDomainMessengerProxy, l1Client)
	require.NoError(t, err)

	// deposit native token transaction
	depositedTx, err := transactions.PadGasEstimate(opts, 1.1, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return l1CDM.SendNativeTokenMessage(opts, opts.From, amount, []byte{}, 200000)
	})
	require.NoError(t, err)

	_, err = wait.ForReceiptOK(context.Background(), l1Client, depositedTx.Hash())
	require.NoError(t, err)

	// Start withdraw token by using L2CrossDomainMessenger
	l2CDM, err := bindings.NewL2CrossDomainMessenger(predeploys.L2CrossDomainMessengerAddr, l2Client)
	require.NoError(t, err)

	l2Opts, err := bind.NewKeyedTransactorWithChainID(sys.cfg.Secrets.Alice, cfg.L2ChainIDBig())
	require.NoError(t, err)
	l2Opts.Value = amount

	// Build calldata for the message
	wnativetokenABI, err := bindings.WNativeTokenMetaData.GetAbi()
	require.NoError(t, err)
	calldata, err := wnativetokenABI.Pack("transfer", opts.From, amount)
	require.NoError(t, err)

	withdrawalTxL2, err := transactions.PadGasEstimate(l2Opts, 1.1, func(l2Opts *bind.TransactOpts) (*types.Transaction, error) {
		return l2CDM.SendMessage(l2Opts, cfg.L1Deployments.L2NativeToken, calldata, 20000)
	})
	require.NoError(t, err)

	withdrawalReceipt, err := wait.ForReceiptOK(context.Background(), l2Client, withdrawalTxL2.Hash())
	require.NoError(t, err)
	require.Equal(t, types.ReceiptStatusSuccessful, withdrawalReceipt.Status)

	balanceBeforeFinalization, err := nativeTokenContract.BalanceOf(&bind.CallOpts{}, opts.From)
	require.NoError(t, err)

	proveReceipt, finalizedReceipt := ProveAndFinalizeWithdrawal(t, cfg, l1Client, sys.EthInstances["sequencer"], cfg.Secrets.Alice, withdrawalReceipt)
	require.Equal(t, types.ReceiptStatusSuccessful, proveReceipt.Status)
	require.Equal(t, types.ReceiptStatusSuccessful, finalizedReceipt.Status)

	balanceAfterFinalization, err := nativeTokenContract.BalanceOf(&bind.CallOpts{}, opts.From)
	require.NoError(t, err)
	require.Equal(t, balanceAfterFinalization, balanceBeforeFinalization.Add(balanceBeforeFinalization, amount))
}
