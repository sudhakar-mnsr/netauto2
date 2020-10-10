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
