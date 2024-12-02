package framework

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(routes []Route, generalMiddlewares []mux.MiddlewareFunc) *mux.Router {
	router := mux.
		NewRouter().
		StrictSlash(true)

	router.Use(generalMiddlewares...)

	for _, route := range routes {
		setupRoute(router, route)
	}

	return router
}

func setupRoute(router *mux.Router, route Route) {
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

func addOptionsRoute(router *mux.Router, route Route) {
	router.Methods("OPTIONS").
		Path(route.Pattern).
		Name(route.Name + "Options").
		HandlerFunc(GetOptions)
}

func GetOptions(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Headers", "*")
	writer.Header().Set("Access-Control-Request-Headers", "Authorization")

	ResponseJSON(writer, http.StatusOK, nil)
}

func Healthy(writer http.ResponseWriter, _ *http.Request) {
	writer.Write([]byte("Healthy"))
}
