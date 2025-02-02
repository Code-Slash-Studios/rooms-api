package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models/"
	"github.com/gorilla/mux"
)

func main() {
	var res models.Reservation
	var response http.Response
	var example string

	fmt.Printf("Test")
}

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
