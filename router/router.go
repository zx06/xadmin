package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UseRouter(e *echo.Echo) {
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
