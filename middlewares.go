package framework

import (
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				ResponseError(writer, http.StatusInternalServerError, "", "")
			}
		}()

		next.ServeHTTP(writer, request)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		writer.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(writer, request)
	})
}
