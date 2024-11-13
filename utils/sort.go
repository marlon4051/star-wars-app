package utils

import (
	"sort"
	"star-wars-app/models"
	"time"
)

func SortPlanets(planetNames []string) []string {
	sort.Strings(planetNames)
	return planetNames
}

func SortFilms(films []models.Film) []models.Film {
	sort.Slice(films, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02", films[i].ReleaseDate)
		dateJ, _ := time.Parse("2006-01-02", films[j].ReleaseDate)
		return dateI.Before(dateJ)
	})
	return films
}
