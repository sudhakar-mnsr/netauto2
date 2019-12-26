package main

import "fmt"

func fib(n int) {
   x, y := 0, 1
   for i := 0; i < n; i++ {
      x, y = y, x+y
      fmt.Printf("%d, ", x)
   }
}

func main() {
   fib(100)
} 
