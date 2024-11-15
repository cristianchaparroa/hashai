package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (s *HealthController) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, "it's working...")
}
