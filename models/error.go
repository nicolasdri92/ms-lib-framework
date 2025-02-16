package models

type Error struct {
	Date       string `json:"date"`
	StatusCode int    `json:"statusCode"`
	Title      string `json:"title"`
	Message    string `json:"message"`
	Code       string `json:"code"`
}
