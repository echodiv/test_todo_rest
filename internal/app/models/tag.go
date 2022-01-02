package models

import "time"

type Tag struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"date_of_creation"`
}
