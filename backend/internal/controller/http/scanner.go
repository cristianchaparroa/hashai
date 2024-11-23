package http

import (
	"context"
	"hashtracker/internal/entities/blockscout"
	"hashtracker/internal/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ScannerController struct {
	uc usecases.ScannerUseCase
}

func NewScannerController(uc usecases.ScannerUseCase) *ScannerController {
	return &ScannerController{
		uc: uc,
	}
}

func (s *ScannerController) GetTransactions(c echo.Context) error {
	r, err := blockscout.NewTransactionRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	root, err := s.uc.Scan(context.Background(), r.Address, r.Level)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, root)
}
