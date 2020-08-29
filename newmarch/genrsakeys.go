/* GenRSAKeys
*/

package main

import (
   "crypto/rand"
   "crypto/rsa"
   "crypto/x509"
   "encoding/gob"
   "encoding/pem"
   "fmt"
   "os"
)

func main() {
reader := rand.Reader
bitSize := 512
key, err := rsa.GenerateKey(reader, bitSize)
checkError(err)
