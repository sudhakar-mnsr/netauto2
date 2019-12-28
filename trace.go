// The trace program uses defer to add entry/exit diagonstics to function

package main

import (
   "log"
   "time"
)

func bigSlowOperation() {
   defer trace("bigSlowOperation")() // dont forget the extra parenthesis
   time.Sleep(10 * time.Second) // simulate slow operation by sleep
}

func trace(msg string) func() {
   start := time.Now()
   log.Printf("enter %s", msg)
   return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
   bigSlowOperation()
}
