package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"xadmin/app/config"
	"xadmin/app/service"
)

func RunServer() {
	// http server
	e := echo.New()

	// middleware
	// e.Use(middleware.AddTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	// jwt
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &service.JwtClaims{},
		SigningKey: config.Config().SigningKey,
	}))
	p := prometheus.NewPrometheus("xadmin", nil)
	p.Use(e)

	// router
	UseRouter(e)

	go func() {
		var addr = fmt.Sprintf(":%d", config.Config().Port)
		if err := e.Start(addr); err != nil && err != http.ErrServerClosed {

			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
