package controllers

import (
	"net/http"

	"aero-api/src/models"
)

func GetFlights(w http.ResponseWriter, r *http.Request) {
	flights := []models.Flight{{FlightNumber: "AA101"}, {FlightNumber: "UA202"}}
	json.NewEncoder(w).Encode(flights)
}
