package app

import (
	"hashtracker/config"
	"hashtracker/internal/controller/http"
	"hashtracker/internal/usecases"
	"hashtracker/internal/usecases/repo/blockscout"
	"hashtracker/internal/usecases/repo/polygon"
	"hashtracker/internal/usecases/repo/thegraph"
	validators "hashtracker/internal/usecases/validators"

	server "hashtracker/pkg/httpserver"

	"github.com/labstack/echo/v4"
)

func Run() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	healthController := http.NewHealthController()
	txRepo := blockscout.NewTransactionRepository()
	ensRepo := thegraph.NewENSRepository(cfg.TheGraph.ApiKey)
	ensValidator := validators.NewENSValidator()

	scannerUseCase := usecases.NewScanner(
		txRepo,
		ensRepo,
		ensValidator,
	)
	scannerController := http.NewScannerController(scannerUseCase)

	polygonRepo := polygon.NewHashReportRepository(cfg)
	reporterController := http.NewReporterController(polygonRepo)

	s := echo.New()
	router := http.NewRouter(s)
	router.AddGet("/health", healthController.Health)
	router.AddGet("/transactions/:address", scannerController.GetTransactions)
	router.AddPost("/reports/:address", reporterController.ReportAddress)
	router.List()

	httpServer := server.New(s, server.Port("8080"))
	httpServer.Run()
}
