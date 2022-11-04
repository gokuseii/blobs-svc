package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/url"
)

type HorizonConfig struct {
	once   comfig.Once
	getter kv.Getter
}

func NewHorizon(getter kv.Getter) *HorizonConfig {
	return &HorizonConfig{
		getter: getter,
	}
}

func (k *HorizonConfig) HorizonURL() *url.URL {
	return k.once.Do(func() interface{} {
		var config struct {
			URL *url.URL `fig:"url"`
		}
		err := figure.
			Out(&config).
			With(figure.BaseHooks).
			From(kv.MustGetStringMap(k.getter, "horizon")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out api"))
		}

		return config.URL
	}).(*url.URL)
}
