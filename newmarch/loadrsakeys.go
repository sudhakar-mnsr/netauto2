/* LoadRSAKeys
*/

package main

import (
   "crypto/rsa"
   "encoding/gob"
   "fmt"
   "os"
)

func main() {
var key rsa.PrivateKey
loadKey("private.key", &key)


