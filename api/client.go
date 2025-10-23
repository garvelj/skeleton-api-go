package api

import (
	"fmt"
	"net/http"
	"skeleton/model"

	"github.com/gin-gonic/gin"
)

func (a *Api) ClientRead(c *gin.Context) {
	var (
		params  = c.Request.URL.Query()
		db      = PostgressConn(a.Cfg.Postgres)
		clients []model.Client
		total   int
		err     error
	)

	clients, total, err = db.ClientRead(params)
	if err != nil {
		// TODO error logging and repsonding
		// log.Fatalf("Could not read clients from database with params [%+v]; err: %s", params, err)
		fmt.Printf("Could not read clients from database with params [%+v]; err: %s", params, err)
		c.JSON(http.StatusOK, "u Svedskoj problemi")
		return
	}

	a.Responder.RespondWithContext(c, http.StatusOK, map[string]any{
		"total": total,
		"items": clients,
	}).Internal("jaga logovo i alertovo ")

	// c.JSON(http.StatusOK, gin.H{
	// 	"total": total,
	// 	"items": clients,
	// })
}
