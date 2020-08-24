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
                      Help: "HTTP request counts",
                   },
                   []string{"path", "method"},
   )
