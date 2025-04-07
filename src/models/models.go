package models

import (
	"container/list"
	"time"
)

type Reservation struct {
	ID     string    `json:"id"`
	RoomID string    `json:"room_id"`
	Name   string    `json:"name"`
	UserID string    `json:"user_id"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
}

type Room struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

type User struct {
	ID         string    `json:"id"`
	FName      string    `json:"fname"`
	LName      string    `json:"lname"`
	Department list.List `json:"department"`
}
