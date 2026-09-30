package main

import (
	"encoding/json"
	"github.com/DenHafiz69/waqt/internal/prayer"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("error loading .env: %v", err)
		return
	}

	http.HandleFunc("/", homeHandler)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
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
