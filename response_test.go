package hh

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matryer/is"
)

func TestSetJsonContentHeader(t *testing.T) {
	rr := httptest.NewRecorder()
	is := is.New(t)

	is.Equal(rr.Header().Get("Content-Type"), "")
	SetJsonContentHeader(rr)
	is.Equal(rr.Header().Get("Content-Type"), "application/json")
}

func TestSendJSON(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		rr := httptest.NewRecorder()
		is := is.New(t)
		type payload struct {
			Name  string `json:"name"`
			Error error  `json:"error,omitempty"`
		}
		err := SendJSON(rr, payload{Name: "jaga"})
		is.NoErr(err)
		is.Equal(rr.Header().Get("Content-Type"), "application/json")
		is.Equal(rr.Body.String(), `{"name":"jaga"}`)
		is.Equal(rr.Code, http.StatusOK)
	})

	t.Run("can set custom stats code", func(t *testing.T) {
		rr := httptest.NewRecorder()
		is := is.New(t)
		type payload struct {
			Name  string `json:"name"`
			Error error  `json:"error,omitempty"`
		}
		err := SendJSON(rr, payload{Name: "jaga"}, http.StatusCreated)
		is.NoErr(err)
		is.Equal(rr.Code, 201)
	})
}
func TestSendJSONObject(t *testing.T) {
	rr := httptest.NewRecorder()
	is := is.New(t)

	err := SendJSONObject(rr, JSONObject{
		"age": 28,
	})
	is.NoErr(err)
	is.Equal(rr.Code, http.StatusOK)
	is.Equal(rr.Header().Get("Content-Type"), "application/json")
	is.Equal(rr.Body.String(), `{"age":28}`)
}

func TestSendErrorMessage(t *testing.T) {
	t.Run("default to status code 500", func(t *testing.T) {
		rr := httptest.NewRecorder()
		is := is.New(t)
		err := SendErrorMessage(rr, "wrong!")
		is.NoErr(err)
		is.Equal(rr.Code, http.StatusInternalServerError)
		is.Equal(rr.Header().Get("Content-Type"), "application/json")
		is.Equal(rr.Body.String(), `{"error":"wrong!"}`)
	})

	t.Run("can set custom error code", func(t *testing.T) {
		rr := httptest.NewRecorder()
		is := is.New(t)
		err := SendErrorMessage(rr, "wrong!", http.StatusBadRequest)
		is.NoErr(err)
		is.Equal(rr.Code, 400)
		is.Equal(rr.Header().Get("Content-Type"), "application/json")
		is.Equal(rr.Body.String(), `{"error":"wrong!"}`)
	})
}
