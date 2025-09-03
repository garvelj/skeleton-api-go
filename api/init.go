package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *Api) initRouter() {
	a.Router = gin.Default()
	a.Router.SetTrustedProxies(nil)
	a.RoutesRegister()
}

func (a *Api) initHttpServer() {
	a.HttpServer = &http.Server{
		Addr:         ":13131", // TODO: Load from config
		Handler:      a.Router,
		ReadTimeout:  5 * time.Second, // TODO: Load from config
		WriteTimeout: 5 * time.Second, // TODO: Load from config
	}
}

func (a *Api) Start(port string) {
	go func() {
		log.Println("Starting skeleton service...")
		if err := a.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error starting Listen and Serve: ", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// waits for shutdown signal
	<-quit
	log.Println("Shutting down skeleton service...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.HttpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Shutdown error: %s\n", err)
	}

	fmt.Println("Skeleton service offline.")
}
