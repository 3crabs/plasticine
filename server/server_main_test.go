package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
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

func (s *server) delete() (*httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return rec, s.router.NewContext(req, rec)
}
