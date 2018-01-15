package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

func hhf_index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a test for index page.")
}

func hhf_getCount(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("companies")
	cnt, err := c.Count()
	if err != nil {
		log.Fatal(err)
	}

	s := cnt

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
}
