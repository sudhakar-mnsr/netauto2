package main

import "fmt"

func main() {
histogram := map[string]int{ "Jan":100, "Feb":445, "Mar":514, "Apr":233, "May":321, "Jun":644, "Jul":113, "Aug":734, "Sep":553, "Oct":344, "Nov":831, "Dec":312, }

table := map[string][]int {
   "Men":[]int{32,55,12,42,53},
   "Women":[]int{44, 42, 23, 41, 65, 44},
}
fmt.Println(histogram)
fmt.Println(table)
}
