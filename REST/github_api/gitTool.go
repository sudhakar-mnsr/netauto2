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

// Struct for holding response of repos fetch API
type Repo struct {
   ID int `json:"id"`
   Name string `json:"name"`
   FullName string `json:"full_name"`
   Forks int `json:"forks"`
   Private bool `json:"private"`
}

// Structs for modelling JSON body in create Gist
type File struct {
   Content string `json:"content"`
}

type Gist struct {
   Description string `json:"description"`
   Public bool `json:"public"`
   Files map[string]File `json:"files"`
}


