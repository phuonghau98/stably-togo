package rest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

type findLowerNearestResponse struct {
	Error string                       `json:"error"`
	Data  findNearestPrimeBodyResponse `json:"data"`
}

func TestFindLowerNearestPrimeHandler(t *testing.T) {
	// Mock server
	router := mux.NewRouter()
	primeHandler := NewPrimeHandler()
	primeHandler.Register(router)
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Utils
	newreq := func(method, url string, bodyPayload interface{}) *http.Request {
		bodyJSON, err := json.Marshal(bodyPayload)
		if err != nil {
			log.Fatalf("Cannot prepare post body: %v", err)
		}
		r, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	testCases := []struct {
		name        string
		request     *http.Request
		err         error
		responseNum int
	}{
		{
			name:        "Should return body contains `num: 11` as response",
			request:     newreq(http.MethodPost, ts.URL+"/api/v1/prime/findnearest", findNearestPrimeBodyRequest{Num: "12"}),
			err:         nil,
			responseNum: 11,
		},
		{
			name:        "Should return body contains `num: -1` as response",
			request:     newreq(http.MethodPost, ts.URL+"/api/v1/prime/findnearest", findNearestPrimeBodyRequest{Num: "2"}),
			err:         nil,
			responseNum: -1,
		},
		{
			name:        "Should return body contains `num: 8971` as response ",
			request:     newreq(http.MethodPost, ts.URL+"/api/v1/prime/findnearest", findNearestPrimeBodyRequest{Num: "8999"}),
			err:         nil,
			responseNum: 8971,
		},
		{
			name: "Should return `invalid request` error as response if num provided with string format",
			request: newreq(http.MethodPost, ts.URL+"/api/v1/prime/findnearest", struct {
				Num string `json:"num"`
			}{Num: "23"}),
			err: ErrPrimeInvalidInputFormat,
		},
		{
			name: "Should return `invalid request` error as response if num provided with object format",
			request: newreq(http.MethodPost, ts.URL+"/api/v1/prime/findnearest", struct {
				Num string `json:"num"`
			}{Num: "{}"}),
			err: ErrPrimeInvalidInputFormat,
		},
		{
			name:    "Should return `ErrPrimeInputNumberOutOfRange` error as response if num provided less than 2",
			request: newreq(http.MethodPost, ts.URL+"/api/v1/prime/findnearest", findNearestPrimeBodyRequest{Num: "-1"}),
			err:     ErrPrimeInputNumberOutOfRange,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Parse body
			resp, err := http.DefaultClient.Do(testCase.request)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			var response findLowerNearestResponse
			err = json.Unmarshal(body, &response)
			if err != nil {
				t.Fatalf("Cannot parse response body: %v", err)
			}

			// Assert error case
			if testCase.err != nil {
				wantedError := testCase.err.Error()
				gotError := response.Error
				if wantedError != gotError {
					log.Fatalf("Should receive error: %v, but got: %v", wantedError, gotError)
				}
			} else {
				// Assert valid case
				wantedNum := testCase.responseNum
				gotNum := response.Data.Num

				if strconv.Itoa(wantedNum) != gotNum {
					log.Fatalf("Should receive num: %v as response, but got: %v", wantedNum, gotNum)
				}
			}
		})
	}
}

func TestFindLowerNearestPrimeHandlerV2(t *testing.T) {
	// Mock server
	router := mux.NewRouter()
	primeHandler := NewPrimeHandler()
	primeHandler.Register(router)
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Utils
	newreq := func(method, url string, bodyPayload interface{}) *http.Request {
		bodyJSON, err := json.Marshal(bodyPayload)
		if err != nil {
			log.Fatalf("Cannot prepare post body: %v", err)
		}
		r, err := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	testCases := []struct {
		name        string
		request     *http.Request
		err         error
		responseNum int
	}{
		{
			name:        "Should return body contains `num: 11` as response",
			request:     newreq(http.MethodPost, ts.URL+"/api/v2/prime/findnearest", findNearestPrimeBodyRequest{Num: "12"}),
			err:         nil,
			responseNum: 11,
		},
		{
			name:        "Should return body contains `num: -1` as response",
			request:     newreq(http.MethodPost, ts.URL+"/api/v2/prime/findnearest", findNearestPrimeBodyRequest{Num: "2"}),
			err:         nil,
			responseNum: -1,
		},
		{
			name:        "Should return body contains `num: 1000000395723132233` as response ",
			request:     newreq(http.MethodPost, ts.URL+"/api/v2/prime/findnearest", findNearestPrimeBodyRequest{Num: "1000000395723132283"}),
			err:         nil,
			responseNum: 1000000395723132233,
		},
		{
			name: "Should return `invalid request` error as response if num provided with string format",
			request: newreq(http.MethodPost, ts.URL+"/api/v2/prime/findnearest", struct {
				Num string `json:"num"`
			}{Num: "23"}),
			err: ErrPrimeInvalidInputFormat,
		},
		{
			name: "Should return `invalid request` error as response if num provided with object format",
			request: newreq(http.MethodPost, ts.URL+"/api/v2/prime/findnearest", struct {
				Num string `json:"num"`
			}{Num: "{}"}),
			err: ErrPrimeInvalidInputFormat,
		},
		{
			name:    "Should return `ErrPrimeInputNumberOutOfRange` error as response if num provided less than 2",
			request: newreq(http.MethodPost, ts.URL+"/api/v2/prime/findnearest", findNearestPrimeBodyRequest{Num: "-1"}),
			err:     ErrPrimeInputNumberOutOfRange,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Parse body
			resp, err := http.DefaultClient.Do(testCase.request)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			var response findLowerNearestResponse
			err = json.Unmarshal(body, &response)
			if err != nil {
				t.Fatalf("Cannot parse response body: %v", err)
			}

			// Assert error case
			if testCase.err != nil {
				wantedError := testCase.err.Error()
				gotError := response.Error
				if wantedError != gotError {
					log.Fatalf("Should receive error: %v, but got: %v", wantedError, gotError)
				}
			} else {
				// Assert valid case
				wantedNum := testCase.responseNum
				gotNum := response.Data.Num

				if strconv.Itoa(wantedNum) != gotNum {
					log.Fatalf("Should receive num: %v as response, but got: %v", wantedNum, gotNum)
				}
			}
		})
	}
}
