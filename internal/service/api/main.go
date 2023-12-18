package api

import (
	"context"
	"github.com/daoprover/listener-svc/internal/data/pg"
	"github.com/daoprover/listener-svc/internal/service/core/master"
	"net"
	"net/http"

	"github.com/daoprover/listener-svc/internal/config"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	config   config.Config
}

func (s *service) run() error {
	s.log.Info("Service started")
	listener := master.NewListener(context.Background(), pg.NewMasterQ(s.config.DB()), s.config)
	go listener.Run()
	r := s.router(listener)

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		config:   cfg,
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
