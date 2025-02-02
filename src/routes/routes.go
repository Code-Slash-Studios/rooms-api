package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models/"
	"github.com/gorilla/mux"
)

var reservations = make(map[string]models.Reservation)

func getReservations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ReservationList []models.Reservation
	for _, reservation := range reservations {
		ReservationList = append(ReservationList, reservation)
	}
	json.NewEncoder(w).Encode(ReservationList)
}

func getReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reservation, exists := reservations[params["id"]]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(reservation)
}

func main() {
	api := mux.NewRouter()
	api.handlefunc("/reservations", getReservations).Methods("GET")
	api.handlefunc("/reservations/{id}", getReservation).Methods("GET")

	fmt.Printf("API listening on port 6000")
	http.ListenAndServe(":6000", api)
}
