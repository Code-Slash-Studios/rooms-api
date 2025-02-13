package routes

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	api := mux.NewRouter()
	api.HandleFunc("/reservations", getReservations).Methods("GET")
	api.HandleFunc("/reservations/{id}", getReservation).Methods("GET")
	api.HandleFunc("/reservations", createReservation).Methods("POST")
	api.HandleFunc("/reservations/{id}", updateReservation).Methods("PUT")
	api.HandleFunc("/reservations/{id}", deleteReservation).Methods("DELETE")

	return api
}
