package main

import (
   "fmt"
   "io"
   "net"
   "net/http"
   "time"
   "os"
)

func main() {
client := &http.Client{
           Transport: &http.Transport{
           DisableKeepAlives: true,
           Dial: (&net.Dialer{
              Timeout: 30 * time.Second,
           },
}       
