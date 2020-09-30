package main

import (
   "encoding/json"
   "fmt"
   "io"
)

type Post struct {
   Id int `json:"id"`
   Content string `json:"content"`
   Author Author `json:"author"`
   Comments []Comment `json:"comments"`
}

type Author struct {
   Id int `json:"id"`
   Name string `json:"name"`
}

type Comment struct {
   Id int `json:"id"`
   Content string `json:"content"`
   Author string `json:"author"`
}

func main() {
post := Post{
   Id: 1,
   Content: "Hello World!",
   Author: Author{
      Id: 2,
      Name: "Sau Sheong",
   },
   Comments: []Comment{
      Comment{
         Id: 1,
         Content: "Have a great day!",
         Author: "Adam",
      },
      Comment{
         Id: 2,
         Content: "How are you today?",
         Author: "Betty",
      },
   },
}
