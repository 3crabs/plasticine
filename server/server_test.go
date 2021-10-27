package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"plasticine/db"
	"strings"
	"testing"
)

var (
	groupJSON = `[{"name":"name"}]`
)

func TestGetGroups(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec, c := s.Get("/groups")
	if assert.NoError(t, s.getGroups(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[]\n", rec.Body.String())
	}
}

func TestAddGroup(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec, c := s.Post("/groups", strings.NewReader(groupJSON))
	if assert.NoError(t, s.addGroups(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

	rec, c = s.Get("/groups")
	if assert.NoError(t, s.getGroups(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "[{\"id\":1,\"name\":\"name\"}]\n", rec.Body.String())
	}
}
