package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"skeleton/conf"
	"skeleton/utils"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *Api) initConfig(confpath string) {
	a.Cfg, a.err = conf.LoadConfig(confpath)
	if a.err != nil {
		log.Fatalf("cannot load conf file; err: %s", a.err)
		return
	}
}

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

func (a *Api) initResponder() {
	if a.err != nil {
		log.Fatalf("cannot create Responder instance; err: %s", a.err)
		return
	}
	a.Responder = utils.NewClient(true, true, log.Printf, utils.AlertSimulate, "json", utils.NewGinResponder())
}

// Start beggins Listen and Serve parting it into a new goroutine.
// It is also waiting for an interruption so it performs shutdown.
// Shuts Down Gracefully; waits for the API to finish all requests that are being handled at the moment of shutdown.
func (a *Api) Start(port string) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// ListenAndServe is parted in a new goroutine so it does not block the handling of the shutdown below
	go func() {
		log.Println("Starting skeleton service...")
		if err := a.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error starting Listen and Serve: ", err)
			return
		}
	}()

	// Listen for the interrupt signal
	<-ctx.Done()

	stop()
	log.Println("Shutting down skeleton service...")

	// This context is used to notify service that it has 10 seconds to finish
	// the requests that it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.HttpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Shutdown error: %s\n", err)
	}

	fmt.Println("Skeleton service offline.")
}
