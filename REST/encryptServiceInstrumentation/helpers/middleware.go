package helpers

import (
   "context"
   "time"
   log "github.com/go-kit/kit/log"
)

// LoggingMiddleware wraps the logs for incoming requests
type LoggingMiddleware struct {
   Logger log.Logger
   Next EncryptService
}
