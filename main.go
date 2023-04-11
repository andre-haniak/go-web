package main

import (
	"net/http"

	"github.com/andre-haniak/go-web/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
