package hh

import (
	"encoding/json"
	"net/http"
)

type JSONObject = map[string]interface{}

func SendJSON(w http.ResponseWriter, data interface{}, code ...int) error {
	SetJsonContentHeader(w)

	if len(code) > 0 {
		w.WriteHeader(code[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}

func SendErrorMessage(w http.ResponseWriter, errorMessage string, code ...int) error {
	SetJsonContentHeader(w)

	if len(code) > 0 {
		w.WriteHeader(code[0])
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	type jsonError struct {
		ErrorMessage string `json:"error"`
	}
	jsonBytes, err := json.Marshal(jsonError{ErrorMessage: errorMessage})
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}

func SendJSONObject(w http.ResponseWriter, o JSONObject) error {
	return SendJSON(w, o)
}

func SetJsonContentHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
