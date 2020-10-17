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
   for rows.Next() {
      rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
      log.Printf("ID:%d, Book:%s, Author: %s\n", tempBook.id, tempBook.name, tempBook.author)
   }
   
   // Update
   statement, _ = db.Prepare("update books set name=? where id=?")
   statement.Exec("The tale of two cities", 1)
   log.Println("Successfully updated the book in database!)
   
   // Delete
   statement, _ = db.Prepare("delete from books where id=?"
   statement.Exec(1)
   log.Println("Successfully deleted the book in database!")
}
