package main

import (
   "database/sql"
   "fmt"
   _ "github.com/lib/pq"
   "log"
   "net/http"
   "time"
)

func main() {
   http.HandleFunc("/", func(w htp.ResponseWriter, r *http.Request) {
      db, err := sql.Open("postgres", "host=localhost dbname=sretest sslmode=disable")
      if err != nil {
         log.Fatal(err)
      }
      defer db.Close()
      
      stmt, err := db.Prepare("Insert into data (metric, value, created) values ($1, $2, $3)")
      if err != nil {
         log.Fatal(err)
      }
      defer stmt.Close()
