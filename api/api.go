package api

import (
	"fmt"
	"net/http"
	"net/url"
	"skeleton/conf"
	"skeleton/storage"

	"github.com/gin-gonic/gin"
)

func New(confpath string) *Api {
	a := &Api{}

	a.initConfig(confpath)
	a.initRouter()
	a.initHttpServer()

	return a
}

type Api struct {
	Cfg        *conf.Cfg
	Router     *gin.Engine
	HttpServer *http.Server

	err error
}

func (a *Api) Ping(c *gin.Context) {

	db, err := storage.NewStorage("postgres", a.Cfg.Postgres)
	if err != nil {
		fmt.Println("err new storage: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("DB: ", db)

	parmas := url.Values{}
	total, err := db.UserCount(parmas)
	if err != nil {
		fmt.Println("err new storage: ", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Total: ", total)

	c.JSON(http.StatusOK, gin.H{
		"ratle": "ratle",
	})
}
