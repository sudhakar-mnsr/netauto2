package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
   Id int `json:"id"`
   Content string `json:"content"`
   Author string `json:"author"`
}
