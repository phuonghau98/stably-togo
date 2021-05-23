package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type PrimeHandler struct{}

func NewPrimeHandler() *PrimeHandler {
	return &PrimeHandler{}
}

func (ph *PrimeHandler) Register(router *mux.Router) {
	router.HandleFunc("/api/v1/prime/findnearest", ph.FindNearestPrime).Methods(http.MethodPost)
}

func (handler PrimeHandler) FindNearestPrime(w http.ResponseWriter, r *http.Request) {
	writeErrorJSONResponse(w, "Not implemented yet", 500)
	return
}
