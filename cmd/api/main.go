package main

import (
	"fmt"
	"log"
	"net/http"
)

const addr = "127.0.0.1"
const port = 8000

type application struct{}

func main() {
	// set application config
	var app application

	// read from command line

	// connect to db

	// declare server
	server := http.Server{
		Addr:    fmt.Sprintf("%v:%v", addr, port),
		Handler: app.Routes(),
	}

	// server execution log
	log.Printf("Executing server at %v:%v", addr, port)

	// start a web server
	err := server.ListenAndServe()

	// handle web server execution error
	if err != nil {
		panic(err)
	}
}
