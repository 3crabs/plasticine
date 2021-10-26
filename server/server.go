package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"plasticine/db"
	"plasticine/models"
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
	s.router.POST("/groups", s.addGroups)
	s.router.GET("/groups", s.getGroups)
}

func (s *server) addGroups(c echo.Context) error {
	var groups []models.Group
	if err := c.Bind(&groups); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err := s.db.AddGroups(groups)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "new groups added")
}

func (s *server) getGroups(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.GetGroups())
}

func (s *server) Run() {
	s.router.Logger.Fatal(s.router.Start(s.port))
}
