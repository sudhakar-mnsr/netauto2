/* FTP server
*/

package main

import (
   "fmt"
   "net"
   "os"
)

const (
   DIR = "DIR"
   CD = "CD"
   PWD = "PWD"
)

func main() {
service := "0.0.0.0:1202"
tcpAddr, err := net.ResolveTCPAddr("tcp", service)
checkError(err)


