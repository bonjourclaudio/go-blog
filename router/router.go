package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

var AppRoutes []RoutePrefix

type RoutePrefix struct {
	Prefix string
	Middleware mux.MiddlewareFunc
	SubRoutes []Route
}

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}
