package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"plasticine/db"
	"plasticine/models"
	"strings"
)

type server struct {
	router *echo.Echo
	db     db.DB
	port   string
}

func NewServer(port string, db db.DB) *server {
	s := &server{
		router: echo.New(),
		db:     db,
		port:   port,
	}
	s.routes()
	return s
}

func (s *server) Get(route string) (*httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(http.MethodGet, route, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return rec, s.router.NewContext(req, rec)
}

func (s *server) Post(route string, body *strings.Reader) (*httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(http.MethodPost, route, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return rec, s.router.NewContext(req, rec)
}

func (s *server) routes() {
	s.router.POST("/groups", s.addGroups)
	s.router.GET("/groups", s.getGroups)
}

func (s *server) addGroups(c echo.Context) error {
	var groups []models.Group
	if err := c.Bind(&groups); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	err := s.db.AddGroups(groups)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "new groups added")
}

func (s *server) getGroups(c echo.Context) error {
	return c.JSON(http.StatusOK, s.db.GetGroups())
}

func (s *server) Run() {
	s.router.Logger.Fatal(s.router.Start(s.port))
}
