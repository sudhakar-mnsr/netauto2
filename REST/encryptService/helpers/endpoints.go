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

// MakeEncryptEndpoint forms endpoint for request/response of encrypt func
func MakeEncryptEndpoint(svc EncryptService) endpoint.Endpoint {
return func(ctx context.Context, request interface{}) (interface{}, error) {
   req := request.(EncryptRequest)
   message, err := svc.Encrypt(ctx, req.Key, req.Text)
   if err != nil {
      return EncryptResponse{message, err.Error()}, nil
   }
   return EncryptResponse{message, ""}, nil
}

