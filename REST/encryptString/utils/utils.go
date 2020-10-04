package utils
import (
   "crypto/aes"
   "crypto/cipher"
   "encoding/base64"
)

// Implements AES encryption algorithm(Rijndael Algorithm)
/* Initialization vector for the AES algorithm
More details visit this link https://en.wikipedia.org/wiki/Advanced_Encryption_Standard */
var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// EncryptString encrypts the string with given key
func EncryptString(key, text string) string {
block, err := aes.NewCipher([]byte(key))
