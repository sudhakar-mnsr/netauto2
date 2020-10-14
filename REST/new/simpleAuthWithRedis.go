package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	redistore "gopkg.in/boj/redistore.v1"
)

var strore, err = redistore.NewRediStore(10, "tcp", ":6379", "", []byte(os.Getenv("SESSION_SECRET")))
var users = map[string]string{"naren": "passme", "admin": "password"}

// HealthcheckHandler returns the date and time
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
