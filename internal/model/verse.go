package model

type Verse struct {
	Number int      `json:"number"`
	Song   string   `json:"song"`
	Band   string   `json:"band"`
	Lines  []string `json:"lines"`
}
