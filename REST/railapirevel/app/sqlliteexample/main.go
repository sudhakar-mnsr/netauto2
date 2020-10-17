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

func dbOperations(db *sql.DB) {
// Create
statement, _ := db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
statement.Exec("A Tale of two cities", "Charless Dickens", 140430547)
log.Println("Inserted the book into database!")

// Read
rows, _ := db.Query("SELECT id, name, author FROM books")
var tempBook Book
