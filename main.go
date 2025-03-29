package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()
	fmt.Fprint(w, "pong")
}

func main() {
	prometheus.MustRegister(pingCounter)

	http.HandleFunc("/ping", pingHandler)
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", nil)
}
