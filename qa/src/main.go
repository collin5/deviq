package main

import (
	"encoding/json"
	"io"
	"net/http"
	"util"
)

type Query struct {
	Lang string `json:"lang"`
	Code string `json:"code"`
	Spec string `json:"spec"`
}

func (q *Query) Exec(w http.ResponseWriter) {
	cont := util.NewContainer("blahblah", q.Lang)
	out, err := cont.Run(q.Code, q.Spec)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.Copy(w, out)
}

func handle(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var query Query
	err := dec.Decode(&query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	query.Exec(w)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":3000", nil)
}
