package main

import (
	"fmt"
	"net/http"
	"rooms-api/src/routes"
)

//func main() {
//	fmt.Println("Hello World")
//	_index.main()
//}

//func main() {
//	api := mux.NewRouter()
//	api.HandleFunc("/reservations", getReservations).Methods("GET")
//	api.HandleFunc("/reservations/{id}", getReservation).Methods("GET")
//	api.HandleFunc("/users", createReservation).Methods("POST")
//	api.HandleFunc("/users/{id}", updateReservation).Methods("PUT")
//	api.HandleFunc("/users/{id}", deleteReservation).Methods("DELETE")
//
//	fmt.Printf("API listening on port 6000")
//	http.ListenAndServe(":6000", api)
//}

func main() {
	router := routes.NewRouter()
	fmt.Println("Server is listening on port 6000")
	http.ListenAndServe(":6000", router)
}
