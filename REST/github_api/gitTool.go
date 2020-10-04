package main

import (
   "encoding/json"
   "fmt"
   "github.com/levigross/grequests"
   "github.com/urfave/cli"
   "io/ioutil"
   "log"
   "os"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}

