package main

import (
   "fmt"
   "encoding/json"
)

type Person struct { 
   Name    string `json:"person_name"` 
   Title   string `json:"person_title"` 
   Address `json:"person_address_obj"` 
} 
 
type Address struct { 
   Street string `json:"person_addr_street"` 
   City   string `json:"person_city"` 
   State  string `json:"person_state"` 
   Postal string `json:"person_postal_code"` 
} 

func main() { 
   p := Person{ 
         Name: "Vladimir Vivien", 
         Title : "Author", 
   } 
   b, _ := json.Marshal(p) 
   fmt.Println(string(b)) 
}
