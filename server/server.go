package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"plasticine/db"
	"plasticine/models"
	"strconv"
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
	s.router.PUT("/groups/:groupId", s.updateGroup)
}

func (s *server) addGroups(c echo.Context) error {
	var group models.Group
	if err := c.Bind(&group); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err := s.db.AddGroup(group)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "new groups added")
}

func (s *server) getGroups(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.GetGroups())
}

func (s *server) updateGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var group models.Group
	if err := c.Bind(&group); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	group.Id = id
	err = s.db.UpdateGroup(group)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "group updated")
}

func (s *server) Run() {
	s.router.Logger.Fatal(s.router.Start(s.port))
}
