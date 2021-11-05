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

func (s *server) addGroupsReq(group models.Group) {
	bytes, _ := json.Marshal(group)
	_, c := s.post(strings.NewReader(string(bytes)))
	_ = s.addGroup(c)
}

func (s *server) getGroupsReq() (*httptest.ResponseRecorder, []models.Group) {
	rec, c := s.get()
	_ = s.getGroups(c)
	var groups []models.Group
	_ = json.Unmarshal([]byte(rec.Body.String()), &groups)
	return rec, groups
}

func (s *server) updateGroupReq(groupId int, group models.Group) {
	bytes, _ := json.Marshal(group)
	_, c := s.put(strings.NewReader(string(bytes)))
	c.SetParamNames("groupId")
	c.SetParamValues(strconv.Itoa(groupId))
	_ = s.updateGroup(c)
}

func (s *server) deleteGroupReq(groupId int) {
	_, c := s.delete()
	c.SetParamNames("groupId")
	c.SetParamValues(strconv.Itoa(groupId))
	_ = s.deleteGroup(c)
}

func (s *server) getGroupStudentsReq(groupId int) (*httptest.ResponseRecorder, []models.User) {
	rec, c := s.get()
	c.SetParamNames("groupId")
	c.SetParamValues(strconv.Itoa(groupId))
	_ = s.getGroupStudents(c)
	var students []models.User
	_ = json.Unmarshal([]byte(rec.Body.String()), &students)
	return rec, students
}

func (s *server) getGroupReq(groupId int) (*httptest.ResponseRecorder, models.GroupInfo) {
	rec, c := s.get()
	c.SetParamNames("groupId")
	c.SetParamValues(strconv.Itoa(groupId))
	_ = s.getGroup(c)
	var group models.GroupInfo
	_ = json.Unmarshal([]byte(rec.Body.String()), &group)
	return rec, group
}

func TestGetGroups(t *testing.T) {
	s := NewServer(":8080", db.NewDefaultDb())

	rec, groups := s.getGroupsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 0, len(groups))
}

func TestAddGroup(t *testing.T) {
	s := NewServer(":8080", db.NewDefaultDb())

	group := models.Group{Name: "name"}
	s.addGroupsReq(group)

	rec, groups := s.getGroupsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 1, len(groups))
	assert.Equal(t, group.Name, groups[0].Name)
}

func TestGetGroup(t *testing.T) {
	s := NewServer(":8080", db.NewDefaultDb())

	group := models.Group{Name: "name"}
	s.addGroupsReq(group)
	_, groups := s.getGroupsReq()

	_, groupInfo := s.getGroupReq(groups[0].Id)

	assert.Equal(t, groupInfo.Id, groups[0].Id)
	assert.Equal(t, groupInfo.Name, groups[0].Name)
	assert.NotNil(t, groupInfo.Lessons)
	assert.Equal(t, 0, len(groupInfo.Lessons))
}

func TestUpdateGroup(t *testing.T) {
	s := NewServer(":8080", db.NewDefaultDb())

	group := models.Group{Name: "name"}
	s.addGroupsReq(group)

	rec, groups := s.getGroupsReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 1, len(groups))
	assert.Equal(t, group.Name, groups[0].Name)

	group.Id = groups[0].Id
	group.Name = "new name"
	s.updateGroupReq(group.Id, group)

	rec, groups = s.getGroupsReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 1, len(groups))
	assert.Equal(t, group.Name, groups[0].Name)
}

func TestDeleteGroup(t *testing.T) {
	s := NewServer(":8080", db.NewDefaultDb())

	group := models.Group{Name: "name"}
	s.addGroupsReq(group)
	_, groups := s.getGroupsReq()
	group = groups[0]

	_, groups = s.getGroupsReq()
	assert.Equal(t, 1, len(groups))

	s.deleteGroupReq(group.Id)

	_, groups = s.getGroupsReq()
	assert.Equal(t, 0, len(groups))
}

func TestGetGroupStudents(t *testing.T) {
	s := NewServer(":8080", db.NewDefaultDb())

	group := models.Group{Name: "name"}
	s.addGroupsReq(group)
	_, groups := s.getGroupsReq()
	group = groups[0]

	student := models.User{
		LastName:  "lastname",
		FirstName: "firstname",
		Role:      models.Student,
		GroupId:   group.Id,
	}
	s.addUserReq(student)
	_, students := s.getStudentsReq()
	student = students[0]

	_, students = s.getGroupStudentsReq(group.Id)
	assert.Equal(t, 1, len(students))
}
