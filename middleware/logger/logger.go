package logger

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		log.Info(fmt.Sprintf("[GIN] %s | %d | %s | %s | Latency: %v",
			time.Now().Format("2006-01-02 15:04:05"),
			c.Writer.Status(),
			c.Request.Method,
			c.Request.URL.Path,
			time.Since(start),
		))
	}
}
