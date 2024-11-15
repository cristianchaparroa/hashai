package app

import (
	"hashtracker/internal/controller/http"
	"hashtracker/internal/usecases"
	"hashtracker/internal/usecases/repo/blockscout"
	server "hashtracker/pkg/httpserver"

	"github.com/labstack/echo/v4"
)

func Run() {
	healthController := http.NewHealthController()
	txRepo := blockscout.NewTransactionRepository()
	scannerUseCase := usecases.NewScanner(txRepo)
	scannerController := http.NewScannerController(scannerUseCase)

	s := echo.New()
	router := http.NewRouter(s)
	router.AddGet("/health", healthController.Health)
	router.AddGet("/transactions/:address", scannerController.GetTransactions)
	router.List()

	httpServer := server.New(s, server.Port("8080"))
	httpServer.Run()
}
