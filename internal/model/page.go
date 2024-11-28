package model

type Page struct {
	Number int `form:"number"`
	Size   int `form:"size"`
}

func NewPage(number, size int) *Page {
	return &Page{Number: number, Size: size}
}
