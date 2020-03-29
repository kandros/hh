package hh

import (
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestIsJsonRequest(t *testing.T) {
	is := is.New(t)
	r := httptest.NewRequest("get", "/users?limit=10", nil)
	is.Equal(GetParam(r, "limit"), "10")
	is.Equal(GetParam(r, "nonexistant"), "")
}
