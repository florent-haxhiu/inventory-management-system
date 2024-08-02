package main

import (
	"net/http"
	"github.com/florent-haxhiu/inventory-management-system/v/routes"
)

func main() {
	r := routes.RouteMe()
	http.ListenAndServe(":3000", r)
}
