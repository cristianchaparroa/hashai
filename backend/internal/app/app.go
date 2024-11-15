package app

import (
	"github.com/labstack/echo/v4"
	"hashtracker/internal/controller/http"
	server "hashtracker/pkg/httpserver"
)

func Run() {
	controller := http.NewHealthController()
	s := echo.New()
	router := http.NewRouter(s)
	router.AddGet("/health", controller.Health)
	router.List()

	httpServer := server.New(s, server.Port("8080"))
	httpServer.Run()
}
