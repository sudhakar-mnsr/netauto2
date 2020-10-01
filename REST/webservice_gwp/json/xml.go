package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
post := Post{
   Id: "1",
   Content: "Hello World!",
   Author: Author{
      Id: "2",
      Name: "Sau Sheong",
   },
} 
