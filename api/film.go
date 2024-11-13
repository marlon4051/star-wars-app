package api

import (
	"encoding/json"
	"net/http"
	"star-wars-app/models"
)

func GetFilms(apiURL string) ([]models.Film, error) {
	resp, err := http.Get(apiURL + "/films/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var filmsResp struct {
		Results []models.Film `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&filmsResp); err != nil {
		return nil, err
	}

	return filmsResp.Results, nil
}
