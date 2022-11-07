package main

import (
	"net/http"
	"online-election-system/router"
)

func main() {
	r := router.Router()

	http.ListenAndServe(":8080", r)
}
