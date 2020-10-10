package helpers

import (
   "context"
   "github.com/go-kit/kit/endpoint"
)

// EncryptService is a blueprint for our service
type EncryptService interface {
   Encrypt(context.Context, string, string) (string, error)
   Decrypt(context.Context, string, string) (string, error)
}
