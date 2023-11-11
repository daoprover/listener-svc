package master

import (
	"context"
	"github.com/daoprover/listener-svc/internal/data"
)

type Order interface {
	GetStatus() string
}

type OrderData struct {
	id      string
	ctx     context.Context
	db      data.MasterQ
	Name    string
	Address string
}

func (o OrderData) GetStatus() string {
	//TODO implement me
	panic("implement me")
}
