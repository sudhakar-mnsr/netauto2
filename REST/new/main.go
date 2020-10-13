package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Book holds data of a book
type Book struct {
   ID int
   ISBN string
   Author string
   PublishedYear string
}

func main() {
// File open for reading, writing and appending
f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
if err != nil {
   fmt.Printf("error opening file: %v", err)
}
defer f.Close()

// This attaches program logs to file
log.SetOutput(f)

