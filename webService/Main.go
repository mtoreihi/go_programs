package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	//log.Fatal(http.ListenAndServe(":8080", router))
	log.Println("Starting the server...")
	log.Fatal(http.ListenAndServeTLS(":8443", "certificate.pem", "key.pem", router))
}
