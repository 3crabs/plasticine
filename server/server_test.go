package server

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"plasticine/db"
	"strconv"
	"strings"
	"testing"
)

func (s *server) get() (*httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return rec, s.router.NewContext(req, rec)
}

func (s *server) post(body *strings.Reader) (*httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(http.MethodPost, "/", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return rec, s.router.NewContext(req, rec)
}

func (s *server) put(body *strings.Reader) (*httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(http.MethodPut, "/", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return rec, s.router.NewContext(req, rec)
}

func (s *server) addGroupsReq(groupJSON string) {
	_, c := s.post(strings.NewReader(groupJSON))
	_ = s.addGroups(c)
}

func (s *server) getGroupsReq() *httptest.ResponseRecorder {
	rec, c := s.get()
	_ = s.getGroups(c)
	return rec
}

func (s *server) updateGroupReq(groupId int, groupJSON string) {
	_, c := s.put(strings.NewReader(groupJSON))
	c.SetParamNames("groupId")
	c.SetParamValues(strconv.Itoa(groupId))
	_ = s.updateGroup(c)
}

func TestGetGroups(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	rec := s.getGroupsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[]\n", rec.Body.String())
}

func TestAddGroup(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	s.addGroupsReq("{\"name\":\"name\"}")

	rec := s.getGroupsReq()

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[{\"id\":1,\"name\":\"name\"}]\n", rec.Body.String())
}

func TestUpdateGroup(t *testing.T) {
	s := NewServer(":8080", db.NewDB())

	s.addGroupsReq("{\"name\":\"name\"}")

	rec := s.getGroupsReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[{\"id\":1,\"name\":\"name\"}]\n", rec.Body.String())

	s.updateGroupReq(1, "{\"name\":\"new name\"}")

	rec = s.getGroupsReq()
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "[{\"id\":1,\"name\":\"new name\"}]\n", rec.Body.String())
}
