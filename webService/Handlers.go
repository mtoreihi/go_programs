package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/gorilla/mux"
)

func hhf_index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a test.")
}

func hhf_getCount(w http.ResponseWriter, r *http.Request) {
	s := "{ count : '25'}"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}
