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

// main handler function
func handleRequest(w http.ResponseWriter, r *http.Request) {
   var err error
   switch r.Method {
   case "GET":
      err = handleGet(w, r)
   case "POST":
      err = handlePost(w, r)
   case "PUT":
      err = handlePut(w, r)
   case "DELETE": 
      err = handleDelete(w, r)
   }
   if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
   }
}
