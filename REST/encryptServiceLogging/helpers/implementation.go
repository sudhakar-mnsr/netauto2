package helpers

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// EncryptServiceInstance is the implementation of interface for micro service
type EncryptServiceInstance struct{}

// Implements AES encryption algorithm(Rijndael Algorithm)
/* Initialization vector for the AES algorithm
   More details visit this link https://en.wikipedia.org/wiki/Advanced_Encryption_Standard */
var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// Encrypt encrypts the string with given key
func (EncryptServiceInstance) Encrypt(_ context.Context, key string, text string) (string, error) {
block, err := aes.NewCipher([]byte(key))
if err != nil {
   panic(err)
}
