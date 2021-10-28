package server

import (
	"github.com/labstack/echo/v4"
	"plasticine/db"
)

type server struct {
	router *echo.Echo
	db     db.DB
	port   string
}

func NewServer(port string, db db.DB) *server {
	s := &server{
		router: echo.New(),
		db:     db,
		port:   port,
	}
	s.routes()
	return s
}

func (s *server) routes() {
	s.routesGroup()
}

func (s *server) Run() {
	s.router.Logger.Fatal(s.router.Start(s.port))
}
