package http

import (
	"context"
	"hashtracker/internal/entities"
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
	r, err := entities.NewTransactionRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	root, err := s.uc.GetTransactions(context.Background(), r.Address, r.Level)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, root)
}
