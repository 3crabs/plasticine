package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"plasticine/models"
	"strconv"
)

func (s *server) routesLesson() {
	s.router.POST("/lessons", s.addLesson)
	s.router.GET("/lessons", s.getLessons)
	s.router.PUT("/lessons/:lessonId", s.updateLesson)
	s.router.DELETE("/lessons/:lessonId", s.deleteLesson)
}

func (s *server) addLesson(c echo.Context) error {
	var lesson models.Lesson
	if err := c.Bind(&lesson); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err := s.db.AddLesson(lesson)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "new lesson added")
}

func (s *server) getLessons(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("groupId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, s.db.GetLessons(id))
}

func (s *server) updateLesson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("lessonId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var lesson models.Lesson
	if err := c.Bind(&lesson); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	lesson.Id = id
	err = s.db.UpdateLesson(lesson)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "lesson updated")
}

func (s *server) deleteLesson(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("lessonId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = s.db.DeleteLesson(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "lesson deleted")
}
