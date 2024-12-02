package framework

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ServerSetup() {
	err := LoadEnv()

	if err != nil {
		log.Fatalf(ErrorDotenvNotFound)
	}

	CheckConnected()
}

func ServerListenAndServe(routes []Route, generalMiddlewares []mux.MiddlewareFunc) {
	router := NewRouter(routes, generalMiddlewares)

	port := GetVariable("PORT")

	server := http.ListenAndServe(port, router)

	log.Fatal(server)
}
