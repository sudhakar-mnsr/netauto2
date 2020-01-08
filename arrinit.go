package main

import "fmt"

type galaxies [14]string

func main() {
   var val [100]int = [100]int{1,2,3,4,5}
   for _, i := range val {
      fmt.Println(i)
   }
   var truth = [256]bool{true}
   for _, i := range truth {
      fmt.Println(i)
   }
   var histogram = [5]map[string]int {
      map[string]int{"A":12, "B":5, "D":11},
      map[string]int{"man":1344, "women":844, "children":577},
   }
   for _, i := range histogram {
      fmt.Println(i)
   }
   var board = [4][2]int{ 
      {33, 23}, 
      {62, 2}, 
      {23, 4}, 
      {51, 88}, 
   }
   for _, i := range board {
      fmt.Println(i)
   }
   // The literal value can be indexed, to make sure only some needed are
   // initialized others being 0
   var msg = [12]rune{0: 'H', 2: 'E', 4: 'L', 6: 'O', 8: '!'}
   for _, i := range msg {
      fmt.Println(i)
   }
   seven := [7]string{"grumpy", "sleepy", "bashful"} 
   fmt.Println(len(seven), cap(seven))
   namedGalaxies := &galaxies{ 
         "Andromeda", 
         "Black Eye", 
         "Bode's", 
   }
   fmt.Println(namedGalaxies)
   fmt.Println(*namedGalaxies)
}
