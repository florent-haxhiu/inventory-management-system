package main

import (
	"net/http"
	"github.com/florent-haxhiu/inventory-management-system/v/routes"
)

func main() {
	r := routes.Router()
	http.ListenAndServe(":3000", r)
}
