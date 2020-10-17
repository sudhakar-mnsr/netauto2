package main

import (
   "database/sql"
   "log"
   _ "github.com/mattn/go-sqlite3"
)

// Book is a placeholder for book
type Book struct {
id int
name string
author string
}

