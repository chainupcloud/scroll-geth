package sync_service

import (
	"context"
	"math/big"

	"github.com/chainupcloud/scroll-geth"
	"github.com/chainupcloud/scroll-geth/common"
	"github.com/chainupcloud/scroll-geth/core/types"
)

// We cannot use ethclient.Client directly as that would lead
// to circular dependency between eth, rollup, and ethclient.
type EthClient interface {
	BlockNumber(ctx context.Context) (uint64, error)
	ChainID(ctx context.Context) (*big.Int, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)
	TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error)
	BlockByHash(ctx context.Context, hash common.Hash) (*types.Block, error)
}
