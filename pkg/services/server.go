package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexivia/ms-lib-framework/api"
	"github.com/nexivia/ms-lib-framework/internal/config"
	"github.com/nexivia/ms-lib-framework/internal/constants"
	"github.com/nexivia/ms-lib-framework/internal/helpers"
	"github.com/nexivia/ms-lib-framework/internal/models"
)

func ServerSetup() {
	err := config.LoadEnv()

	if err != nil {
		log.Fatalf(constants.ErrorDotenvNotFound)
	}

	helpers.CheckConnected()
}

func ServerListenAndServe(microservice string, routes []models.Route, generalMiddlewares []gin.HandlerFunc) {
	router := api.NewRouter(routes, generalMiddlewares)

	port := config.GetVariable("PORT")

	server := http.ListenAndServe(":"+port, router)

	log.Fatal(server)
}
