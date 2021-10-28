package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"plasticine/db"
	"strconv"
	"strings"
	"testing"
)

func (s *server) addSubjectReq(subjectJSON string) {
	_, c := s.post(strings.NewReader(subjectJSON))
	_ = s.addSubject(c)
}

func (s *server) getSubjectsReq() *httptest.ResponseRecorder {
	rec, c := s.get()
	_ = s.getSubjects(c)
	return rec
}

func (s *server) updateSubjectReq(subjectId int, subjectJSON string) {
	_, c := s.put(strings.NewReader(subjectJSON))
	c.SetParamNames("subjectId")
	c.SetParamValues(strconv.Itoa(subjectId))
	_ = s.updateSubject(c)
}

func TestGetSubjects(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec := s.getSubjectsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[]\n", rec.Body.String())
}

func TestAddSubject(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	s.addSubjectReq("{\"name\":\"name\"}")

	rec := s.getSubjectsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[{\"id\":1,\"name\":\"name\"}]\n", rec.Body.String())
}

func TestUpdateSubject(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	s.addSubjectReq("{\"name\":\"name\"}")

	rec := s.getSubjectsReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[{\"id\":1,\"name\":\"name\"}]\n", rec.Body.String())

	s.updateSubjectReq(1, "{\"name\":\"new name\"}")

	rec = s.getSubjectsReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[{\"id\":1,\"name\":\"new name\"}]\n", rec.Body.String())
}
