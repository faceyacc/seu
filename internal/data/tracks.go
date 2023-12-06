package data

import "time"

type Track struct {
	ID          int64             `json:"id"`
	CreatedAt   time.Time         `json:"created_at"`
	Title       string            `json:"title"`
	Artist      []string          `json:"artist"`
	Year        int32             `json:"year,omitempty,string"`
	Description string            `json:"description,omitempty"`
	Duration    int64             `json:"track_duration,omitempty"` // In seconds
	Genres      []string          `json:"genre,omitempty"`
	Links       map[string]string `json:"links"`
	Version     int32             `json:"-"`
}
