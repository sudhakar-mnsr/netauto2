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

// Fetches the repos for the given Github users
func getStats(url string) *grequests.Response {
   resp, err := grequests.Get(url, requestOptions)
   // you can modify the request by passing an optional RequestOptions struct
   if err != nil {
      log.Fatalln("Unable to make request: ", err)
   }
   return resp
}

// Reads the files provided and creates Gist on github
func createGist(url string, args []string) *grequests.Response {
// get first two arguments
description := args[0]
// remaining arguments are filenames with path
var fileContents = make(map[string]File)
for i := 1; i < len(args); i++ {
   dat, err := ioutil.ReadFile(args[i])
   if err != nil {
      log.Println("Please check the filenames.Absolute path or same directory are allowed")
      return nil
   }
   
   var file File
   file.Content = string(dat)
   fileContents[args[i]] = file
}

func main() {
app := cli.NewApp()
// define command for our client
app.Commands = []cli.Command{
   {
      Name: "fetch",
      Aliases: []string{"f"}
      Usage: "Fetch the repo details with user. [Usage]: goTool fetch user")
      Action: func(c *cli.Context) error {
         if c.NArg() > 0 {
            // Github API Logic
            var repos []Repo
            user := c.Args()[0]
            var repoUrl = fmt.Sprintf("https://api.github.com/users/%s/repos", user)
            resp := getStats(repoUrl)
            resp.JSON(&repos)
            log.Println(repos)
         } else {
            log.Println("Please give a username. See -h to see help")
         }
         return nil
      },
   },
   { 
