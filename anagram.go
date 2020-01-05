package main 
 
import ( 
   "bufio" 
   "bytes" 
   "fmt" 
   "os" 
   "errors" 
)

// sorts letters in a word (i.e. "morning" -> "gimnnor")
func sortRunes(str string) string {
   runes := bytes.Runes([]byte(str))
   var temp rune
   for i := 0; i < len(runes); i++ {
      for j := i + 1; j < len(runes); j++ {
         if runes[j] < runes[i] {
            temp = runes[i] {
            runes[i], runes[j] = runes[j], temp
         }
      }
   }
   return string(runes)
}
