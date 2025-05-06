package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	. "rooms-api/src/models"

	"github.com/gorilla/mux"
)

func GetReservations(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("SELECT ID, RoomID, Name, UserID, Start, End FROM reservations order by Start ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	reservations := []Reservation{}
	for rows.Next() {
		var res Reservation
		if err := rows.Scan(&res.ID, &res.RoomID, &res.Name, &res.UserID, &res.Start, &res.End); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reservations = append(reservations, res)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}

func GetReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var res Reservation
	query := "SELECT ID, RoomID, Name, UserID, Start, End FROM reservations WHERE ID = ?"
	err := DB.QueryRow(query, id).Scan(&res.ID, &res.RoomID, &res.Name, &res.UserID, &res.Start, &res.End)
	if err == sql.ErrNoRows {
		http.Error(w, "Reservation `id` not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func GetReservationsByRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomID := params["room_id"]
	rows, err := DB.Query("SELECT ID, RoomID, Name, UserID, Start, End FROM reservations WHERE RoomID = ? order by Start ASC", roomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	reservations := []Reservation{}
	for rows.Next() {
		var res Reservation
		if err := rows.Scan(&res.ID, &res.RoomID, &res.Name, &res.UserID, &res.Start, &res.End); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reservations = append(reservations, res)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}

func GetReservationsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]
	rows, err := DB.Query("SELECT ID, RoomID, Name, UserID, Start, End FROM reservations WHERE UserID = ? order by Start ASC", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	reservations := []Reservation{}
	for rows.Next() {
		var res Reservation
		if err := rows.Scan(&res.ID, &res.RoomID, &res.Name, &res.UserID, &res.Start, &res.End); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reservations = append(reservations, res)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}

func CreateReservation(w http.ResponseWriter, r *http.Request) {
	var res Reservation
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var conflictCount int
	checkQuery := `SELECT COUNT(*) FROM appointments WHERE room_id = ? AND start_time < ? AND end_time > ?`
	checkErr := DB.QueryRow(checkQuery, res.RoomID, res.End, res.Start).Scan(&conflictCount)
	if checkErr != nil {
		http.Error(w, checkErr.Error(), http.StatusInternalServerError)
		return
	}
	if conflictCount > 0 {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, "Conflict detected: room is already booked")
		return
	}

	query := `INSERT INTO reservations (RoomID, Name, UserID, Start, End) VALUES (?, ?, ?, ?, ?)`
	result, err := DB.Exec(query, res.RoomID, res.Name, res.UserID, res.Start, res.End)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := result.LastInsertId()
	res.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpdateReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var res Reservation
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var conflictCount int
	checkQuery := `SELECT COUNT(*) FROM appointments WHERE room_id = ? AND start_time < ? AND end_time > ?`
	checkErr := DB.QueryRow(checkQuery, res.RoomID, res.End, res.Start).Scan(&conflictCount)
	if checkErr != nil {
		http.Error(w, checkErr.Error(), http.StatusInternalServerError)
		return
	}
	if conflictCount > 0 {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, "Conflict detected: room is already booked")
		return
	}

	query := `UPDATE reservations SET RoomID = ?, Name = ?, UserID = ?, Start = ?, End = ? WHERE ID = ?`
	_, err := DB.Exec(query, res.RoomID, res.Name, res.UserID, res.Start, res.End, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res.ID = res.ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func DeleteReservation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	_, err := DB.Exec("DELETE FROM reservations WHERE ID = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetNext(w http.ResponseWriter, r *http.Request) {
	roomRows, err := DB.Query("SELECT id, name, department FROM rooms")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer roomRows.Close()
	rooms := []Room{}
	for roomRows.Next() {
		var room Room
		if err := roomRows.Scan(&room.ID, &room.Name, &room.Department); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		rooms = append(rooms, room)
	}

	reservations := []Reservation{}
	for _, room := range rooms {
		resRow, err := DB.Query("SELECT ID, RoomID, Name, UserID, Start, End FROM reservations WHERE RoomID = ? AND Start > utc_timestamp order by Start ASC limit 1", room.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resRow.Close()
		for resRow.Next() {
			var res Reservation
			if err := resRow.Scan(&res.ID, &res.RoomID, &res.Name, &res.UserID, &res.Start, &res.End); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			reservations = append(reservations, res)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}
