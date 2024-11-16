package entities

import (
	"github.com/labstack/echo/v4"
)

type ReportRequest struct {
	Address string
}

func NewReportRequest(c echo.Context) (*ReportRequest, error) {
	address := c.Param("address")
	return &ReportRequest{Address: address}, nil
}
