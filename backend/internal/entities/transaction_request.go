package entities

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type TransactionRequest struct {
	Address string
	Level   int
}

func NewTransactionRequest(c echo.Context) (*TransactionRequest, error) {
	address := c.Param("address")
	level := c.QueryParam("level")

	levelInt, err := strconv.ParseInt(level, 10, 32)
	if err != nil {
		return nil, err
	}

	return &TransactionRequest{
		Address: address,
		Level:   int(levelInt),
	}, nil
}
