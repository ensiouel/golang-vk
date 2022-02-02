package objects

import "time"

type Audio struct {
	ID       int64         `json:"id"`
	OwnerID  int64         `json:"owner_id"`
	Artist   string        `json:"artist"`
	Title    string        `json:"title"`
	Duration time.Duration `json:"duration"`
	Url      string        `json:"url"`
	LyricsID int64         `json:"lyrics_id,omitempty"`
	AlbumID  int64         `json:"album_id,omitempty"`
	GenreID  int64         `json:"genre_id"`
	Date     int64         `json:"date"`
	NoSearch int           `json:"no_search,omitempty"`
	IsHQ     int           `json:"is_hq"`
}
