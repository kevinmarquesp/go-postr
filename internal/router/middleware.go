package router

import "net/http"

// The return will determine if the router has permition start the route
// controller (true) or not (false).
func Middleware(sw http.ResponseWriter, r *http.Request) bool {
	return true
}
