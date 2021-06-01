package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)
// UseRouter 注册路由
func UseRouter(e *echo.Echo) {
	// 前端静态文件
	e.File("/", "public/index.html")
	e.Static("/", "public/")
	api := e.Group("/api")
	api.GET("/", func(c echo.Context) error {
		c.JSONPretty(http.StatusOK, echo.Map{
			"real_ip":  c.RealIP(),
			"request":  c.Request().Header,
			"response": c.Response().Header(),
		}, "    ")
		return nil
	})
}
