package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

var (
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	responseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_response_status_total",
			Help: "Total number of HTTP responses by status code",
		},
		[]string{"status"},
	)

	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)
)

// Middleware for custom Prometheus metrics
func prometheusMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Path()
		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path)) // Start the timer

		// Call the next handler
		err := next(c)

		// Capture the response status code
		statusCode := c.Response().Status

		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc() // Increment status code counter
		totalRequests.WithLabelValues(path).Inc()                      // Increment total requests counter

		timer.ObserveDuration() // Stop the timer and observe the duration

		return err
	}
}

func main() {
	prometheus.MustRegister(httpDuration, responseStatus, totalRequests)

	e := echo.New()
	//e.Use(echoprometheus.NewMiddleware("echo"))
	e.Use(prometheusMiddleware)

	e.GET("/metrics", echoprometheus.NewHandler())
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := e.Start(":8081"); err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()

	<-done

	fmt.Println("Server shutting down gracefully...")
}
