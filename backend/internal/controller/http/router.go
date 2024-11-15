package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server *echo.Echo
}

func NewRouter(s *echo.Echo) *Router {
	s.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	return &Router{s}
}

func (r *Router) AddPost(path string, fn func(c echo.Context) error) *Router {
	r.server.POST(path, fn)
	return r
}

func (r *Router) AddGet(path string, fn func(c echo.Context) error) *Router {
	r.server.GET(path, fn)
	return r
}

func (r *Router) List() {
	data, err := json.MarshalIndent(r.server.Routes(), "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(data))
}
