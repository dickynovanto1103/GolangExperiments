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

func recordHistogramMetric() {
	for {
		histogramOps.With(prometheus.Labels{"code": "200"}).Observe(0.2)
		time.Sleep(1*time.Second)
	}
}

var(
	counterOps = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "try_prometheus",
		Name: "counter",
		Help: "just counter metric",
	})

	histogramOps = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "try_prometheus",
		Name:        "histogram",
		Help:        "just histogram",
		Buckets:     []float64{0.1, 0.2, 1},
	}, []string{"code"})
)

func main() {
	//prometheus.CounterOpts{
	//	Name: "counter",
	//	Help: "just metric for gettinc ounter",
	//}
	go recordMetric()
	go recordHistogramMetric()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}