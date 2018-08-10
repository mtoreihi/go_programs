package main

import (
	"log"
	"github.com/kardianos/service"
	"net/http"
)

func LicenseServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is the test result."))
}

func myserver() {
	http.HandleFunc("/getLicense", LicenseServer)
	//err := http.ListenAndServeTLS(":8443", "F:/Go/go_programs/windows_service/server.crt", "F:/Go/go_programs/windows_service/server.key", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
}

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	myserver()
}


func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "ParsRightLicenseManager",
		DisplayName: "ParsRight License Manager",
		Description: "License manager for ParsRight.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}