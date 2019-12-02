package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestIsDivisor(t *testing.T) {
	router := setupRouter()
	ping := performRequest(router, "GET", "/ping", nil)
	if ping.Code != http.StatusOK {
		t.Errorf("/ping failed; got: %v, want: %v", ping.Code, http.StatusOK)
	}
	req := DIVISION{
		DIVIDEND: 9,
		DIVISOR:  3,
	}
	correctAnswer := DIVISION{
		DIVISOR:  9,
		DIVIDEND: 3,
		RESULT:   true,
	}
	body, err := json.Marshal(req)
	if err != nil {
		t.Error(err)
	}
	api := performRequest(router, "POST", "/api/v1/", bytes.NewBuffer(body))
	if api.Code != http.StatusOK {
		t.Errorf("/ping failed; got: %v, want: %v", api.Code, http.StatusOK)
	}
	var answer DIVISION
	err = json.Unmarshal(api.Body.Bytes(), &answer)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(correctAnswer, answer) {
		t.Errorf("Not equal answer for: %v; got: %v, want: %v", req, answer, correctAnswer)
	}
}
