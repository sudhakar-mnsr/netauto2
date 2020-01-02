package main

import "fmt"

type food interface {
   eat()
}

type veggie string
func (v veggie) eat() {
   fmt.Println("Eating", v)
}

type meat string
func (m meat) eat() {
   fmt.Println("Eating tasty", m)
}

func eat(f food) {
   veg, ok := f.(veggie)
   if ok {
      if veg == "okra" {
         fmt.Println("Yuk! not eating", veg)
      }else{
         veg.eat()
      }
      return
   }
   mt, ok := f.(meat)
   if ok {
      if mt == "beef" {
         fmt.Println("Yuk! not eating", mt)
      }else{
         mt.eat()
      }
      return
   }
   fmt.Println("Not eating whatever that is: ", f)
}

func main() {
   m1 := meat("beef")
   m2 := meat("mutton")
   v1 := veggie("okra")
   v2 := veggie("cabbage")
   eat(m1)
   eat(m2)
   eat(v1)
   eat(v2)
}

