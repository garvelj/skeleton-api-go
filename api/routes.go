package api

import (
	"github.com/gin-gonic/gin"
)

func (a *Api) RoutesRegister() {
	v1 := a.Router.Group("/v1")

	v1.Use(gin.Recovery()) //  gin.Logger()

	v1.GET("/skeleton", a.Ping)
}
