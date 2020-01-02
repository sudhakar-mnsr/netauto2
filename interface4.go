package main

import (
   "fmt"
)

func main() {
   var anyType interface{}
   anyType = 77.0
   anyType = "Iam a string now"
   fmt.Println(anyType)

   printAnyType("The car is slow")
   m := map[string] string{"ID":"12345", "name":"Kerry"}

   printAnyType(m)
   printAnyType(1234567)
}

func printAnyType(val interface{}) {
   fmt.Println(val)
}
