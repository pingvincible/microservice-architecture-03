package controllers

import (
	"fmt"
	"net/http"
)

func ReturnHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ReturnHealth")
	fmt.Fprintf(w, "{\"status\":\"OK\"}")
}

func ReturnVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: ReturnVersion")
	fmt.Fprintf(w, "{\"version\":\"0.8\"}")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: home")
	fmt.Fprintf(w, "{\"hello\":\"world\"}")
}
