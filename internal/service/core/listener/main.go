package listener

import (
	"context"
	"github.com/daoprover/listener-svc/internal/data"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"math/big"
)

func NewListener(ctx context.Context, client *ethclient.Client, startBlock *big.Int, db data.MasterQ) *ListenerData {
	return &ListenerData{
		ctx:        ctx,
		client:     client,
		startBlock: startBlock,
		db:         db,
	}
}

func (l *ListenerData) Run() error {
	currentBlock := l.startBlock
	for {
		block, err := l.client.BlockByNumber(l.ctx, currentBlock)
		if err != nil {
			return errors.Wrap(err, "failed to get block  by  number")
		}

		for _, tx := range block.Transactions() {
		   err = l.db.New().TransactionsQ().Insert(l.prepareTx(tx))
		   if  err != nil {
			    return errors.Wrap(err,  "failed to insert tx")
		   }
		}
		currentBlock = currentBlock.Add(currentBlock, big.NewInt(1))
	}
}

func (l *ListenerData) prepareTx(tx *types.Transaction) *data.Transaction {
	return &data.Transaction{
		Hash:      tx.Hash().Hex(),
		Recipient: tx.To().Hex(),
		Sender: tx.
	}
}
