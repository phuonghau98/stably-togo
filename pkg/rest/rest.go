package rest

import (
	"encoding/json"
	"net/http"
)

type errorJSONResponse struct {
	Error string `json:"error"`
}

type successJSONResponse struct {
	Data interface{} `json:"data"`
}

func writeErrorJSONResponse(w http.ResponseWriter, msg string, status int) {
	writeJSON(w, errorJSONResponse{Error: msg}, status)
}

func writeSuccessJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	writeJSON(w, successJSONResponse{Data: data}, status)
}

func writeJSON(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)
	w.Write(content)
}
