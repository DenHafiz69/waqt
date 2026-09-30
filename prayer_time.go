package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
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
	Hijri   string   `json:"hijri"`
	Imsak   UnixTime `json:"imsak"`
	Fajr    UnixTime `json:"fajr"`
	Syuruk  UnixTime `json:"syuruk"`
	Dhuha   UnixTime `json:"dhuha"`
	Dhuhr   UnixTime `json:"dhuhr"`
	Asr     UnixTime `json:"asr"`
	Maghrib UnixTime `json:"maghrib"`
	Isha    UnixTime `json:"isha"`
}

type MonthlyPrayerTime struct {
	Zone        string            `json:"zone"`
	Year        int               `json:"year"`
	Month       string            `json:"month"`
	MonthNumber int               `json:"month_number"`
	LastUpdated string            `json:"last_updated"`
	Prayers     []DailyPrayerTime `json:"prayers"`
}

func GetMonthlyPrayerTime(Month int) MonthlyPrayerTime {

	requestUrl := "https://api.waktusolat.app/v2/solat/gps/3.068498/101.630263?year=2026&month=8"

	req, _ := http.NewRequest("GET", requestUrl, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("request error: %v", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("io read error: %v", err)
	}

	schedule := MonthlyPrayerTime{}

	if err := json.Unmarshal(body, &schedule); err != nil {
		log.Printf("unmarshal error: %v", err)
	}

	return schedule
}

func GetAndWriteNewMonthToFile(currentMonth int) (MonthlyPrayerTime, error) {

	schedule := GetMonthlyPrayerTime(currentMonth)

	file, err := os.Create("monthly_schedule.json")
	if err != nil {
		return MonthlyPrayerTime{}, fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(schedule); err != nil {
		return MonthlyPrayerTime{}, fmt.Errorf("error encoding file: %v", err)
	}

	return schedule, nil
}

func GetLocalSchedule(filepath string) (MonthlyPrayerTime, error) {

	currentMonth := int(time.Now().Month())
	currentYear := int(time.Now().Year())

	// If current month and year match, use local schedule
	file, err := os.Open(filepath)
	if errors.Is(err, fs.ErrNotExist) {
		// If file not exist, get the file, and return the schedule
		schedule, err := GetAndWriteNewMonthToFile(currentMonth)
		if err != nil {
			return MonthlyPrayerTime{}, fmt.Errorf("error getting monthly schedule: %v", err)
		}
		return schedule, nil
	} else if err != nil {
		return MonthlyPrayerTime{}, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	schedule := MonthlyPrayerTime{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&schedule); err != nil {
		return MonthlyPrayerTime{}, fmt.Errorf("error decoding file: %v", err)
	}

	// Check if month and year match with local schedule
	scheduleMonth, err := strconv.Atoi(schedule.Month)
	if err != nil {
		return MonthlyPrayerTime{}, fmt.Errorf("convert to int error: %v", err)
	}

	scheduleYear := int(schedule.Year)

	if scheduleMonth != currentMonth || scheduleYear != currentYear {
		schedule, err := GetAndWriteNewMonthToFile(currentMonth)
		if err != nil {
			return MonthlyPrayerTime{}, fmt.Errorf("error getting monthly schedule: %v", err)
		}
		return schedule, nil
	}

	return schedule, nil
}

func GetDailyPrayerTime() DailyPrayerTime {

	// Call the GetLocalSchedule
	schedule, err := GetLocalSchedule("monthly_schedule.json")
	if err != nil {
		log.Printf("error getting local schedule: %v", err)
	}

	day := time.Now().Day()

	fmt.Println("Day:", day)

	TodaySchedule := DailyPrayerTime{}

	for _, v := range schedule.Prayers {
		if v.Day == day {
			TodaySchedule = v
			break
		}
	}

	return TodaySchedule
}
