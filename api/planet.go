package api

import (
	"encoding/json"
	"net/http"
	"star-wars-app/models"
)

func GetPlanet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var planet models.Planet
	if err := json.NewDecoder(resp.Body).Decode(&planet); err != nil {
		return "", err
	}

	return planet.Name, nil
}
