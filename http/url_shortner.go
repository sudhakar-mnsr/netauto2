package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter7/urlshortener/helper"
	base62 "github.com/Hands-On-Restful-Web-services-with-Go/chapter7/urlshortener/utils"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// DBClient stores the database session information. Needs to be initialized once
type DBClient struct {
	db *sql.DB
}

// Record Model is a HTTP response
type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}
