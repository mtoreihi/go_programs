package main

import (
	"net/http"
	"log"
)

func LicenseServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is the test result."))
}

func main() {
	http.HandleFunc("/getLicense", LicenseServer)
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
	//err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen and Server: ", err)
	}
}