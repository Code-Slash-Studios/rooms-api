package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	. "rooms-api/src/models"

	"github.com/gorilla/mux"
)

func GetRooms(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT id, name, department FROM rooms")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	rooms := []Room{}
	for rows.Next() {
		var room Room
		if err := rows.Scan(&room.ID, &room.Name, &room.Department); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rooms = append(rooms, room)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var room Room
	err := DB.QueryRow("SELECT id, name, department FROM rooms WHERE id = ?", id).Scan(&room.ID, &room.Name, &room.Department)
	if err == sql.ErrNoRows {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := DB.Exec("INSERT INTO rooms (name, department) VALUES (?, ?)", room.Name, room.Department)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := result.LastInsertId()
	room.ID = strconv.FormatInt(id, 10)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var room Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := DB.Exec("UPDATE rooms SET name = ?, department = ? WHERE id = ?", room.Name, room.Department, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	room.ID = room.ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	_, err := DB.Exec("DELETE FROM rooms WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
