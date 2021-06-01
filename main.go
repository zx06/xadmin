package main

import (
	"time"

	"github.com/gin-contrib/expvar"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/requestid"
	"go.uber.org/zap"
)

func main() {
	r := gin.New()
	logger, _ := zap.NewProduction()

	r.Use(requestid.New())

	r.Use(ginzap.Ginzap(logger, time.RFC3339, false))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	r.GET("/debug/vars", expvar.Handler())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.Run()
}
