package contracts

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tokamak-network/tokamak-thanos/op-bindings/bindings"
	"github.com/tokamak-network/tokamak-thanos/op-challenger/game/types"
	"github.com/tokamak-network/tokamak-thanos/op-service/sources/batching"
	"github.com/tokamak-network/tokamak-thanos/op-service/txmgr"
)

const (
	methodGameCount   = "gameCount"
	methodGameAtIndex = "gameAtIndex"
	methodGameImpls   = "gameImpls"
	methodCreateGame  = "create"
	methodGames       = "games"
)

type DisputeGameFactoryContract struct {
	multiCaller *batching.MultiCaller
	contract    *batching.BoundContract
}

func NewDisputeGameFactoryContract(addr common.Address, caller *batching.MultiCaller) (*DisputeGameFactoryContract, error) {
	factoryAbi, err := bindings.DisputeGameFactoryMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to load dispute game factory ABI: %w", err)
	}
	return &DisputeGameFactoryContract{
		multiCaller: caller,
		contract:    batching.NewBoundContract(factoryAbi, addr),
	}, nil
}

func (f *DisputeGameFactoryContract) GetGameFromParameters(ctx context.Context, traceType uint32, outputRoot common.Hash, l2BlockNum uint64) (common.Address, error) {
	result, err := f.multiCaller.SingleCall(ctx, batching.BlockLatest, f.contract.Call(methodGames, traceType, outputRoot, common.BigToHash(big.NewInt(int64(l2BlockNum))).Bytes()))
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to fetch game from parameters: %w", err)
	}
	return result.GetAddress(0), nil
}

func (f *DisputeGameFactoryContract) GetGameCount(ctx context.Context, blockHash common.Hash) (uint64, error) {
	result, err := f.multiCaller.SingleCall(ctx, batching.BlockByHash(blockHash), f.contract.Call(methodGameCount))
	if err != nil {
		return 0, fmt.Errorf("failed to load game count: %w", err)
	}
	return result.GetBigInt(0).Uint64(), nil
}

func (f *DisputeGameFactoryContract) GetGame(ctx context.Context, idx uint64, blockHash common.Hash) (types.GameMetadata, error) {
	result, err := f.multiCaller.SingleCall(ctx, batching.BlockByHash(blockHash), f.contract.Call(methodGameAtIndex, new(big.Int).SetUint64(idx)))
	if err != nil {
		return types.GameMetadata{}, fmt.Errorf("failed to load game %v: %w", idx, err)
	}
	return f.decodeGame(result), nil
}

func (f *DisputeGameFactoryContract) GetGameImpl(ctx context.Context, gameType uint32) (common.Address, error) {
	result, err := f.multiCaller.SingleCall(ctx, batching.BlockLatest, f.contract.Call(methodGameImpls, gameType))
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to load game impl for type %v: %w", gameType, err)
	}
	return result.GetAddress(0), nil
}

func (f *DisputeGameFactoryContract) GetGamesAtOrAfter(ctx context.Context, blockHash common.Hash, earliestTimestamp uint64) ([]types.GameMetadata, error) {
	count, err := f.GetGameCount(ctx, blockHash)
	if err != nil {
		return nil, err
	}
	batchSize := uint64(f.multiCaller.BatchSize())
	rangeEnd := count

	var games []types.GameMetadata
	for {
		if rangeEnd == uint64(0) {
			// rangeEnd is exclusive so if its 0 we've reached the end.
			return games, nil
		}
		rangeStart := uint64(0)
		if rangeEnd > batchSize {
			rangeStart = rangeEnd - batchSize
		}
		calls := make([]*batching.ContractCall, 0, rangeEnd-rangeStart)
		for i := rangeEnd - 1; ; i-- {
			calls = append(calls, f.contract.Call(methodGameAtIndex, new(big.Int).SetUint64(i)))
			// Break once we've added the last call to avoid underflow when rangeStart == 0
			if i == rangeStart {
				break
			}
		}

		results, err := f.multiCaller.Call(ctx, batching.BlockByHash(blockHash), calls...)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch games: %w", err)
		}

		for _, result := range results {
			game := f.decodeGame(result)
			if game.Timestamp < earliestTimestamp {
				return games, nil
			}
			games = append(games, game)
		}
		rangeEnd = rangeStart
	}
}

func (f *DisputeGameFactoryContract) GetAllGames(ctx context.Context, blockHash common.Hash) ([]types.GameMetadata, error) {
	count, err := f.GetGameCount(ctx, blockHash)
	if err != nil {
		return nil, err
	}

	calls := make([]*batching.ContractCall, count)
	for i := uint64(0); i < count; i++ {
		calls[i] = f.contract.Call(methodGameAtIndex, new(big.Int).SetUint64(i))
	}

	results, err := f.multiCaller.Call(ctx, batching.BlockByHash(blockHash), calls...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch games: %w", err)
	}

	var games []types.GameMetadata
	for _, result := range results {
		games = append(games, f.decodeGame(result))
	}
	return games, nil
}

func (f *DisputeGameFactoryContract) CreateTx(traceType uint32, outputRoot common.Hash, l2BlockNum uint64) (txmgr.TxCandidate, error) {
	call := f.contract.Call(methodCreateGame, traceType, outputRoot, common.BigToHash(big.NewInt(int64(l2BlockNum))).Bytes())
	return call.ToTxCandidate()
}

func (f *DisputeGameFactoryContract) decodeGame(result *batching.CallResult) types.GameMetadata {
	gameType := result.GetUint32(0)
	timestamp := result.GetUint64(1)
	proxy := result.GetAddress(2)
	return types.GameMetadata{
		GameType:  gameType,
		Timestamp: timestamp,
		Proxy:     proxy,
	}
}