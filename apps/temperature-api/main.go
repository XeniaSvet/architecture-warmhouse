package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Response структура для JSON-ответа
type TemperatureResponse struct {
	Location    string  `json:"location"`
	SensorID    string  `json:"sensorId"`
	Temperature float64 `json:"temperature"`
	Unit        string  `json:"unit"`
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	location := r.URL.Query().Get("location")
	sensorID := r.URL.Query().Get("sensorId")

	// Если не указана локация, определяем её по sensorId
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// Если не указан sensorId, определяем его по локации
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}

	randomTemp := 18.0 + rand.Float64()*(26.0-18.0)

	response := TemperatureResponse{
		Location:    location,
		SensorID:    sensorID,
		Temperature: jsonRound(randomTemp, 1),
		Unit:        "Celsius",
	}

	json.NewEncoder(w).Encode(response)
}

func jsonRound(val float64, precision int) float64 {
	p := 1.0
	for i := 0; i < precision; i++ {
		p *= 10
	}
	return float64(int(val*p)) / p
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/temperature", temperatureHandler)

	fmt.Println("Сервер запущен на http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}
