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
