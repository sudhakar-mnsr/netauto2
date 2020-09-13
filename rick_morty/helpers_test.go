package rickandmorty

import "testing"

var slice = []string{
	"hello",
	"world",
	"hola",
	"mundo",
	"Allo",
	"Welt",
}

var sliceIntegers = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
}

func Test_containsStringElementPresent(t *testing.T) {
   sliceToIterate := []string{
      "hello",
      "hola",
      "Allo",
      "world",
      "mundo",
      "Welt",
   }
   
   for _, element := range sliceToIterate {
      if !containsString(slice, element) {
         t.Errorf("The %s (%T) is not present in %v", element, element, slice)
      }
   }
}
