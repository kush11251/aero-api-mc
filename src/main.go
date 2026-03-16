package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc="/flights", getFlights).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getFlights(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([{"flight_number": "AA101"}, {"flight_number": "UA202"}])
}