package api

import (
	"go-zookeeper-client-app/util"
	"log"
)

type API struct {
	cfg *util.Config
}

func NewHandler() (*API, error) {
	a := new(API)

	var err error

	a.cfg, err = util.ConfInitialize()
	if err != nil {
		log.Println("[NewHandler] failed config initialize :", err)
		return nil, err
	}

	return a, nil
}

func (a *API) Close() {
}

func (a *API) GetApiPort() string {
	return a.cfg.ApiPort
}
