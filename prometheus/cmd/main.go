package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var pingCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "ping_request_count",
	Help: "Number of requests handled by ping endpoint.",
})

func ping(w http.ResponseWriter, req *http.Request) {
	pingCounter.Inc()

	fmt.Fprintf(w, "pong!")
}

func main() {
	prometheus.MustRegister(pingCounter)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/ping", ping)

	http.ListenAndServe(":8090", nil)
}
