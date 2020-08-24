package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
rand.Seed(time.Now().Unix())
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
data := map[string]float64{
        "hello": 1.0
        "now": float64(time.Now().Unix()),
        "rand": rand.Float64(),
}


