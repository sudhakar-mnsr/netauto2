package main

import "fmt"

func main() {
hist := make(map[string]int)
hist["Jun"] = 644
hist["Oct"] = 648
hist["Jul"] = 645
remove(hist, "Oct")
fmt.Println(len(hist))
}

func remove(store


