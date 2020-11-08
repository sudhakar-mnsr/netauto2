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
