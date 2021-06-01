package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloHandler Hello
// @Description Hello
// @Tags Hello
// @Accept  json
// @Produce  json
// @Param X-Request-Id header string false "可选,追踪请求,不传该值则由服务端自行生成"
// @Router /api/hello [get]
func HelloHandler(c echo.Context) error {
	err := c.JSONPretty(http.StatusOK, echo.Map{
		"real_ip":  c.RealIP(),
		"request":  c.Request().Header,
		"response": c.Response().Header(),
	}, "    ")
	if err != nil {
		return err
	}
	return nil
}