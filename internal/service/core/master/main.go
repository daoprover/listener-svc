package master

import (
	"context"
	"encoding/hex"
	"github.com/daoprover/listener-svc/internal/data"
	"github.com/daoprover/listener-svc/internal/service/core/github"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"golang.org/x/crypto/sha3"
	"time"
)

type Listener interface {
	Run() error
	AddToQuery(name string, address string) (*string, error)
}

type ListenerData struct {
	ctx    context.Context
	db     data.MasterQ
	query  chan OrderData
	orders map[string]Order
	uuid   uuid.UUID
}

func NewListener(ctx context.Context, client *ethclient.Client, db data.MasterQ) Listener {
	return &ListenerData{
		ctx:    ctx,
		db:     db,
		query:  make(chan OrderData),
		orders: make(map[string]Order),
	}
}

func (l *ListenerData) Run() error {
	for {
		select {
		case <-l.ctx.Done():
			return nil
		case order := <-l.query:
			glistener := github.NewGithubListener(order.ctx, order.db, order.id, order.Name, order.Address)
			if err := glistener.Run(); err != nil {
				return errors.Wrap(err, "Failed to  get info  from github")
			}

		}
	}
}

func (l *ListenerData) AddToQuery(name string, address string) (*string, error) {
	sha := sha3.New256()
	_, err := sha.Write([]byte(time.Now().String()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to write to sha instance")
	}

	order := OrderData{
		Name:    name,
		Address: address,
		ctx:     l.ctx,
		db:      l.db,
		id:      hex.EncodeToString(sha.Sum(nil)),
	}

	l.orders[order.id] = order
	l.query <- order
	return &order.id, nil
}
