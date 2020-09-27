package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sausheong/gwp/Chapter_2_Go_ChitChat/chitchat/data"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Configuration struct {
   Address string
   ReadTimeout int64
   WriteTimeout int64
   Static string
}

var config Configuration
var logger *log.Logger

// Convenience function for printing to stdout
func p(a ...interface{}) {
   fmt.Println(a)
}

func init() {
   loadConfig()
   file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
   if err != nil {
      log.Fatalln("Failed to open log file", err)
   }
   logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

func loadConfig() {
   file, err := os.Open("config.json")
   if err != nil {
      log.Fatalln("Cannot open config file", err)
   }
   
   decoder := json.NewDecoder(file)
   config = Configuration{}
   err = decoder.Decode(&config)
   if err != nil {
      log.Fatalln("Cannot get configuration from file", err)
   }
}

// Convenience function to redirect to the error message page
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
   url := []string{"/err?msg=", msg}
   http.Redirect(writer, request, strings.Join(url, ""), 302)
}

