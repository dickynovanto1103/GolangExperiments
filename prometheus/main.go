package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func recordMetric() {
	for {
		counterOps.Inc()
		time.Sleep(1*time.Microsecond)
	}
}

var(
	counterOps = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "try_prometheus",
		Name: "counter",
		Help: "just counter metric",
	})
)

func main() {
	//prometheus.CounterOpts{
	//	Name: "counter",
	//	Help: "just metric for gettinc ounter",
	//}
	go recordMetric()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}