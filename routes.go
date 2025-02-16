package framework

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicolasdri92/ms-lib-framework/models"
)

func NewRouter(routes []models.Route, middlewares []mux.MiddlewareFunc) *mux.Router {
	router := mux.
		NewRouter().
		StrictSlash(true)

	router.Use(middlewares...)

	for _, route := range routes {
		setupRoute(router, route)
	}

	return router
}

func setupRoute(router *mux.Router, route models.Route) {
	subRouter := router.
		PathPrefix("/api/v1").
		Methods(route.Method).
		Path(route.Pattern).
		Subrouter()

	subRouter.
		Name(route.Name).
		Handler(route.HandleFunc)

	if route.Middlewares != nil {
		subRouter.Use(route.Middlewares...)
	}

	addOptionsRoute(subRouter, route)
}

func addOptionsRoute(router *mux.Router, route models.Route) {
	router.Methods(http.MethodOptions).
		Path(route.Pattern).
		Name(route.Name + http.MethodOptions).
		HandlerFunc(getOptions)
}

func getOptions(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Headers", "*")
	writer.Header().Set("Access-Control-Request-Headers", "Authorization")

	ResponseJSON(writer, nil)
}

// func Healthy(writer http.ResponseWriter, _ *http.Request) {
// 	writer.Write([]byte("Healthy"))
// }
