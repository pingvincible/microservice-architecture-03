package routes

import (
	"github.com/gorilla/mux"
	"github.com/pingvincible/crud-service/pkg/controllers"
)

var RegisterAppRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	router.HandleFunc("/health", controllers.ReturnHealth).Methods("GET")
	router.HandleFunc("/version", controllers.ReturnVersion).Methods("GET")
}
