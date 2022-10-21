package api

import (
	"github.com/gin-gonic/gin"
	"go-zookeeper-client-app/util"
	"log"
	"net/http"
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

func (a *API) HealthCheck(gCtx *gin.Context) {
	gCtx.JSON(http.StatusOK, map[string]string{"result": "success"})
	return
}
