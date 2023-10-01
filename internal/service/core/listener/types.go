package listener

import (
	"context"
	"github.com/daoprover/listener-svc/internal/data"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Listener interface {
	Run()
}

type ListenerData struct {
	ctx        context.Context
	client     *ethclient.Client
	startBlock *big.Int
	db         data.MasterQ
}
