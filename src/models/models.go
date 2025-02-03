package models

import (
	"time"
)

type Reservation struct {
	ID    string
	Name  string
	Room  string
	Start time.Time
	End   time.Time
}

func (r Reservation) getID() string {
	return r.ID
}

func main() {}
