package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"plasticine/db"
)

type server struct {
	router          *echo.Echo
	routerOpenGroup *echo.Group
	routerAuthGroup *echo.Group
	db              db.DB
	port            string
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
	s.routerOpenGroup = s.router.Group("/api")
	s.routerAuthGroup = s.router.Group("/api", s.Auth())
	s.routesGroup()
	s.routesLesson()
	s.routesPlace()
	s.routesSubject()
	s.routesUser()
}

func (s *server) Run() {
	s.router.Logger.Fatal(s.router.Start(s.port))
}

func (s *server) Auth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "admin" {
			return true, nil
		}
		return false, nil
	})
}
