package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"plasticine/db"
	"strings"
	"testing"
)

func TestGetGroups(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec, c := s.Get()
	if assert.NoError(t, s.getGroups(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[]\n", rec.Body.String())
	}
}

func TestAddGroup(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec, c := s.Post(strings.NewReader("{\"name\":\"name\"}"))
	_ = s.addGroups(c)

	rec, c = s.Get()
	if assert.NoError(t, s.getGroups(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[{\"id\":1,\"name\":\"name\"}]\n", rec.Body.String())
	}
}

func TestUpdateGroup(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	_, c := s.Post(strings.NewReader("{\"name\":\"name\"}"))
	_ = s.addGroups(c)
	rec, c := s.Get()
	if assert.NoError(t, s.getGroups(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[{\"id\":1,\"name\":\"name\"}]\n", rec.Body.String())
	}

	_, c = s.Put(strings.NewReader("{\"name\":\"new name\"}"))
	c.SetParamNames("groupId")
	c.SetParamValues("1")
	_ = s.updateGroup(c)

	rec, c = s.Get()
	if assert.NoError(t, s.getGroups(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[{\"id\":1,\"name\":\"new name\"}]\n", rec.Body.String())
	}
}
