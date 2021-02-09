package main

import (
	"net/http"

	"github.com/raphaelkieling/orders/application"
)

func main() {
	router := application.CreateRouter()
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
