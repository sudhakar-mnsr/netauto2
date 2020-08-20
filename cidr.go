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
