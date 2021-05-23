package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phuonghau98/stably-togo/pkg/service/prime"
)

// Custom errors to embrace reusability
type PrimeRequestValidationError string

func (e PrimeRequestValidationError) Error() string {
	return string(e)
}

var (
	ErrInputNumberOutOfRange = PrimeRequestValidationError("Input number should be larger than 1")
	ErrInvalidInputFormat    = PrimeRequestValidationError("Invalid request")
)

type PrimeHandler struct{}

func NewPrimeHandler() *PrimeHandler {
	return &PrimeHandler{}
}

func (ph *PrimeHandler) Register(router *mux.Router) {
	router.HandleFunc("/api/v1/prime/findnearest", ph.FindLowerNearestPrimeV1).Methods(http.MethodPost)
}

type findNearestPrimeBodyRequest struct {
	Num int `json:"num",omitempty`
}

func (f findNearestPrimeBodyRequest) validate() error {
	if f.Num <= 1 {
		return ErrInputNumberOutOfRange
	}
	return nil
}

type findNearestPrimeBodyResponse struct {
	Num int `json:"num"`
}

func (handler PrimeHandler) FindLowerNearestPrimeV1(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var reqBody findNearestPrimeBodyRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		writeErrorJSONResponse(w, ErrInvalidInputFormat.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// validate inputs
	if err := reqBody.validate(); err != nil {
		writeErrorJSONResponse(w, err.Error(), 400)
		return
	}

	// Process the request
	foundPrimeNumber, err := prime.FindLowerNearestPrimeNumber(reqBody.Num)
	if err != nil {
		writeErrorJSONResponse(w, fmt.Sprintf("%v", foundPrimeNumber), 500)
		return
	}

	writeSuccessJSONResponse(w, &findNearestPrimeBodyResponse{Num: foundPrimeNumber}, 200)
	return
}
