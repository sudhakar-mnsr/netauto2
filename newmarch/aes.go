/* AES
*/

package main

import (
   "bytes"
   "crypto/aes"
   "fmt"
)

func main() {
key := []byte("my key, len 16 b")
cipher, err := aes.NewCipher(key)
if err != nil {
   fmt.Println(err.Error())
}
src := []byte("hello 16b block")


