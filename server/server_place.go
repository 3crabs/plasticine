package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"plasticine/models"
	"strconv"
)

func (s *server) routesPlace() {
	s.routerAuthGroup.POST("/places", s.addPlace)
	s.routerOpenGroup.GET("/places", s.getPlaces)
	s.routerAuthGroup.PUT("/places/:placeId", s.updatePlace)
	s.routerAuthGroup.DELETE("/places/:placeId", s.deletePlace)
}

func (s *server) addPlace(c echo.Context) error {
	var place models.Place
	if err := c.Bind(&place); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err := s.db.AddPlace(place)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "new place added")
}

func (s *server) getPlaces(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.GetPlaces())
}

func (s *server) updatePlace(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("placeId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	var place models.Place
	if err := c.Bind(&place); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	place.Id = id
	err = s.db.UpdatePlace(place)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "place updated")
}

func (s *server) deletePlace(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("placeId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = s.db.DeletePlace(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "place deleted")
}
