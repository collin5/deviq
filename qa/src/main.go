package main

import (
	"net/http"

	"./qa"
)

func main() {
	http.HandleFunc("/", qa.HandleFunc)
	http.ListenAndServe(":3000", nil)
}
