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
	bytes, _ := json.Marshal(user)
	_, c := s.post(strings.NewReader(string(bytes)))
	_ = s.addUser(c)
}

func (s *server) updateUserReq(userId int, user models.User) {
	bytes, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, c := s.put(strings.NewReader(string(bytes)))
	c.SetParamNames("userId")
	c.SetParamValues(strconv.Itoa(userId))
	_ = s.updateUser(c)
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

func TestUpdateStudent(t *testing.T) {
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

	student.Id = students[0].Id
	student.FirstName = "new firstname"
	student.LastName = "new lastname"
	s.updateUserReq(student.Id, student)

	rec, students = s.getStudentsReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, student.LastName, students[0].LastName)
	assert.Equal(t, student.FirstName, students[0].FirstName)
	assert.Equal(t, student.Role, students[0].Role)
}

func TestUpdateTeacher(t *testing.T) {
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

	teacher.Id = teachers[0].Id
	teacher.FirstName = "new firstname"
	teacher.LastName = "new lastname"
	s.updateUserReq(teacher.Id, teacher)

	rec, teachers = s.getTeachersReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, teacher.LastName, teachers[0].LastName)
	assert.Equal(t, teacher.FirstName, teachers[0].FirstName)
	assert.Equal(t, teacher.Role, teachers[0].Role)
}

func TestAddGroupForStudent(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	group := models.Group{Name: "name"}
	s.addGroupsReq(group)
	_, groups := s.getGroupsReq()
	group = groups[0]

	student := models.User{
		LastName:  "lastname",
		FirstName: "firstname",
		Role:      models.Student,
	}
	s.addUserReq(student)

	_, students := s.getStudentsReq()
	assert.Equal(t, 0, students[0].GroupId)

	student = students[0]
	student.GroupId = group.Id
	s.updateUserReq(student.Id, student)

	_, students = s.getStudentsReq()
	assert.Equal(t, group.Id, students[0].GroupId)
}
