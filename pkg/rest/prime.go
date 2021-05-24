package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/phuonghau98/stably-togo/pkg/service/prime"
)

// Custom errors to embrace reusability
type PrimeRequestValidationError string

func (e PrimeRequestValidationError) Error() string {
	return string(e)
}

var (
	ErrPrimeInputNumberOutOfRange = PrimeRequestValidationError("Input number should be larger than 1")
	ErrPrimeInvalidInputFormat    = PrimeRequestValidationError("Invalid request")
)

// Validations

type findNearestPrimeBodyRequest struct {
	Num int `json:"num",omitempty`
}

func (f findNearestPrimeBodyRequest) validate() error {
	if f.Num <= 1 {
		return ErrPrimeInputNumberOutOfRange
	}
	return nil
}

type findNearestPrimeBodyResponse struct {
	Num string `json:"num"`
}

// Hanlders

type PrimeHandler struct{}

func NewPrimeHandler() *PrimeHandler {
	return &PrimeHandler{}
}

func (ph *PrimeHandler) Register(router *mux.Router) {
	router.HandleFunc("/api/v1/prime/findnearest", ph.FindLowerNearestPrimeV1).Methods(http.MethodPost)
}

func (handler *PrimeHandler) FindLowerNearestPrimeV1(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var reqBody findNearestPrimeBodyRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		writeErrorJSONResponse(w, ErrPrimeInvalidInputFormat.Error(), http.StatusBadRequest)
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

	writeSuccessJSONResponse(w, &findNearestPrimeBodyResponse{Num: strconv.Itoa(foundPrimeNumber)}, 200)
	return
}
