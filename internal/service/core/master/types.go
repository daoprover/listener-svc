package master

import (
	"context"
	"github.com/daoprover/listener-svc/internal/data"
	"github.com/daoprover/listener-svc/internal/service/core/cryptoapi"
)

type Order interface {
	GetStatus() string
	GetData() OrderData
}

type OrderData struct {
	id       string
	ctx      context.Context
	db       data.MasterQ
	Name     string
	Address  string
	TimeFrom uint
	TimeTo   uint
	Holders  map[string]cryptoapi.Holder
}

func (o OrderData) GetStatus() string {
	//TODO implement me
	return ""
}

func (o OrderData) GetData() OrderData {
	return o
}

func (o OrderData) SetHolders(holders map[string]cryptoapi.Holder) {
	o.Holders = holders
}
