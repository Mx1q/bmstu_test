package main

import (
	"encoding/json"
	"github.com/766b/chi-prometheus"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const (
	srvAddr     = ":8082"
	metricsAddr = ":8081"
)

//var (
//	httpDuration = apps.NewHistogramVec(
//		apps.HistogramOpts{
//			Name:    "http_request_duration_seconds",
//			Help:    "Duration of HTTP requests in seconds",
//			Buckets: apps.DefBuckets,
//		},
//		[]string{"path"},
//	)
//
//	responseStatus = apps.NewCounterVec(
//		apps.CounterOpts{
//			Name: "http_response_status_total",
//			Help: "Total number of HTTP responses by status code",
//		},
//		[]string{"status"},
//	)
//
//	totalRequests = apps.NewCounterVec(
//		apps.CounterOpts{
//			Name: "http_requests_total",
//			Help: "Total number of HTTP requests",
//		},
//		[]string{"path"},
//	)
//)
//
//// Middleware for custom Prometheus metrics
//func prometheusMiddleware() middleware.Middleware {
//	return M(c *gin.Context) {
//		start := time.Now()
//
//		// Process request
//		c.Next() // Call the next handler
//
//		// Measure duration
//		duration := time.Since(start).Seconds()
//		path := c.Request.URL.Path
//
//		// Capture metrics
//		httpDuration.WithLabelValues(path).Observe(duration)
//		statusCode := strconv.Itoa(c.Writer.Status())
//		responseStatus.WithLabelValues(statusCode).Inc()
//		totalRequests.WithLabelValues(path).Inc()
//	}
//}

type ResponseStruct struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func main() {
	r := chi.NewRouter()

	m := chiprometheus.NewMiddleware("chi")
	r.Use(m)

	r.Handle("/metrics", promhttp.Handler())

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ResponseStruct{Status: http.StatusText(http.StatusOK), Data: "pong"})
	})

	http.ListenAndServe(":8081", r)
}
