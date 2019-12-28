package main

import "fmt"

func f(...int) {}
func g([]int) {}

func main() {
   fmt.Printf("%T\n", f)
   fmt.Printf("%T\n", g)
}
