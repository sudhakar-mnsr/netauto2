package main

import (
"log"
"net/http"
httptransport "github.com/go-kit/kit/transport/http"
"github.com/narenaryan/encryptService/helpers"
)

func main() {
svc := helpers.EncryptServiceInstance{}
encryptHandler := httptransport.NewServer(helpers.MakeEncryptEndpoint(svc)
helpers.DecodeEncryptRequest,\
helpers.EncodeResponse)
