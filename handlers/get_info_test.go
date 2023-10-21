package db

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	// Testing
	"github.com/stretchr/testify/assert"
)

func TestGetPokemonInfoById(t *testing.T) {
	log.Println("Testing GetPokemonInfoById")

	InitDB()

	// Initialize the Gin router
	router := gin.Default()

	// Set up a test route that calls the GetPokemonInfoById function
	router.GET("/pokemon/:id", GetAllPokemonInfoById)

	// Create a test request to the route
	req, err := http.NewRequest("GET", "/pokemon/281", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	recorder := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code 200")

	// Check the response body
	expectedBody := `{"national_id":281,"name":"Kirlia","speed":50,"s_def":65,"s_attack":65,"def":35,"attack":45,"hp":38,"type1":"Psychic","type2":"Fairy","height":0.8,"weight":20.2,"ability1":"Synchronize","ability2":"Trace","hidden_ability":"Telepathy","pre_evo":{"national_id":280,"name":"Ralts","img_url":"https://play.pokemonshowdown.com/sprites/dex/ralts.png"},"evolutions":[{"national_id":282,"name":"Gardevoir","img_url":"https://play.pokemonshowdown.com/sprites/dex/gardevoir.png"},{"national_id":475,"name":"Gallade","img_url":"https://play.pokemonshowdown.com/sprites/dex/gallade.png"}],"img_url":"https://play.pokemonshowdown.com/sprites/dex/kirlia.png","sound_url":"https://play.pokemonshowdown.com/audio/cries/kirlia.mp3"}`

	assert.Equal(t, expectedBody, recorder.Body.String(), "Response body does not match expected")
	// INSERT INTO pokemons VALUES (281, 'Kirlia', 50, 65, 65, 35, 45, 38, 'Psychic', 'Fairy', 0.8, 20.2, 'Synchronize', 'Trace', 'Telepathy', 280);
}
