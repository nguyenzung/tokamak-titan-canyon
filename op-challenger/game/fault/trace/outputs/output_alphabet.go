package outputs

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/tokamak-network/tokamak-thanos/op-challenger/game/fault/contracts"
	"github.com/tokamak-network/tokamak-thanos/op-challenger/game/fault/trace"
	"github.com/tokamak-network/tokamak-thanos/op-challenger/game/fault/trace/alphabet"
	"github.com/tokamak-network/tokamak-thanos/op-challenger/game/fault/trace/split"
	"github.com/tokamak-network/tokamak-thanos/op-challenger/game/fault/types"
	"github.com/tokamak-network/tokamak-thanos/op-challenger/metrics"
)

func NewOutputAlphabetTraceAccessor(
	logger log.Logger,
	m metrics.Metricer,
	prestateProvider types.PrestateProvider,
	rollupClient OutputRootProvider,
	splitDepth types.Depth,
	prestateBlock uint64,
	poststateBlock uint64,
) (*trace.Accessor, error) {
	outputProvider := NewTraceProviderFromInputs(logger, prestateProvider, rollupClient, splitDepth, prestateBlock, poststateBlock)
	alphabetCreator := func(ctx context.Context, localContext common.Hash, depth types.Depth, agreed contracts.Proposal, claimed contracts.Proposal) (types.TraceProvider, error) {
		provider := alphabet.NewTraceProvider(agreed.L2BlockNumber, depth)
		return provider, nil
	}
	cache := NewProviderCache(m, "output_alphabet_provider", alphabetCreator)
	selector := split.NewSplitProviderSelector(outputProvider, splitDepth, OutputRootSplitAdapter(outputProvider, cache.GetOrCreate))
	return trace.NewAccessor(selector), nil
}
