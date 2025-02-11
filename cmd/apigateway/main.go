package main

import (
	"apigateway/router"
	"net/http"
)

func main() {
	r := router.NewRouter()
	r.InitHandlers()
	http.ListenAndServe(":8080", &r)
}
