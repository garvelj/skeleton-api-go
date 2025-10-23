package utils

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Client struct {
	shouldLog   bool
	shouldAlert bool
	logf        func(string, ...any)
	alertf      func(string, ...any)
	out         string
	Responder
}

func NewClient(ShouldLog, ShouldAlert bool, Logf, Alertf func(string, ...any), Out string, Responder Responder) *Client {
	return &Client{
		shouldLog:   ShouldLog,
		shouldAlert: ShouldAlert,
		logf:        Logf,
		alertf:      Alertf,
		out:         Out,
		Responder:   Responder,
	}
}

func (c *Client) Log(message string, args ...any) {
	if c.shouldLog {
		c.logf(message, args...)
	}
}

func (c *Client) Alert(message string, args ...any) {
	if c.shouldAlert {
		c.alertf(message, args)
	}
}

func (c *Client) Internal(message string, args ...any) {
	c.Log(message, args...)
	c.Alert(message, args)
}

func (c *Client) RespondWithContext(ctx context.Context, statusCode int, body any) *Client {
	c.Responder.RespondWithContext(ctx, statusCode, body)
	return c
}

type Responder interface {
	RespondWithContext(context.Context, int, any)
}

type GinResponder struct{}

func NewGinResponder() GinResponder {
	return GinResponder{}
}

func (g GinResponder) RespondWithContext(ctx context.Context, statusCode int, body any) {
	c, ok := ctx.(*gin.Context)
	if !ok {
		panic("context %v not of type *gin.Context")
	}
	c.JSON(statusCode, body)
}

func AlertSimulate(message string, args ...any) {
	message = fmt.Sprintf(message, args...)
	fmt.Println("\n!ALERT! ", message, " ⚠️")
}
