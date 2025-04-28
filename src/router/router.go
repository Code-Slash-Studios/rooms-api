package router

import (
	"github.com/gorilla/mux"

	"rooms-api/src/routes"
)

func Router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	api := r.PathPrefix("/api").Subrouter()
	pi := r.PathPrefix("/pi").Subrouter()

	api.HandleFunc("/reservations", routes.GetReservations).Methods("GET")
	api.HandleFunc("/reservations/{id}", routes.GetReservation).Methods("GET")
	api.HandleFunc("/reservations", routes.CreateReservation).Methods("POST")
	api.HandleFunc("/reservations/{id}", routes.UpdateReservation).Methods("PUT")
	api.HandleFunc("/reservations/{id}", routes.DeleteReservation).Methods("DELETE")
	api.HandleFunc("/reservations/room/{room_id}", routes.GetReservationsByRoom).Methods("GET")
	api.HandleFunc("/reservations/user/{user_id}", routes.GetReservationsByUser).Methods("GET")
	api.HandleFunc("/reservations/rooms", routes.GetNextReservations).Methods("GET")

	api.HandleFunc("/rooms", routes.GetRooms).Methods("GET")
	api.HandleFunc("/rooms/{id}", routes.GetRoom).Methods("GET")
	api.HandleFunc("/rooms", routes.CreateRoom).Methods("POST")
	api.HandleFunc("/rooms/{id}", routes.UpdateRoom).Methods("PUT")
	api.HandleFunc("/rooms/{id}", routes.DeleteRoom).Methods("DELETE")

	api.HandleFunc("/users", routes.GetUser).Methods("GET")
	api.HandleFunc("/users/{id}", routes.GetUser).Methods("GET")
	api.HandleFunc("/users", routes.CreateUser).Methods("POST")
	api.HandleFunc("/users/{id}", routes.UpdateUser).Methods("PUT")
	api.HandleFunc("/users/{id}", routes.DeleteUser).Methods("DELETE")

	pi.HandleFunc("/", routes.GetReservations).Methods("GET")
	pi.HandleFunc("/{room_id}", routes.GetReservationsByRoom).Methods("GET")

	return r
}
