package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("\n%s: %s\t |%d| %s\t %s %s \n",
			params.TimeStamp,
			params.ClientIP,
			params.StatusCode,
			params.Latency,
			params.Method,
			params.Path,
		)
	})
}
