package main

import (
   "encoding/json"
   "fmt"
   "io/ioutil"
)

type Post struct {
   Id int `json:"id"`
   Content string `json:"content"`
   Author Author `json:"author"`
   Comments []Comment `json:"comments"`
}
