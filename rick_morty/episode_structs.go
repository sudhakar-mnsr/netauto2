package rickandmorty

import (
"errors"
"strconv"
)

type Episode struct {
   ID int `json:"id"`
   Name string `json:"name"`
   Air_Date string `json:"air_date"`
   Episode string `json:"episode"`
   Characters []string `json:"characters"`
   URL string `json:"url"`
   Created string `json:"created"`
}
