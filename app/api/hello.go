package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TaskListHandler 任务列表
// @Description 任务列表
// @Tags 筛分聚类
// @Accept  json
// @Produce  json
// @Param X-Request-Id header string false "可选,追踪请求,不传该值则由服务端自行生成"
// @Router /api/tasks [get]
func Hello(c echo.Context) error {
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