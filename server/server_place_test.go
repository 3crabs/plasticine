package server

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"plasticine/db"
	"plasticine/models"
	"strconv"
	"strings"
	"testing"
)

func (s *server) addPlaceReq(place models.Place) {
	bytes, _ := json.Marshal(place)
	_, c := s.post(strings.NewReader(string(bytes)))
	_ = s.addPlace(c)
}

func (s *server) getPlacesReq() (*httptest.ResponseRecorder, []models.Place) {
	rec, c := s.get()
	_ = s.getPlaces(c)
	var places []models.Place
	_ = json.Unmarshal([]byte(rec.Body.String()), &places)
	return rec, places
}

func (s *server) updatePlaceReq(placeId int, place models.Place) {
	bytes, _ := json.Marshal(place)
	_, c := s.put(strings.NewReader(string(bytes)))
	c.SetParamNames("placeId")
	c.SetParamValues(strconv.Itoa(placeId))
	_ = s.updatePlace(c)
}

func (s *server) deletePlaceReq(placeId int) {
	_, c := s.delete()
	c.SetParamNames("placeId")
	c.SetParamValues(strconv.Itoa(placeId))
	_ = s.deletePlace(c)
}

func TestGetPlaces(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec, places := s.getPlacesReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 0, len(places))
}

func TestAddPlace(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	place := models.Place{Name: "name"}
	s.addPlaceReq(place)

	rec, places := s.getPlacesReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 1, len(places))
	assert.Equal(t, place.Name, places[0].Name)
}

func TestUpdatePlace(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	place := models.Place{Name: "name"}
	s.addPlaceReq(place)

	rec, places := s.getPlacesReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 1, len(places))
	assert.Equal(t, place.Name, places[0].Name)

	place.Id = places[0].Id
	place.Name = "new name"
	s.updatePlaceReq(1, place)

	rec, places = s.getPlacesReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 1, len(places))
	assert.Equal(t, place.Name, places[0].Name)
}

func TestDeletePlace(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	place := models.Place{Name: "name"}
	s.addPlaceReq(place)

	_, places := s.getPlacesReq()
	assert.Equal(t, 1, len(places))

	s.deletePlaceReq(places[0].Id)

	_, places = s.getPlacesReq()
	assert.Equal(t, 0, len(places))
}
