package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/DenHafiz69/waqt/internal/prayer"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading .env: %v", err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("GET /healthz", healthz)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Println("Listening on :8080")
	log.Fatal(srv.ListenAndServe())
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := prayer.GetDailyPrayerTime()

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("encoding response: %v", err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
