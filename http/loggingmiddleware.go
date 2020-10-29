package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request!")
	w.Write([]byte("OK"))
	log.Println("Finished processing request")
}
