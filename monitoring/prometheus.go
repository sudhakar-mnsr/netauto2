package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
   httpRequests = prometheus.NewCounterVec(
                  prometheus.CounterOpts{
                     Name: "http_request_count_total",
                     Help: "Http request counts",
                  },
                  []string{"path", "method"},
   }
   httpDurations = prometheus.NewSummaryVec(
                   prometheus.SummaryOpts{
                      Name: "http_request_durations_microseconds",
                      Help: "HTTP request latency distributions.",
                      Objectives: map[float64]float64{},
                   },
                   []string{"path", "method"},
   )
)

func init() {
   prometheus.MustRegister(httpRequests)
   prometheus.MustRegister(httpDurations)
}

func MetricHandler(handlerFunc func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
   return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
      start := time.Now()
      handlerFunc(writer, request)

      elapsed := time.Since(start)
      msElapsed := elapsed / time.Microsecond

      httpRequests.WithLabelValues(request.URL.Path, request.Method).Inc()
      httpDurations.WithLabelValues(request.URL.Path, request.Method).Observe(float64(msElapsed))
   })
}
