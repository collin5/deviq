package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"../src/qa"
)

func TestHandler(t *testing.T) {
	form := []byte(`{"lang": "python"}`)
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(form))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	qa.Handle(w, r)
}

func TestQueryExec(t *testing.T) {

}
