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

func (s *server) getStudentsReq() *httptest.ResponseRecorder {
	rec, c := s.get()
	_ = s.getStudents(c)
	return rec
}

func (s *server) getTeachersReq() *httptest.ResponseRecorder {
	rec, c := s.get()
	_ = s.getTeachers(c)
	return rec
}

func (s *server) addUserReq(userJSON string) {
	_, c := s.post(strings.NewReader(userJSON))
	_ = s.addUser(c)
}

func (s *server) updateUserReq(userId int, userJSON string) {
	_, c := s.put(strings.NewReader(userJSON))
	c.SetParamNames("userId")
	c.SetParamValues(strconv.Itoa(userId))
	_ = s.updateGroup(c)
}

func TestGetStudents(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec := s.getStudentsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[]\n", rec.Body.String())
}

func TestGetTeachers(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec := s.getTeachersReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[]\n", rec.Body.String())
}

//func TestAddStudent(t *testing.T) {
//	s := NewServer(":8080", db.NewDB())
//
//	s.addUserReq("{\"name\":\"name\"}")
//
//	rec := s.getGroupsReq()
//
//	assert.Equal(t, http.StatusOK, rec.Code)
//	assert.Equal(t, "[{\"id\":1,\"name\":\"name\"}]\n", rec.Body.String())
//}

//func TestUpdateGroup(t *testing.T) {
//	s := NewServer(":8080", db.NewDB())
//
//	s.addGroupsReq("{\"name\":\"name\"}")
//
//	rec := s.getGroupsReq()
//	assert.Equal(t, http.StatusOK, rec.Code)
//	assert.Equal(t, "[{\"id\":1,\"name\":\"name\"}]\n", rec.Body.String())
//
//	s.updateGroupReq(1, "{\"name\":\"new name\"}")
//
//	rec = s.getGroupsReq()
//	assert.Equal(t, http.StatusOK, rec.Code)
//	assert.Equal(t, "[{\"id\":1,\"name\":\"new name\"}]\n", rec.Body.String())
//}
