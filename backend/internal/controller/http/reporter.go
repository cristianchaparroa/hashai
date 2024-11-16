package http

import (
	"context"
	"hashtracker/internal/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReporterController struct {
	rp usecases.PolygonRepository
}

func NewReporterController(rp usecases.PolygonRepository) *ReporterController {
	return &ReporterController{
		rp: rp,
	}
}

func (r *ReporterController) ReportAddress(c echo.Context) error {
	response, err := r.rp.Resolve(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response.Tsx)
}
