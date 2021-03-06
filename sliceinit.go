package main

import "fmt"

func main() {

var ids []string = []string{"fe225","ac144","3b12c"}
   for _, i := range ids {
      fmt.Println(i)
   }
vector := []float64{12.4, 44, 126, 2, 11.5}
   for _, i := range vector {
      fmt.Println(i)
   }
months := []string {
         "Jan", "Feb", "Mar", "Apr", 
         "May", "Jun", "Jul", "Aug", 
         "Sep", "Oct", "Nov", "Dec", 
         }
   for _, i := range months {
      fmt.Println(i)
   }
tables := []map[string][]int {
         { 
               "age":{53, 13, 5, 55, 45, 62, 34, 7}, 
               "pay":{124, 66, 777, 531, 933, 231}, 
         }, 
    } 
   for _, i := range tables {
      fmt.Println(i)
   }
    graph  := [][][][]int{ 
         {{{44}, {3, 5}}, {{55, 12, 3}, {22, 4}}}, 
         {{{22, 12, 9, 19}, {7, 9}}, {{43, 0, 44, 12}, {7}}},     
    }
   for _, i := range graph {
      fmt.Println(i)
   }
   // Create slice from another slice
   var halfyr = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun",}
   q1 := halfyr[:3]
   q2 := halfyr[3:]
   fmt.Println(q1)
   fmt.Println(q2)
   // Create slice from another array
   var months2 [12]string = [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun","Jul", "Aug", "Sep", "Oct", "Nov", "Dec",}
   q3 := months2[6:9]
   q4 := months2[9:]
   summer := months2[3:6:6]
   fmt.Println(q3)
   fmt.Println(q4)
   fmt.Println(summer)
   months3 := make([]string,6,12)
   fmt.Println(cap(months3)) 
   months4 := make([]string, 3, 3)
   months4 = append(months4,"Jul","Aug","Sep")
   months4 = append(months4, []string{"Jul","Aug","Sep"}...)
   months4 = append(months4,"OCT","NOV","DEC")
   fmt.Println(months4)
   fmt.Println(len(months4), cap(months4))
   v := make([]string, len(months4), cap(months4))
   copy(v, months4)
   fmt.Println(v)
   fmt.Println(len(v), cap(v))
}
