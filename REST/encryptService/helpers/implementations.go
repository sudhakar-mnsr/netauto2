package helpers
import (
    "context"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "errors"
)

// EncryptServiceInstance is the implementation of interface for microservice
type EncryptServiceInstance struct{}
// Implements AES encryption algorithm(Rijndael Algorithm)
// Initialization vector for the AES algorithm 
var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
// Encrypt encrypts the string with given key
func (EncryptServiceInstance) Encrypt(_ context.Context, key string, text string) (string, error) {
   block, err := aes.NewCipher([]byte(key))
   if err != nil {
      panic(err)
   }
   plaintext := []byte(text)
   cfb := cipher.NewCFBEncrypter(block, initVector)
   cipertext := make([]byte, len(plaintext))
   cfb.XORKeyStream(ciphertext, plaintext)
   return base64.StdEncoding.EncodeToString(cipertext), nil
}
