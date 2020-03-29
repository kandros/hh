package hh

import (
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestGetParam(t *testing.T) {
	is := is.New(t)
	req := httptest.NewRequest("", "/users?limit=10", nil)
	is.Equal(GetParam(req, "limit"), "10")
	is.Equal(GetParam(req, "nonexistant"), "")
}

func TestIsJsonRequest(t *testing.T) {
	is := is.New(t)
	req := httptest.NewRequest("", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	is.True(IsJsonRequest(req))
}
