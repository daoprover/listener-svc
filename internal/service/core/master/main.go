package master

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/daoprover/listener-svc/internal/config"
	"github.com/daoprover/listener-svc/internal/data"
	"github.com/daoprover/listener-svc/internal/service/core/cryptoapi"
	"github.com/daoprover/listener-svc/internal/service/core/github"
	"github.com/google/uuid"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"golang.org/x/crypto/sha3"
	"time"
)

type Listener interface {
	Run() error
	AddToQuery(name string, address string, timeTo uint, timeForm uint) (*string, error)
	ValidateOrder(id string) (OrderData, error)
}

type ListenerData struct {
	ctx    context.Context
	db     data.MasterQ
	query  chan *OrderData
	orders map[string]Order
	uuid   uuid.UUID
	config config.Config
}

func NewListener(ctx context.Context, db data.MasterQ, config config.Config) Listener {
	return &ListenerData{
		ctx:    ctx,
		db:     db,
		config: config,
		query:  make(chan *OrderData),
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

			ethListener := cryptoapi.NewCryptoAPI(l.config.ThirdPartyConfig().ApiPath, l.config.ThirdPartyConfig().ApiKey)
			responceTx, err := ethListener.GetInternalTransactionByAddress(order.Address, order.TimeFrom, order.TimeTo)
			if err != nil {
				continue
			}

			fmt.Println("responce tx: ", responceTx)

			holders := ethListener.GetTokensHoldersByTime(*responceTx, order.TimeFrom, order.TimeTo)
			order.Holders = holders
			fmt.Printf("set: ", order.Holders)
			fmt.Printf("data: ", holders)

		}
	}
}

func (l *ListenerData) AddToQuery(name string, address string, timeTo, timeFrom uint) (*string, error) {
	sha := sha3.New256()
	_, err := sha.Write([]byte(time.Now().String()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to write to sha instance")
	}

	order := OrderData{
		Name:     name,
		Address:  address,
		TimeFrom: timeFrom,
		TimeTo:   timeTo,
		ctx:      l.ctx,
		db:       l.db,
		id:       hex.EncodeToString(sha.Sum(nil)),
	}

	l.orders[order.id] = order
	l.query <- &order
	return &order.id, nil
}

func (l *ListenerData) ValidateOrder(id string) (OrderData, error) {
	order := l.orders[id]
	if order == nil {
		return OrderData{}, errors.New("Failed to get order")
	}

	return order.GetData(), nil
}
