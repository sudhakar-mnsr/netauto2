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


