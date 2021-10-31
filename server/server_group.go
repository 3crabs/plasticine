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
	s.router.GET("/groups/:groupId", s.getGroup)
	s.router.PUT("/groups/:groupId", s.updateGroup)
	s.router.DELETE("/groups/:groupId", s.deleteGroup)
	s.router.GET("/groups/:groupId/students", s.getGroupStudents)
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

func (s *server) getGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	group, err := s.db.GetGroup(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, group)
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

func (s *server) deleteGroup(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = s.db.DeleteGroup(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "group deleted")
}

func (s *server) getGroupStudents(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, s.db.GetGroupStudents(id))
}
