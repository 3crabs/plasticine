package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"plasticine/models"
	"strconv"
)

func (s *server) routesSubject() {
	s.router.POST("/subjects", s.addSubject)
	s.router.GET("/subjects", s.getGroups)
	s.router.PUT("/subjects/:subjectId", s.updateGroup)
}

func (s *server) addSubject(c echo.Context) error {
	var subject models.Subject
	if err := c.Bind(&subject); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err := s.db.AddSubject(subject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "new subjects added")
}

func (s *server) getSubjects(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.GetSubjects())
}

func (s *server) updateSubject(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("subjectId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var subject models.Subject
	if err := c.Bind(&subject); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	subject.Id = id
	err = s.db.UpdateSubject(subject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "subject updated")
}
