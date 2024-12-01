package framework

import (
	"fmt"
	"net/http"
	"os"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				ResponseError(w, "LIB_FW_PANIC_MIDDLEWARE", http.StatusInternalServerError, "HIGH", fmt.Sprint(r), "Error", fmt.Sprint(r))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN") != "" {
			w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN"))
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		next.ServeHTTP(w, r)
	})
}
