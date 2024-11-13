package utils

import (
	"reflect"
	"star-wars-app/models"
	"testing"
)

func TestSortPlanets(t *testing.T) {
	planets := []string{"Endor", "Tatooine", "Alderaan"}
	expected := []string{"Alderaan", "Endor", "Tatooine"}

	sortedPlanets := SortPlanets(planets)
	if !reflect.DeepEqual(sortedPlanets, expected) {
		t.Errorf("Expected result: %v, but get: %v", expected, sortedPlanets)
	}
}

func TestSortFilms(t *testing.T) {
	films := []models.Film{
		{Title: "The Empire Strikes Back", ReleaseDate: "1980-05-17"},
		{Title: "Return of the Jedi", ReleaseDate: "1983-05-25"},
		{Title: "A New Hope", ReleaseDate: "1977-05-25"},
	}
	expected := []models.Film{
		{Title: "A New Hope", ReleaseDate: "1977-05-25"},
		{Title: "The Empire Strikes Back", ReleaseDate: "1980-05-17"},
		{Title: "Return of the Jedi", ReleaseDate: "1983-05-25"},
	}

	sortedFilms := SortFilms(films)
	if !reflect.DeepEqual(sortedFilms, expected) {
		t.Errorf("Resultado esperado %v, pero se obtuvo %v", expected, sortedFilms)
	}
}
