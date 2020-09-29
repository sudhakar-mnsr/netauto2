package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
   Id int `json:"id"`
   Content string `json:"content"`
   Author string `json:"author"`
}

func main() {
   server := http.Server{
      Addr: ":8080",
   }
   http.HandleFunc("/post/", handleRequest)
   server.ListenAndServe()
}
