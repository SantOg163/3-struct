package api

import (
	"demo/struct/config"
)

type Api struct {
	conf config.Config
}

func NewApi(conf config.Config) *Api {
	return &Api{conf: conf}
}
