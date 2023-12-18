package config

import (
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type ThirdPartyConfiger interface {
	ThirdPartyConfig() *ThirdPartyConfig
}

type ThirdPartyConfig struct {
	ApiKey  string `fig:"api_key"`
	ApiPath string `fig:"api_path"`
}

func NewThirdPartyConfiger(getter kv.Getter) ThirdPartyConfiger {
	return &thirdPartyConfig{
		getter: getter,
	}
}

type thirdPartyConfig struct {
	getter kv.Getter
	once   comfig.Once
}

func (c *thirdPartyConfig) ThirdPartyConfig() *ThirdPartyConfig {
	return c.once.Do(func() interface{} {
		raw := kv.MustGetStringMap(c.getter, "third_party")
		config := NetworksConfig{}
		err := figure.Out(&config).With(figure.BaseHooks).From(raw).Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out"))
		}
		return &config
	}).(*ThirdPartyConfig)
}
