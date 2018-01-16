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

type t_count struct {
	name  string `json: "name"`
	count int    `json: "count"`
}

func hhf_getCount2(w http.ResponseWriter, r *http.Request) {

	j := "{count:13000}"

	//BSON, _ := json.NewEncoder(w)
	//log.Println(BSON)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(j); err != nil {
		panic(err)
	}
}
