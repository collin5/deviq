package test

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"qa"
	"strings"
	"sync"
	"testing"
)

var (
	expected = "working"
	wait     sync.WaitGroup
)

func Exec(q *qa.Query, w http.ResponseWriter) {
	defer wait.Done()
	io.Copy(w, strings.NewReader(expected))
}

func TestHandler(t *testing.T) {
	wait.Add(1)
	qa.Exec = Exec
	form := []byte(`{"lang": "python"}`)
	r := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(form))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	qa.Handle(w, r)
	wait.Wait()
	actual, _ := ioutil.ReadAll(w.Result().Body)
	if string(actual) != expected {
		t.Errorf("Expected %s, got %s instead.", expected, actual)
	}
}
