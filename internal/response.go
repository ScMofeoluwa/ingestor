package ingestor

import (
	"encoding/json"
	"net/http"
)

type SuccessPayload struct {
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"code"`
}

type ErrorPayload struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
}

func SuccessResponse(w http.ResponseWriter, payload *SuccessPayload) {
	writeJSONResponse(w, payload.StatusCode, payload)
}

func ErrorResponse(w http.ResponseWriter, payload *ErrorPayload) {
	writeJSONResponse(w, payload.StatusCode, payload)
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
