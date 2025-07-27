package internal

import "time"

type Event struct {
	Source    string    `json:"source"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	Timestamp time.Time `json:"timestamp"`
}
