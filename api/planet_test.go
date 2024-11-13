package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlanet(t *testing.T) {
	planetMockResponse := `{
		"name": "Tatooine"
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(planetMockResponse))
	}))
	defer server.Close()

	planetName, err := GetPlanet(server.URL)
	if err != nil {
		t.Fatalf("Failed to get planets: %v", err)
	}

	expectedName := "Tatooine"
	if planetName != expectedName {
		t.Errorf("Expect name '%s', but actual name is '%s'", expectedName, planetName)
	}
}
