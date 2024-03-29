package app

import (
	"xadmin/app/api"

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

	// swagger
	g.GET("/docs/*", echoSwagger.WrapHandler)
	g.GET("/hello", api.HelloHandler)
}
