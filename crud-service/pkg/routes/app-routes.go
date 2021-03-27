package routes

import (
	"github.com/gorilla/mux"
	"github.com/pingvincible/crud-service/pkg/controllers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var RegisterAppRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/health", controllers.ReturnHealth).Methods("GET")
	router.HandleFunc("/version", controllers.ReturnVersion).Methods("GET")
	router.Path("/metrics").Handler(promhttp.Handler())
}
