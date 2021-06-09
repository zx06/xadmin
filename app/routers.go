package app

import (
	"xadmin/app/api"
	"xadmin/app/config"
	"xadmin/app/service"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

// UseRouter 注册路由
func UseRouter(e *echo.Echo) {
	// 前端静态文件
	e.File("/", "public/index.html")
	e.Static("/", "public/")
	// api
	apiGroup := e.Group("/api")
	useApiRouter(apiGroup)
}

func useApiRouter(g *echo.Group) {
	// jwt
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &service.JwtClaims{},
		SigningKey: config.Config().SigningKey,
	}))
	// swagger
	g.GET("/docs/*", echoSwagger.WrapHandler)
	g.GET("/hello", api.HelloHandler)
}
