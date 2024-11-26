package model

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Band   string `gorm:"band"`
	Name   string `gorm:"name"`
	Lyrics string `gorm:"lyrics"`
}

type SongRequestCreate struct {
	Band string `json:"band" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type SongRequestUpdate struct {
	Band   string `json:"band"`
	Name   string `json:"name"`
	Lyrics string `json:"lyrics"`
}

type SongResponse struct {
	Band   string `json:"band" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Lyrics string `json:"lyrics"`
}

func (s *Song) ToResponse() *SongResponse {
	return &SongResponse{
		Band:   s.Band,
		Name:   s.Name,
		Lyrics: s.Lyrics,
	}
}
