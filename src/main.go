package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

var requestDuration = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "request_duration_seconds",
		Help:    "Determines how long a request takes in second",
		Buckets: []float64{1, 2, 4},
	},
	[]string{"path", "method"},
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	// start timer
	start := time.Now()

	time.Sleep(time.Duration(rand.IntN(5)) * time.Second)

	pingCounter.Inc()

	// calculate duration
	duration := time.Since(start).Seconds()

	requestDuration.WithLabelValues("GET", "/ping").Observe(duration)

	fmt.Fprint(w, "pong: ", duration)
}

func main() {
	prometheus.MustRegister(pingCounter, requestDuration)

	http.HandleFunc("/ping", pingHandler)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
