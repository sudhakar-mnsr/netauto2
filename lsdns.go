package main

import (
   "context"
   "flag"
   "fmt"
   "net"
   "os"
)

var (
   ip string
   host string
   ns bool
   mx bool
   txt bool
   cname bool
)
