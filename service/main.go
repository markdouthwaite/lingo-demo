package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/markdouthwaite/lingo"
	"log"
	"net/http"
	"os"
	"time"
)

type response struct {
	Status string            `json:"status,omitempty"`
	Errors map[string]string `json:"errors,omitempty"`
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	code := http.StatusOK
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response{
		Status: http.StatusText(code),
	})
}

func getEnvWithDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func main() {
	host := getEnvWithDefault("HOST", "0.0.0.0")
	port := getEnvWithDefault("PORT", "8000")

	model := lingo.LoadClassifier("artifacts/breast-cancer-1.h5")
	handler := lingo.NewClassifierHandler(model)

	router := mux.NewRouter()
	router.HandleFunc("/predict", handler).Methods("POST")
	router.HandleFunc("/health", healthCheck).Methods("GET")

	server := &http.Server{
		Handler:      router,
		Addr:         host + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}
