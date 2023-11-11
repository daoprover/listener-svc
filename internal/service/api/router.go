package api

import (
	"github.com/daoprover/listener-svc/internal/data/pg"
	"github.com/daoprover/listener-svc/internal/service/api/handlers"
	"github.com/daoprover/listener-svc/internal/service/core/master"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(listener master.Listener) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxMasterQ(pg.NewMasterQ(s.config.DB())),
			handlers.CtxMasterRunner(listener),
		),
	)
	r.Route("/integrations/listener-svc", func(r chi.Router) {
		r.Route("/order", func(r chi.Router) {
			r.Post("/", handlers.CreateInfoOrder)
		})
	})

	return r
}
