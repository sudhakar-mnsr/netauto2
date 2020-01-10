package main

import "fmt"

type person struct {
   name string
   title string
}

func updateNamePtr(p *person, name string) {
   p.name = name
}
func updateName(p person, name string) {
   p.name = name
}

func main() {
   p := person{}
   p.name = "unknown"
   updateName(p, "Sudhakar MNSR")
   fmt.Println(p.name)
   p_ptr := new(person)
   p_ptr.name = "unknown"
   updateNamePtr(p_ptr, "Sudhakar MNSR")
   fmt.Println(p_ptr.name)
}
