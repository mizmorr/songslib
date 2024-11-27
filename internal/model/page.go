package model

type Page struct {
	Number int `json:"number"`
	Size   int `json:"size"`
}

func NewPage(number, size int) *Page {
	return &Page{Number: number, Size: size}
}
