/* ThreadedIPEchoServer
*/

package main

import (
"fmt"
"net"
"os"
)

func main() {
service := ":1200"
listener, err := net.Listen("tcp", service)
checkError(err)


