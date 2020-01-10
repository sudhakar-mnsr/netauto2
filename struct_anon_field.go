package main

import "fmt"

type diameter int
type name struct {
   long string
   short string
   symbol rune
}

func main() {
   earth := plannet{
      diameter: 7926
      name: name{
         long: "Earth",
         short: "E",
         symbol: '\u2641',
      },
      desc: "Third rock from the sun",
   }
   fmt.Println(earth)
}
