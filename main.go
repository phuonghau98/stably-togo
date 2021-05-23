package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/phuonghau98/stably-togo/pkg/rest"
	"github.com/phuonghau98/stably-togo/pkg/util"
	"github.com/rs/cors"
)

var (
	HTTP_BINDING_ADDR                  = ":8080"
	ENV                                = "development"
	DOT_ENV_FILENAME                   = ".env"
	UI_BUILD_FOLDER_KEY                = "UI_BUILD_FOLDER"
	UI_BUILD_INDEX_FILENAME_KEY        = "UI_BUILD_INDEX_FILENAME"
	UI_CLIENT_URL_KEY                  = "UI_CLIENT_URL"
	ENV_KEY                            = "ENV"
	ENV_PRODUCTION                     = "production"
	ENV_DEVELOPMENT                    = "development"
	DEFAULT_UI_BUILD_FOLDER_PATH       = "ui/build"
	DEFAULT_UI_BUILD_INDEX_FILENAME    = "index.html"
	DEFAULT_CLIENT_URL                 = "http://localhost:3000"
	DEFAULT_ENV                        = "development"
	HTTP_SERVER_WRITE_TIMEOUT_DURATION = 30 * time.Second
	HTTP_SERVER_READ_TIMEOUT_DURATION  = 30 * time.Second
)

func main() {
	// Load dot env file
	err := godotenv.Load(DOT_ENV_FILENAME)
	if err != nil {
		log.Println("No .env file detected")
	}

	// env variables
	uiBuildFolderPath := util.Getenv(UI_BUILD_FOLDER_KEY, DEFAULT_UI_BUILD_FOLDER_PATH)
	uiIndexFilename := util.Getenv(UI_BUILD_INDEX_FILENAME_KEY, DEFAULT_UI_BUILD_INDEX_FILENAME)
	uiClientURL := util.Getenv(UI_CLIENT_URL_KEY, DEFAULT_CLIENT_URL)
	env := util.Getenv(ENV_KEY, ENV_DEVELOPMENT)
	// Init mux router
	r := mux.NewRouter()
	primeHandler := rest.NewPrimeHandler()
	primeHandler.Register(r)

	// Handle react build in production
	spa := rest.SPAHandler{StaticPath: uiBuildFolderPath, IndexPath: uiIndexFilename}
	r.PathPrefix("/").Handler(spa)

	// Cors
	corsOptions := cors.Options{}
	if env != ENV_PRODUCTION {
		corsOptions.AllowedOrigins = []string{uiClientURL}
	}

	corsInstance := cors.New(corsOptions)

	srv := &http.Server{
		Handler:      corsInstance.Handler(r),
		Addr:         HTTP_BINDING_ADDR,
		WriteTimeout: HTTP_SERVER_WRITE_TIMEOUT_DURATION,
		ReadTimeout:  HTTP_SERVER_READ_TIMEOUT_DURATION,
	}

	log.Println("Binding http server to ", HTTP_BINDING_ADDR)
	log.Fatal(srv.ListenAndServe())
}
