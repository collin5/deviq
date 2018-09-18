package main

import (
	"net/http"
	"qa"
)

func main() {
	http.HandleFunc("/", qa.Handle)
	http.ListenAndServe(":3000", nil)
}
