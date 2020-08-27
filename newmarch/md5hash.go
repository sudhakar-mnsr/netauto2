/* Go has support for MD4, MD5, RIPEMD-160, SHA1, SHA224, SHA256
   SHA384 and SHA512. All follow same pattern as below
   A function new (or similar) returns hash object
   hash has an io.Writer
   write data to be hashed to this writer
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
   
   for n := 0; n < hashSize; n += 4 {
      var val uint32
      val = uint32(hashValue[n]) << 24 +
            uint32(hashValue[n+1]) << 16 +
            uint32(hashValue[n+2]) << 8 +
            uint32(hashValue[n+3])
      fmt.Printf("%x ", val)
   }
   fmt.Println()
}
