package routes

import (
	"encoding/json"
	"net/http"

	. "rooms-api/src/models"

	"github.com/gorilla/mux"
)

var rooms = make(map[string]Room)

func getRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var RoomList []Room
	for _, room := range rooms {
		RoomList = append(RoomList, room)
	}
	json.NewEncoder(w).Encode(RoomList)
}

func getRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	room, exists := rooms[params["id"]]
	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(room)
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	var room Room
	json.NewDecoder(r.Body).Decode(&room)
	rooms[room.ID] = room
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}

func updateRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedRoom Room
	json.NewDecoder(r.Body).Decode(&updatedRoom)
	if _, exists := rooms[params["id"]]; !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}
	updatedRoom.ID = params["id"]
	rooms[params["id"]] = updatedRoom
	json.NewEncoder(w).Encode(updatedRoom)
}

func deleteRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if _, exists := rooms[params["id"]]; !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}
	delete(rooms, params["id"])
	w.WriteHeader(http.StatusNoContent)
}
