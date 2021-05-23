package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	HTTP_BINDING_ADDR = ":8080"
)

func main() {

	r := mux.NewRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    HTTP_BINDING_ADDR,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Binding http server to ", HTTP_BINDING_ADDR)
	log.Fatal(srv.ListenAndServe())
}
