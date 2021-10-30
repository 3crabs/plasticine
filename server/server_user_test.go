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

func (s *server) getStudentsReq() (*httptest.ResponseRecorder, []models.User) {
	rec, c := s.get()
	_ = s.getStudents(c)
	var users []models.User
	_ = json.Unmarshal([]byte(rec.Body.String()), &users)
	return rec, users
}

func (s *server) getTeachersReq() (*httptest.ResponseRecorder, []models.User) {
	rec, c := s.get()
	_ = s.getTeachers(c)
	var users []models.User
	_ = json.Unmarshal([]byte(rec.Body.String()), &users)
	return rec, users
}

func (s *server) addUserReq(user models.User) {
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, c := s.post(strings.NewReader(string(bytes)))
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

	rec, students := s.getStudentsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 0, len(students))
}

func TestGetTeachers(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec, teachers := s.getTeachersReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 0, len(teachers))
}

func TestAddStudent(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	student := models.User{
		LastName:  "lastname",
		FirstName: "firstname",
		Role:      models.Student,
	}
	s.addUserReq(student)

	rec, students := s.getStudentsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, student.LastName, students[0].LastName)
	assert.Equal(t, student.FirstName, students[0].FirstName)
	assert.Equal(t, student.Role, students[0].Role)
}

func TestAddTeacher(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	teacher := models.User{
		LastName:  "lastname",
		FirstName: "firstname",
		Role:      models.Teacher,
	}
	s.addUserReq(teacher)

	rec, teachers := s.getTeachersReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, teacher.LastName, teachers[0].LastName)
	assert.Equal(t, teacher.FirstName, teachers[0].FirstName)
	assert.Equal(t, teacher.Role, teachers[0].Role)
}

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
