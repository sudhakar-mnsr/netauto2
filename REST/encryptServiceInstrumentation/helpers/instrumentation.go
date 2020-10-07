package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

// InstrumentingMiddleware is a struct representing middleware
type InstrumentingMiddleware struct {
   RequestCount metrics.Counter
   RequestLatency metrics.Histogram
   Next EncryptService
}
