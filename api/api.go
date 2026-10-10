package api

import "1-converter/config"

type Api struct {
	config *config.Encrypter
}

func NewApi(cfg *config.Encrypter) *Api {
	return &Api{
		config: cfg,
	}
}

func (a *Api) GetKey() string {
	return a.config.Key
}
