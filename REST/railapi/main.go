package main

import (
   "database/sql"
   "encoding/json"
   "log"
   "net/http"
   "time"
   "github.com/Hands-On-Restful-Web-services-with-Go/chapter4/dbutils"
   "github.com/emicklei/go-restful"
   _ "github.com/mattn/go-sqlite3"
)

// DB Driver visible to whole program
var DB *sql.DB

// TrainResource is the model for holding rail information

