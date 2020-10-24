package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// HealthCheck API returns date time to client
func HealthCheck(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	io.WriteString(w, currentTime.String())
}
