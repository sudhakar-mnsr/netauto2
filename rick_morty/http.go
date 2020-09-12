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
   
   case map[string]string:
      if k == "params" {
         hasParams = true
      }
   case map[string]int:
      if k == "params" {
         integer := strconv.FormatInt(int64(v.(map[string]int)["integer"]), 10)
         endpoint = endpoint + integer
         delete(options, "params")
      }
   case []int:
      params := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v.([]int))), ","), "[]")
      endpoint = endpoint + params
   default: 
      err := sliceIntToString(v.([]int), ",")
      return nil, errors.New(err)
   }
}

