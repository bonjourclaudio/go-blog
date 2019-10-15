package customRouter

import (
	"net/http"
)

var Routes []RoutePrefix

type RoutePrefix struct {
	IsSecure bool
	Prefix string
	SubRoutes []Route
}

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}