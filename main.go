package main

import (
	"net/http"

	"github.com/katamayo/NewFramework/app/models"
	"github.com/katamayo/NewFramework/route"
)

func main() {
	// db := sql.Open("postgres", "...")
	print("Welcome, Katamayo New Framework with Golang\n")
	router := route.RouteApp(&models.AppContext{})
	http.ListenAndServe(":8080", router)
}
