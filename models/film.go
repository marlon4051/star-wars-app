package models

type Film struct {
	Title       string   `json:"title"`
	ReleaseDate string   `json:"release_date"`
	Planets     []string `json:"planets"`
}
