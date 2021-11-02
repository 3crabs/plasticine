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

func (s *server) addSubjectReq(subject models.Subject) {
	bytes, _ := json.Marshal(subject)
	_, c := s.post(strings.NewReader(string(bytes)))
	_ = s.addSubject(c)
}

func (s *server) getSubjectsReq() (*httptest.ResponseRecorder, []models.Subject) {
	rec, c := s.get()
	_ = s.getSubjects(c)
	var subjects []models.Subject
	_ = json.Unmarshal([]byte(rec.Body.String()), &subjects)
	return rec, subjects
}

func (s *server) updateSubjectReq(subjectId int, subject models.Subject) {
	bytes, _ := json.Marshal(subject)
	_, c := s.put(strings.NewReader(string(bytes)))
	c.SetParamNames("subjectId")
	c.SetParamValues(strconv.Itoa(subjectId))
	_ = s.updateSubject(c)
}

func (s *server) getSubjectReq(subjectId int) (*httptest.ResponseRecorder, models.Subject) {
	rec, c := s.get()
	c.SetParamNames("subjectId")
	c.SetParamValues(strconv.Itoa(subjectId))
	_ = s.getSubject(c)
	var subject models.Subject
	_ = json.Unmarshal([]byte(rec.Body.String()), &subject)
	return rec, subject
}

func TestGetSubjects(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec, subjects := s.getSubjectsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 0, len(subjects))
}

func TestAddSubject(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	subject := models.Subject{Name: "name"}
	s.addSubjectReq(subject)

	rec, subjects := s.getSubjectsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 1, len(subjects))
	assert.Equal(t, subject.Name, subjects[0].Name)
}

func TestUpdateSubject(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	subject := models.Subject{Name: "name"}
	s.addSubjectReq(subject)

	rec, subjects := s.getSubjectsReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 1, len(subjects))
	assert.Equal(t, subject.Name, subjects[0].Name)

	subject.Id = subjects[0].Id
	subject.Name = "new name"
	s.updateSubjectReq(1, subject)

	rec, subjects = s.getSubjectsReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 1, len(subjects))
	assert.Equal(t, subject.Name, subjects[0].Name)
}

func TestGetSubject(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	subject := models.Subject{Name: "name"}
	s.addSubjectReq(subject)
	_, subjects := s.getSubjectsReq()
	subject = subjects[0]

	_, subject = s.getSubjectReq(subject.Id)

	assert.Equal(t, "name", subject.Name)
}
