package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Println("Starting the server...")
	log.Fatal(http.ListenAndServe(":8443", router))

	//log.Fatal(http.ListenAndServeTLS(":8443", "certificate.pem", "key.pem", router))

}
