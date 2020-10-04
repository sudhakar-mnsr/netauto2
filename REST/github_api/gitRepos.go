package main

import (
   "github.com/levigross/grequests"
   "log"
   "os"
)

var GITHUB_TOKEN = os.Getenv("GITHUB_TOKEN")
var requestOptions = &grequests.RequestOptions{Auth: []string{GITHUB_TOKEN, "x-oauth-basic"}}

type Repo struct {
   ID int `json:"id"`
   Name string `json:"name"`
   FullName string `json:"full_name"`
   Forks int `json:"forks"`
   Private bool `json:"private"`
}
