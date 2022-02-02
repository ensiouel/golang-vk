package objects

import (
	"fmt"
	"strings"
)

type PhotoSize struct {
	Type   string `json:"type"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type PhotoSizes []PhotoSize

type Photo struct {
	ID      int64      `json:"id"`
	AlbumID int64      `json:"album_id"`
	OwnerID int64      `json:"owner_id"`
	UserID  int64      `json:"user_id"`
	Text    string     `json:"text"`
	Date    int64      `json:"date"`
	Sizes   PhotoSizes `json:"sizes"`
	Width   int        `json:"width,omitempty"`
	Height  int        `json:"height,omitempty"`
}

type Crop struct {
	X  int `json:"x"`  // Координата X левого верхнего угла в процентах.
	Y  int `json:"y"`  // Координата Y левого верхнего угла в процентах.
	X2 int `json:"x2"` // Координата X правого нижнего угла в процентах.
	Y2 int `json:"y2"` // Координата Y правого нижнего угла в процентах.
}

type Rect struct {
	X  int `json:"x"`  // Координата X левого верхнего угла в процентах.
	Y  int `json:"y"`  // Координата Y левого верхнего угла в процентах.
	X2 int `json:"x2"` // Координата X правого нижнего угла в процентах.
	Y2 int `json:"y2"` // Координата Y правого нижнего угла в процентах.
}

type Image struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Photos []Photo

func (sizes *PhotoSizes) Max() *PhotoSize {
	return &(*sizes)[len(*sizes)-1]
}

func (sizes *PhotoSizes) Min() *PhotoSize {
	return &(*sizes)[0]
}

func (photo *Photo) ToAttachments() string {
	return fmt.Sprintf("photo%d_%d", photo.OwnerID, photo.ID)
}

func (photos *Photos) ToAttachments() string {
	var attachments []string

	for _, photo := range *photos {
		attachments = append(attachments, photo.ToAttachments())
	}
	return strings.Join(attachments, ",")
}
