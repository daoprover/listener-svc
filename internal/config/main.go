package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	NetworksConfiger
	comfig.Listenerer
	ThirdPartyConfiger
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	NetworksConfiger
	comfig.Listenerer
	ThirdPartyConfiger
	getter kv.Getter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:             getter,
		NetworksConfiger:   NewEthRPCConfiger(getter),
		Databaser:          pgdb.NewDatabaser(getter),
		Copuser:            copus.NewCopuser(getter),
		ThirdPartyConfiger: NewThirdPartyConfiger(getter),
		Listenerer:         comfig.NewListenerer(getter),
		Logger:             comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}
