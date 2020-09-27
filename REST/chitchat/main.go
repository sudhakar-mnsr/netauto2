package main

import (
   "net/http"
   "time"
)

func main() {
p("ChitChat", version(), "started at", config.Address)

// handle static assets
mux := http.NewServeMux()
files := http.FileServer(http.Dir(config.Static))
mux.Handle("/static/", http.StripPrefix("/static/", files))

// all route patterns matched here
// route handler functions defined in other files

// index
mux.HandleFunc("/", index)
// error
mux.HandleFunc("/err", err)

// defined in route_thread.go
mux.HandleFunc("/thread/new", newThread)
mux.HandleFunc("/thread/create",createThread)
mux.HandleFunc("/thread/post",postThread)
mux.HandleFunc("/thread/read",readThread)
