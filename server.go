package framework

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerSetup() {
	err := LoadEnv()

	if err != nil {
		log.Fatalf(ErrorDotenvNotFound)
	}

	CheckConnected()
}

func ServerListenAndServe(microservice string, routes []Route, generalMiddlewares []gin.HandlerFunc) {
	router := NewRouter(routes, generalMiddlewares)

	port := GetVariable("PORT")

	server := http.ListenAndServe(":"+port, router)

	log.Fatal(server)
}
