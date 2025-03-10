package routes

import (
	"encoding/json"
	"net/http"

	. "rooms-api/src/models"

	"github.com/gorilla/mux"
)

var reservations = make(map[string]Reservation)

func getReservations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ReservationList []Reservation
	for _, reservation := range reservations {
		ReservationList = append(ReservationList, reservation)
	}
	json.NewEncoder(w).Encode(ReservationList)
}

func getReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reservation, exists := reservations[params["id"]]
	if !exists {
		http.Error(w, "Reservation not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(reservation)
}

func createReservation(w http.ResponseWriter, r *http.Request) {
	var reservation Reservation
	json.NewDecoder(r.Body).Decode(&reservation)
	reservations[reservation.ID] = reservation
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reservation)
}

func updateReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedReservation Reservation
	json.NewDecoder(r.Body).Decode(&updatedReservation)
	if _, exists := reservations[params["id"]]; !exists {
		http.Error(w, "Reservation not found", http.StatusNotFound)
		return
	}
	updatedReservation.ID = params["id"]
	reservations[params["id"]] = updatedReservation
	json.NewEncoder(w).Encode(updatedReservation)
}

func deleteReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, exists := reservations[params["id"]]; !exists {
		http.Error(w, "Reservation not found", http.StatusNotFound)
		return
	}
	delete(reservations, params["id"])
	w.WriteHeader(http.StatusNoContent)
}
