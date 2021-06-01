package router

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
	useApiRouter(e, "/api")
}

func useApiRouter(e *echo.Echo, prefix string) {
	r := e.Group(prefix)
	// swagger
	r.GET("/docs/*", echoSwagger.WrapHandler)
	r.GET("/hello", api.HelloHandler)
}
