package network

import (
	"context"
	"github.com/daoprover/listener-svc/internal/data"
)

func NewNetworkListener(ctx context.Context, db data.MasterQ, id string, name string, address string, timestampFrom, timestampTo string) NetworkListeer {
	return &NetworkListeerData{
		Ctx:           ctx,
		DB:            db,
		ID:            id,
		Name:          name,
		Address:       address,
		timestampFrom: timestampFrom,
		timestampTo:   timestampTo,
	}
}

func (n *NetworkListeerData) Run() error {

	return nil
}

func (n *NetworkListeerData) GetContractMeta() error {

	return nil
}

func (n *NetworkListeerData) GetConfirmedTransactions() error {
	return nil
}
