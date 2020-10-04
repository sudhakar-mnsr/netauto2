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
