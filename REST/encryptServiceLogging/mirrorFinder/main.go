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

func findFastest(urls []string) response {
urlChan := make(chan string)
latencyChan := make(chan time.Duration)

for _, url := range urls {
mirrorURL := url
go func() {
   log.Println("Started probing: ", mirrorURL)
   start := time.Now()
