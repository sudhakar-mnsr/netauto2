/* GetHeadInfo
*/

package main

import (
   "fmt"
   "io/ioutil"
   "net"
   "os"
)

func main() {
if len(os.Args) != 2 {
   fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
   os.Exit(1)
}
