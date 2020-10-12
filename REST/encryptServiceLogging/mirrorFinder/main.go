package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter1/mirrors"
)

type response struct {
   FastestURL string `json:"fastest_url"`
   Latency time.Duration `json:"latency"`
}
