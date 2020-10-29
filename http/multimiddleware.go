package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area uint64
}

func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		// Filtering requests by MIME type
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}
