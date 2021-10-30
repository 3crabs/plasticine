package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"plasticine/models"
	"strconv"
)

func (s *server) routesGroup() {
	s.router.POST("/groups", s.addGroup)
	s.router.GET("/groups", s.getGroups)
	s.router.PUT("/groups/:groupId", s.updateGroup)
}

func (s *server) addGroup(c echo.Context) error {
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
