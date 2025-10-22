package api

import (
	"github.com/gin-gonic/gin"
)

func (a *Api) RoutesRegister() {
	v1 := a.Router.Group("/v1/skeleton")

	v1.Use(gin.Recovery()) //  gin.Logger()

	v1.GET("/ping", a.Ping)

	clients := v1.Group("/clients")
	clients.GET("", a.ClientRead)
}
