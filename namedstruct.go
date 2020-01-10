package main

import "fmt"

type person struct {
   name string
   address address
}
type address struct {
   street string
   city, state string
   postal string
}

func makePerson() person {
   addr := address{"Peeramchervu", "Hyderabad", "Telangana", "500091"}
   return person{
      name: "Sudhakar MNSR",
      address: addr,
   }
}
func main() {
   person1 := makePerson()
   fmt.Println(person1)
}
