package main

import (
   "log"
   "net/http"
   "os"
   kitlog "github.com/go-kit/kit/log"
   httptransport "github.com/go-kit/kit/transport/http"
   "github.com/narenaryan/encryptService/helpers"
)

func main() {
logger := kitlog.NewLogfmtLogger(os.Stderr)
var svc helpers.EncryptService
svc = helpers.EncryptServiceInstance{}
svc = helpers.LoggingMiddleware{Logger: logger, Next: svc}
encryptHandler := httptransport.NewServer(helpers.MakeEncryptEndpoint(svc),
   helpers.DecodeEncryptRequest, 
   helpers.EncodeResponse)

