package main

import (
	"net/http"
	"training/routes"
)

func main() {
	router := routes.InitialRoute()
	http.ListenAndServe(":8080", router)
}
