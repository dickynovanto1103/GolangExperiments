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
		time.Sleep(1*time.Second)
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
		Buckets:     prometheus.DefBuckets,
	}, []string{"code"})

	httpCounterOps = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "test",
		Name: "http",
	})
)

func StartPrometheusHandler() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
	httpCounterOps.Inc()
}

func main() {
	//prometheus.CounterOpts{
	//	Name: "counter",
	//	Help: "just metric for gettinc ounter",
	//}
	StartPrometheusHandler()
	http.HandleFunc("/test", handler)
	http.ListenAndServe(":8080", nil)
}