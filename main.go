package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0)
	return nil
}

func (u UnixTime) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, "%d", u.Time.Unix()), nil
}

type DailyPrayerTime struct {
	Day     int      `json:"day"`
	Hijri   UnixTime `json:"hijri"`
	Imsak   UnixTime `json:"imsak"`
	Fajr    UnixTime `json:"fajr"`
	Syuruk  UnixTime `json:"syuruk"`
	Dhuha   UnixTime `json:"dhuha"`
	Dhuhr   UnixTime `json:"dhuhr"`
	Asr     UnixTime `json:"asr"`
	Maghrib UnixTime `json:"maghrib"`
	Isha    UnixTime `json:"isha"`
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

	requestUrl := "https://api.waktusolat.app/v2/solat/gps/3.068498/101.630263?year=2026&month=8"

	req, _ := http.NewRequest("GET", requestUrl, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	schedule := MonthlySchedule{}

	if err := json.Unmarshal(body, &schedule); err != nil {
		fmt.Printf("Unmarshal failed: %v", err)
	}

	fmt.Println(schedule)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go!")
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
