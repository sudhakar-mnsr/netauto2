package main

import (
   "log"
   "net/http"
   "os"
   
   stdprometheus "github.com/prometheus/client_golang/prometheus"
   "github.com/prometheus/client_golang/prometheus/promhttp"
   
   kitlog "github.com/go-kit/kit/log"
   httptransport "github.com/go-kit/kit/transport/http"
   "github.com/sudhakar-mnsr/encryptService/helpers"
   
   kitprometheus "github.com/go-kit/kit/metrics/prometheus"
)

func main() {
logger := kitlog.NewLogfmtLogger(os.Stderr)

fieldKeys := []string{"method", "error"}
requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
   Namespace: "encryption",
   Subsystem: "my_service",
   Name: "request_count",
   Help: "Number of requests received.",
}, fieldKeys)
requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
   Namespace: "encryption",
   Subsystem: "my_service",
   Name: "request_latency_microseconds",
   Help: "Total duration of requests in microseconds.",
}, fieldKeys)

