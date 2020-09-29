package main

import ( 
   "database/sql"
   _ "github.com/lib/pq"
)     

var Db *sql.DB

// Connect to DB
func init() {
   var err error
   Db, err = sql.Open(postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
   if err != nil {
      panic(err)
   }
}

// Get a single post
func retrieve(id int) (post Post, err error) {
   post = Post{}
   err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
   return
}

// Create a new post
func (post *Post) create() (err error) {
   statement := "insert into posts (content, author) values ($1, $2) returning id"
   stmt, err := Db.Prepare(statement)
   if err != nil {
      return
   }
   defer stmt.Close()
   err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
   return
}
