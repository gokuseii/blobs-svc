package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/connectors/keyer"
	"gitlab.com/tokend/connectors/submit"
	"net/url"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	keyer.Keyer
	submit.Submission
	HorizonURL() *url.URL
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	types.Copuser
	comfig.Listenerer
	keyer.Keyer
	submit.Submission
	getter kv.Getter
	*HorizonConfig
}

func New(getter kv.Getter) Config {
	return &config{
		getter:        getter,
		Databaser:     pgdb.NewDatabaser(getter),
		Copuser:       copus.NewCopuser(getter),
		Listenerer:    comfig.NewListenerer(getter),
		Logger:        comfig.NewLogger(getter, comfig.LoggerOpts{}),
		Keyer:         keyer.NewKeyer(getter),
		Submission:    submit.NewSubmission(getter),
		HorizonConfig: NewHorizon(getter),
	}
}
