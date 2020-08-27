/* A hash has an io.Writer
   you write data to be hashed to this writer
   query the number of bytes in the hash value by Size
   query hash value by Sum
   MD5:
   hash value is 16-byte array
   printed as four hexadecimal numbers (each of 4 bytes)
*/

package main

import (
   "crypto/md5"
   "fmt"
)

func main() {
hash := md5.New()
bytes := []byte("hello\n")
hash.Write(bytes)
hashValue := hash.Sum(nil)
hashSize := hash.Size()

