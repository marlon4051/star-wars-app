package main

import (
	"encoding/json"
	"fmt"
	"log"
	"star-wars-app/api"
	"star-wars-app/utils"
)

func main() {
	// get films
	films, err := api.GetFilms("https://swapi.dev/api")
	if err != nil {
		log.Fatalf("Failed to get films: %v", err)
	}

	films = utils.SortFilms(films)
	result := make(map[string][]string)

	for _, film := range films {
		planetNames := []string{}
		for _, planetURL := range film.Planets {
			planet, err := api.GetPlanet(planetURL)
			if err != nil {
				log.Printf("Error al obtener los planetas: %v", err)
				continue
			}
			planetNames = append(planetNames, planet)
		}

		uniquePlanets := make(map[string]struct{})
		for _, planet := range planetNames {
			uniquePlanets[planet] = struct{}{}
		}

		var uniquePlanetNames []string
		for planet := range uniquePlanets {
			uniquePlanetNames = append(uniquePlanetNames, planet)
		}
		uniquePlanetNames = utils.SortPlanets(uniquePlanetNames)

		result[film.Title] = uniquePlanetNames
	}
	prettyJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatalf("Error al generar el JSON: %v", err)
	}

	fmt.Println(string(prettyJSON))
}
