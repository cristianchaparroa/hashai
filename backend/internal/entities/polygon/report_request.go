package polygon

import (
	"errors"
	"github.com/labstack/echo/v4"
	"math/big"
	"time"
)

type ReportRequest struct {
	Address  string
	Comments string `json:"comments"`
	Source   string `json:"source"`
	Date     *big.Int
}

func NewReportRequest(c echo.Context) (*ReportRequest, error) {
	address := c.Param("address")
	r := &ReportRequest{}
	if err := c.Bind(r); err != nil {
		return nil, errors.New("is not possible to bind request request body")
	}

	return &ReportRequest{
		Address:  address,
		Comments: r.Comments,
		Source:   r.Source,
		Date:     big.NewInt(time.Now().Unix()),
	}, nil
}
