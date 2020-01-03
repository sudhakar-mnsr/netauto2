package main

import "fmt"

func div(dvdn, dvsr int) (q, r int) {
   r = dvdn
   for r >= dvsr {
      q++
      r = r - dvsr
   }
   return 
}

func add(x, y int) (res int) {
   res = x + y
   return
}
func sub(x, y int) (res int) {
   res = x - y
   return
}
func mul(x, y int) (res int) {
   res = x * y
   return
}

func main() {
   var opAdd func(int, int) int = add
   opSub := sub
   opDiv := div
   var opMul func(int, int) int = mul
   fmt.Printf("opAdd=%d\n", opAdd(12,44))
   fmt.Printf("opSub=%d\n", opSub(12,44))
   x,y := opDiv(122,44)
   fmt.Printf("opDiv=%d-%d\n", x,y)
   fmt.Printf("opMul=%d\n", opMul(12,44))
}
