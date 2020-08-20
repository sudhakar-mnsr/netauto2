package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"os"
)

var (
	cidr string
)

func init() {
	flag.StringVar(&cidr, "c", "", "the CIDR address")
}

// This program implements a CIDR subnet calculator.
// It takes a CIDR address prefix an calculates ip-ranges,
// total hosts, wildcard mask, etc.

func main() {
flag.Parse()
if cidr == "" {
   fmt.Println("CIDR address missing")
   os.Exit(1)
}

ip, ipnet, err := net.ParseCIDR(cidr)
if err != nil {
   fmt.Println("failed parsing CIDR address: ", err)
   os.Exit(1)
}

