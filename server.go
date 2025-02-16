package framework

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicolasdri92/ms-lib-framework/models"
)

func ServerSetup() {
	err := LoadEnv()

	if err != nil {
		log.Fatalf(ErrorDotenvNotFound)
	}

	CheckConnected()
}

func ListenAndServe(routes []models.Route, middlewares []mux.MiddlewareFunc) {
	addr := GetVariable("PORT")
	handler := NewRouter(routes, middlewares)
	err := http.ListenAndServe(":"+addr, handler)

	log.Fatal(err)
}
