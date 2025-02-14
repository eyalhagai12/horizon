package server

import "github.com/labstack/echo/v4"

type registerRoutesFunc func(*echo.Group, Env)

type Server struct {
	Env Env
	App *echo.Echo
}

func NewServer(env Env, app *echo.Echo) *Server {
	return &Server{Env: env, App: app}
}

func (s *Server) Start(port string) error {
	return s.App.Start(port)
}

func (s *Server) Close() error {
	return s.App.Close()
}

func (s *Server) RegisterRoutes(registerRoutes ...registerRoutesFunc) {
	api := s.App.Group("/api")
	for _, rr := range registerRoutes {
		rr(api, s.Env)
	}
}
