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
	id      int    `json: "id"`
	content string `json: "content"`
}

func hhf_getCount2(w http.ResponseWriter, r *http.Request) {

	//j := "{id:1,content:Hello, World!}"
	//m := t_count{1345, "Hello World!"}

	b := []byte(`{"id":999999,"content":"Hello World!"}`)

	//BSON, _ := json.NewEncoder(w)
	//log.Println(BSON)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Credentials", "(optional)")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
