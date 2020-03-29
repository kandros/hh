package hh

import "net/http"

func IsJsonRequest(r *http.Request) bool {
	return r.Header.Get("Content-Type") == "application/json"
}

func GetParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}
