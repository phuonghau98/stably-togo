package rest

import (
	"encoding/json"
	"math/big"
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
	ErrPrimeInputNumberOutOfRange = PrimeRequestValidationError("Input number should be larger than 1")
	ErrPrimeInvalidInputFormat    = PrimeRequestValidationError("Invalid request")
)

// Validations

type findNearestPrimeBodyRequest struct {
	Num string `json:"num",omitempty`
}

func (f findNearestPrimeBodyRequest) validate() error {
	n := new(big.Int)
	n, ok := n.SetString(f.Num, 10)
	if !ok {
		return ErrPrimeInvalidInputFormat
	}
	if n.Cmp(big.NewInt(1)) < 0 {
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
	router.HandleFunc("/api/v2/prime/findnearest", ph.FindLowerNearestPrimeV2).Methods(http.MethodPost)
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
		writeErrorJSONResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert string to bigint
	n := new(big.Int)
	n, ok := n.SetString(reqBody.Num, 10)
	if !ok {
		writeErrorJSONResponse(w, ErrPrimeInvalidInputFormat.Error(), http.StatusBadRequest)
		return
	}

	// Process the request
	foundPrimeNumber, err := prime.FindLowerNearestPrimeNumber(n)
	if err != nil {
		writeErrorJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeSuccessJSONResponse(w, &findNearestPrimeBodyResponse{Num: foundPrimeNumber}, 200)
	return
}

func (handler *PrimeHandler) FindLowerNearestPrimeV2(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var reqBody findNearestPrimeBodyRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		writeErrorJSONResponse(w, ErrPrimeInvalidInputFormat.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// validate inputs
	if err := reqBody.validate(); err != nil {
		writeErrorJSONResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert string to bigint
	n := new(big.Int)
	n, ok := n.SetString(reqBody.Num, 10)
	if !ok {
		writeErrorJSONResponse(w, ErrPrimeInvalidInputFormat.Error(), http.StatusBadRequest)
		return
	}

	// Process the request
	foundPrimeNumber, err := prime.FindLowerNearestPrimeNumberV2(n)
	if err != nil {
		writeErrorJSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeSuccessJSONResponse(w, &findNearestPrimeBodyResponse{Num: foundPrimeNumber}, 200)
	return
}
