package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/pingvincible/crud-service/pkg/routes"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	routes.RegisterUserRoutes(myRouter)
	routes.RegisterAppRoutes(myRouter)

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {
	fmt.Println("CRUD service")
	handleRequests()
}
