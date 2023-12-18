package network

import (
	"context"
	"github.com/daoprover/listener-svc/internal/data"
)

type NetworkListeer interface {
}

type NetworkListeerData struct {
	ID            string
	Ctx           context.Context
	DB            data.MasterQ
	Name          string
	Address       string
	timestampFrom string
	timestampTo   string
}
