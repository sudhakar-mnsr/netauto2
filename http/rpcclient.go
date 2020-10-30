  
package main

import (
	"log"
	"net/rpc"
)

type Args struct {
}

func main() {
	var reply int64
	args := Args{}
	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
