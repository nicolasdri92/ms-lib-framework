package framework

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nicolasdri92/ms-lib-framework/models"
)

func ResponseJSON(writer http.ResponseWriter, results interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(results)
}

func ResponseError(writer http.ResponseWriter, statusCode int, title string, message string) {
	ResponseJSON(writer, models.Error{
		Date:       time.Now().Format(time.RFC3339),
		StatusCode: statusCode,
		Title:      title,
		Message:    message,
	})
}
