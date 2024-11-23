package http

import (
	"context"
	"fmt"
	"hashtracker/internal/entities/polygon"
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
	report, err := polygon.NewReportRequest(c)
	if err != nil {
		fmt.Println("--> error request: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response, err := r.rp.CreateReport(context.Background(), report)
	if err != nil {
		fmt.Println("--> error create report: ", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
