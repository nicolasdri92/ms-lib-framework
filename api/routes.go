package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexivia/ms-lib-framework/internal/constants"
	"github.com/nexivia/ms-lib-framework/internal/models"
)

func NewRouter(routes []models.Route, generalMiddlewares []gin.HandlerFunc) *gin.Engine {
	router := gin.Default()

	router.Use(generalMiddlewares...)

	for _, route := range routes {
		routeGroup := router.Group(route.Pattern)

		if route.Middlewares != nil {
			routeGroup.Use(route.Middlewares...)
		}

		switch route.Method {
		case http.MethodGet:
			routeGroup.GET("", route.HandleFunc)
		case http.MethodPost:
			routeGroup.POST("", route.HandleFunc)
		case http.MethodPut:
			routeGroup.PUT("", route.HandleFunc)
		case http.MethodDelete:
			routeGroup.DELETE("", route.HandleFunc)
		case http.MethodPatch:
			routeGroup.PATCH("", route.HandleFunc)
		default:
			log.Printf(constants.HttpMethodNotSupported, route.Method, route.Pattern)
		}

		addOptionsRoute(routeGroup)
	}

	return router
}

func addOptionsRoute(routeGroup *gin.RouterGroup) {
	routeGroup.OPTIONS("", GetOptions)
}

func GetOptions(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	c.Writer.Header().Set("Access-Control-Request-Headers", "Authorization")
	c.JSON(http.StatusOK, gin.H{})
}

func Healthy(c *gin.Context) {
	c.String(http.StatusOK, "Healthy")
}
