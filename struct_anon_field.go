package main

import "fmt"

type diameter int
type name struct {
   long string
   short string
   symbol rune
}
type plannet struct {
   diameter
   name
   desc string
}

func main() {
   earth := plannet{
      diameter: 7926,
      name: name{
         long: "Earth",
         short: "E",
         symbol: '\u2641',
      },
      desc: "Third rock from the sun",
   }
   fmt.Println(earth)
   jupiter := plannet{}
   jupiter.diameter = 88846
   jupiter.name.long = "Jupiter"
   jupiter.name.short = "J"
   jupiter.name.symbol = '\u2643'
   jupiter.desc = "A ball of gas"
   fmt.Println(jupiter)
   saturn := plannet{}
   saturn.diameter = 120536
   saturn.long = "Saturn"
   saturn.short = "S"
   saturn.symbol = '\u2644'
   saturn.desc = "Slow mover"
   fmt.Println(saturn)
}
