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

func createReservation(w http.ResponseWriter, r *http.Request) {
	var reservation models.Reservation
	json.NewDecoder(r.Body).Decode(&reservation)
	reservations[reservation.ID] = reservation
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reservation)
}

func updateReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedReservation models.Reservation
	json.NewDecoder(r.Body).Decode(&updatedReservation)
	if _, exists := reservations[params["id"]]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	updatedReservation.ID = params["id"]
	reservations[params["id"]] = updatedReservation
	json.NewEncoder(w).Encode(updatedReservation)
}

func deleteReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, exists := reservations[params["id"]]; !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(reservations, params["id"])
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	api := mux.NewRouter()
	api.handlefunc("/reservations", getReservations).Methods("GET")
	api.handlefunc("/reservations/{id}", getReservation).Methods("GET")
	api.HandleFunc("/users", createReservation).Methods("POST")
	api.HandleFunc("/users/{id}", updateReservation).Methods("PUT")
	api.HandleFunc("/users/{id}", deleteReservation).Methods("DELETE")

	fmt.Printf("API listening on port 6000")
	http.ListenAndServe(":6000", api)
}
