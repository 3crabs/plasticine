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

func TestGetGroups(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec, groups := s.getGroupsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 0, len(groups))
}

func TestAddGroup(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	group := models.Group{Name: "name"}
	s.addGroupsReq(group)

	rec, groups := s.getGroupsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, 1, len(groups))
	assert.Equal(t, group.Name, groups[0].Name)
}

func TestUpdateGroup(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

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
