package main

import "fmt"

type person struct {
   name string
   title string
}

func updateName(p person, name string) {
   p.name = name
}

func main() {
   p := person{}
   p.name = "unknown"
   updateName(p, "Sudhakar MNSR")
   fmt.Println(p.name)
}
