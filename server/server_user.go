package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"plasticine/models"
	"strconv"
)

func (s *server) routesUser() {
	s.router.GET("/students", s.getStudents)
	s.router.GET("/teachers", s.getStudents)
	s.router.POST("/users", s.addUser)
	s.router.PUT("/users/:userId", s.updateUser)
	s.router.GET("/users/:userId", s.getUserInfo)
	s.router.DELETE("/users/:userId", s.deleteUser)
}

func (s *server) getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.GetUsersByRole(models.Student))
}

func (s *server) getTeachers(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.GetUsersByRole(models.Teacher))
}

func (s *server) addUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err := s.db.AddUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "new users added")
}

func (s *server) updateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	user.Id = id
	err = s.db.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "user updated")
}

func (s *server) getUserInfo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	userInfo, err := s.db.GetUserInfo(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, userInfo)
}

func (s *server) deleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = s.db.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "user deleted")
}
