package handlers

import (
	"context"
	"github.com/daoprover/listener-svc/internal/config"
	"github.com/daoprover/listener-svc/internal/data"
	"github.com/daoprover/listener-svc/internal/service/core/master"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	masterRunnerCtxKey
	configCtxKey
	masterqCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxMasterRunner(entry master.Listener) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, masterRunnerCtxKey, entry)
	}
}

func MasterRunner(r *http.Request) master.Listener {
	return r.Context().Value(masterRunnerCtxKey).(master.Listener)
}

func CtxConfig(entry config.Config) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, configCtxKey, entry)
	}
}

func Config(r *http.Request) config.Config {
	return r.Context().Value(configCtxKey).(config.Config)
}

func CtxMasterQ(entry data.MasterQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, masterqCtxKey, entry)
	}
}

func MasterQ(r *http.Request) data.MasterQ {
	return r.Context().Value(masterqCtxKey).(data.MasterQ).New()
}
