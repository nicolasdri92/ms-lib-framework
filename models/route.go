package models

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandleFunc  http.HandlerFunc
	Middlewares []mux.MiddlewareFunc
}
