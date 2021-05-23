package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/phuonghau98/stably-togo/pkg/rest"
	"github.com/rs/cors"
)

var (
	HTTP_BINDING_ADDR = ":8080"
	ENV               = "development"
)

func main() {

	r := mux.NewRouter()
	primeHandler := rest.NewPrimeHandler()
	primeHandler.Register(r)

	// Cors
	corsOptions := cors.Options{}
	if ENV != "production" {
		corsOptions.AllowedOrigins = []string{"http://localhost:3000"}
	}

	corsInstance := cors.New(corsOptions)

	srv := &http.Server{
		Handler: corsInstance.Handler(r),
		Addr:    HTTP_BINDING_ADDR,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Binding http server to ", HTTP_BINDING_ADDR)
	log.Fatal(srv.ListenAndServe())
}
