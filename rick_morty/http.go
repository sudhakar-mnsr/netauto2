package rickandmorty

import (
   "encoding/json"
   "errors"
   "fmt"
   "net/http"
   "strconv"
   "strings"
   "time"
)

func makePetition(options map[string]interface{}) (interface{}, error) {
client := &http.Client{
          Timeout: 5 * time.Second,
}

endpoint := ""
hasParams := false

for k, v := range options {
switch v.(type) {
case string:
   if k == "endpoint" {
      endpoint = v.(string)
   }

