package model

import "time"

type Event struct {
	UUID   int       `json:"uuid"`
	UserID int       `json:"userID"`
	Event  string    `json:"event"`
	Date   time.Time `json:"date"`
}
