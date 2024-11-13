package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Respuesta simulada de la API
var filmsMockResponse = `{
    "results": [
        {
            "title": "A New Hope",
            "release_date": "1977-05-25",
            "planets": ["http://swapi.dev/api/planets/1/"]
        },
        {
            "title": "The Empire Strikes Back",
            "release_date": "1980-05-17",
            "planets": ["http://swapi.dev/api/planets/2/"]
        }
    ]
}`

func TestGetFilms(t *testing.T) {
	// mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(filmsMockResponse))
	}))
	defer server.Close()

	// mock server
	films, err := GetFilms(server.URL)
	if err != nil {
		t.Fatalf("Failed to get films: %v", err)
	}

	if len(films) != 2 {
		t.Fatalf("Waiting for to films %d", len(films))
	}

	expectedTitle := "A New Hope"
	if films[0].Title != expectedTitle {
		t.Errorf("Se esperaba el título '%s', pero se obtuvo '%s'", expectedTitle, films[0].Title)
	}
}
