package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pingvincible/crud-service/pkg/routes"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"net/http"
	"strconv"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

var requestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "app_request_count",
		Help: "Application Request Count",
	},
	[]string{"method", "endpoint", "http_status"},
)

var requestLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "app_request_latency_seconds",
	Help: "Application Request Latency",
}, []string{"method", "endpoint"})

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		method := r.Method
		endpoint, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(requestLatency.WithLabelValues(method, endpoint))
		rw := NewResponseWriter(w)
		next.ServeHTTP(rw, r)

		statusCode := rw.statusCode

		requestCount.WithLabelValues(method, endpoint, strconv.Itoa(statusCode)).Inc()

		timer.ObserveDuration()
	})
}

func init() {
	prometheus.Register(requestCount)
	prometheus.Register(requestLatency)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(prometheusMiddleware)
	routes.RegisterUserRoutes(myRouter)
	routes.RegisterAppRoutes(myRouter)

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	fmt.Println("CRUD service")
	handleRequests()
}
