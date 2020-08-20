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
   
   // Given IPv4 block 192.168.100.14/24.
   // The followings uses IPNet to get:
   // - The routing address for the subnet - 192.168.100.0
   // - one-bits of the network mask 24 of 32 i.e 255.255.255.0
   // - Total hosts in the network 2^8
   // Wildcard the inverse of the subnet i.e 0.0.0.255
   // Max address of the subnet 192.168.100.255
   ones, totalBits := ipnet.Mask.Size()
   size := totalBits - ones
   totalHosts := math.Pow(2, float64(size))
   wildcardIP := wildcard(net.IP(ipnet.Mask))
   last := lastIP(ip, net.IPMask(wildcardIP))
   
   fmt.Println()
   fmt.Printf("CIDR: %s\n", cidr)
   fmt.Println("==========================================================")
   fmt.Printf("CIDR Block:                     %s\n", cidr)
   fmt.Printf("Network:                        %s\n", ipnet.IP)
   fmt.Printf("IP Range:                       %s - %s\n", ipnet.IP, last)
   fmt.Printf("Total Hosts:                    %0.0f\n", totalhosts)
   fmt.Printf("Netmask:                        %s\n", net.IP(ipnet.Mask))
   fmt.Printf("Wildcard Mask:                  %s\n", wildcardIP)
   fmt.Println()
}

// wildcard returns the opposite of the netmask for the network
func wildcard(mask net.IP) net.IP {
   var ipVal net.IP
   for _, octet := range mask {
      ipVal = append(ipVal, ^octet)
   }
   return ipVal
}

// lastIP calculates the highest addressable IP for the given subnet.
// It loops through each octet of the subnets IP address and applies 
// bitwise OR operation to each corresponding octet from mask
func lastIP(ip net.IP, mask net.IPMask) net.IP {
   ipIN := ip.To4()
   if ipIn == nil {
      ipIn = ip.TO16()
      if ipIn == nil {
         return nil
      }
   }
