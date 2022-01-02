package models

import "time"

type Task struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Created     time.Time `json:"date_of_creation"`
	TagId       int
}
