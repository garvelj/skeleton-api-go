package api

import (
	"net/http"
	"skeleton/conf"
	"skeleton/storage"
	"skeleton/utils"

	"github.com/gin-gonic/gin"
)

func New(confpath string) *Api {
	a := &Api{}

	a.initConfig(confpath)
	a.initRouter()
	a.initHttpServer()
	a.initResponder()

	return a
}

type Api struct {
	Cfg        *conf.Cfg
	Router     *gin.Engine
	HttpServer *http.Server
	Responder  *utils.Client

	err error
}

func (a *Api) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func PostgressConn(cfg conf.DbCfg) storage.Storage {
	storage, _ := storage.NewStorage("postgres", cfg)
	return storage
}
