package api

import (
	"github.com/gin-gonic/gin"
	"go-zookeeper-client-app/util"
	"log"
	"net/http"
)

type API struct {
	cfg       *util.Config
	aLog      *util.Logger
	zookeeper *util.ZKCon
}

func NewHandler() (*API, error) {
	a := new(API)

	var err error

	a.cfg, err = util.ConfInitialize()
	if err != nil {
		log.Println("[NewHandler] failed config initialize :", err)
		return nil, err
	}

	a.aLog, err = util.LogInitialize(a.cfg.LogInfo.LogPath, a.cfg.LogInfo.LogLevel)
	if err != nil {
		log.Println("[NewHandler] failed log initialize : ", err)
		return nil, err
	}

	a.zookeeper, err = util.ZookeeperInitialize(a.cfg)
	if err != nil {
		log.Println("[NewHandler] failed zookeeper initialize : ", err)
	}

	return a, nil
}

func (a *API) Close() {
}

func (a *API) GetApiPort() string {
	return a.cfg.ApiInfo.Port
}

func (a *API) HealthCheck(gCtx *gin.Context) {
	gCtx.JSON(http.StatusOK, map[string]string{"result": "success"})
	return
}
