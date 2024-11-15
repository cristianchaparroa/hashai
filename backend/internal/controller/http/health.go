package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (s *HealthController) Health(c echo.Context) error {
	return c.JSON(http.StatusOK, "it's working...")
}
