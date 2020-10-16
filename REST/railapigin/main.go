package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter4/dbutils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// DB Driver visible to whole program
var DB *sql.DB
