package http

import (
	"context"
	"fmt"
	"hashtracker/internal/entities"
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
	report, err := entities.NewReportRequest(c)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := r.rp.Resolve(context.Background(), report.Address)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response.Tsx)
}
