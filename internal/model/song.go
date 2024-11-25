package model

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Group  string `gorm:"group"`
	Name   string `gorm:"name"`
	Lyrics string `gorm:"lyrics"`
}

type SongRequestCreate struct {
	Group string `json:"group" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type SongRequestUpdate struct {
	Group  string `json:"group"`
	Name   string `json:"name"`
	Lyrics string `json:"lyrics"`
}

type SongResponse struct {
	Group  string `json:"group" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Lyrics string `json:"lyrics"`
}

func (s *Song) ToResponse() *SongResponse {
	return &SongResponse{
		Group:  s.Group,
		Name:   s.Name,
		Lyrics: s.Lyrics,
	}
}
