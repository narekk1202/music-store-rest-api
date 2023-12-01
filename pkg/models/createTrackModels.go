package models

type CreateTrack struct {
	Artist string `json:"artist" binding:"required"`
	Title  string `json:"title" binding:"required"`
}
