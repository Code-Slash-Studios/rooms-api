package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	api := mux.NewRouter()
	api.HandleFunc("/reservations", getReservations).Methods("GET")
	api.HandleFunc("/reservations/{id}", getReservation).Methods("GET")
	api.HandleFunc("/users", createReservation).Methods("POST")
	api.HandleFunc("/users/{id}", updateReservation).Methods("PUT")
	api.HandleFunc("/users/{id}", deleteReservation).Methods("DELETE")

	//fmt.Printf("API listening on port 6000")
	//http.ListenAndServe(":6000", api)
	return api
}
