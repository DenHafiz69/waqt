package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type DailyPrayerTime struct {
	Day     int       `json:"day"`
	Hijri   time.Time `json:"hijri"`
	Imsak   time.Time `json:"imsak"`
	Fajr    time.Time `json:"fajr"`
	Syuruk  time.Time `json:"syuruk"`
	Dhuha   time.Time `json:"dhuha"`
	Dhuhr   time.Time `json:"dhuhr"`
	Asr     time.Time `json:"asr"`
	Maghrib time.Time `json:"maghrib"`
	Isha    time.Time `json:"isha"`
}

type MonthlySchedule struct {
	Zone        string            `json:"zone"`
	Year        int               `json:"year"`
	Month       string            `json:"month"`
	MonthNumber int               `json:"month_number"`
	LastUpdated string            `json:"last_updated"`
	Prayers     []DailyPrayerTime `json:"prayers"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go!")
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
