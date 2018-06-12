package main

import (
	"net/http"
	"log"
	"github.com/patrickmn/go-cache"
	"time"
)

var (
	c = cache.New(10 * time.Second, 10 * time.Second)
)

func LicenseServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is the test result."))
	w.Write([]byte(provideData(req)))
}

func main() {
	http.HandleFunc("/getLicense", LicenseServer)
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
	//err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen and Server: ", err)
	}
}

func provideData(req *http.Request) string {
	data, found := c.Get("data")
	println(time.Now().String())
	if found {
		println("reading from cache")
		return data.(string)
	} else {
		println("rebuilding cache ..............................")
		c.Set("data", "123456789", cache.DefaultExpiration)
		data, found := c.Get("data")
		if found {
			return data.(string)
		} else {
			return ""
		}
	}
}