package api

import (
	"net/http"
	"skeleton/conf"

	"github.com/gin-gonic/gin"
)

func New(confpath string) *Api {
	a := &Api{}

	a.initRouter()
	a.initHttpServer()

	return a
}

type Api struct {
	Cfg        conf.Cfg
	Router     *gin.Engine
	HttpServer *http.Server
}

func (a *Api) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ratle": "ratle",
	})
}
