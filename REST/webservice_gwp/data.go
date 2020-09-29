package main

import ( 
   "database/sql"
   _ "github.com/lib/pq"
)     

var Db *sql.DB

// Connect to DB
func init() {
   var err error
   Db, err = sql.Open(postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
   if err != nil {
      panic(err)
   }
}
