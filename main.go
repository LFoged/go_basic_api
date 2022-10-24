package main

import (
	"log"
	"net/http"
	p "plrlsight_api1/product"
)

const apibasePath = "/api"

func main() {
	p.SetupRoutes(apibasePath)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
