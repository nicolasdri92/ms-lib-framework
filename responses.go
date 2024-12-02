package framework

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func ResponseJSON(w http.ResponseWriter, status int, results interface{}) {
	response, err := json.Marshal(results)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func ResponseError(w http.ResponseWriter, code string, status int, types string, trace string, title string, message string) {
	err := Error{
		Codigo:    strconv.Itoa(status),
		Error:     message,
		Timestamp: time.Now().Format("2006-01-02T00:00:00"),
		Code:      code,
		Status:    status,
		Type:      types,
		Trace:     trace,
		Title:     title,
		Message:   message,
	}

	ResponseJSON(w, status, err)
}
