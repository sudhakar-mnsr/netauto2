package main

import (
   "database/sql"
   "encoding/json"
   "log"
   "net/http"
   "time"
   "github.com/emicklei/go-restful"
   _ "github.com/mattn/go-sqlite3"
   "github.com/sudhakar-mnsr/dbutils"
)

func main() {
   db, err :=sql.Open("sqlite3", "./railapi.db")
   if err != nil {
      log.Println("Driver creation failed!")
   }
   dbutils.Initialize(db)
}
