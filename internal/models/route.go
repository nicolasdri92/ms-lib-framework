package models

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string            // Nombre descriptivo de la ruta (ej., "Obtener usuarios").
	Method      string            // Método HTTP permitido para la ruta (ej., "GET", "POST", "PUT", "DELETE").
	Pattern     string            // Patrón de URL asociado a la ruta (ej., "/users", "/products/:id").
	HandleFunc  gin.HandlerFunc   // Función que procesará las solicitudes que coincidan con el patrón.
	Middlewares []gin.HandlerFunc // Lista de middlewares específicos que se ejecutarán antes de HandleFunc.
}
