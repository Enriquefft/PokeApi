package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

func buildUrl(args ...string) string {
	var urlBuilder strings.Builder
	for _, arg := range args {
		urlBuilder.WriteString(arg)
	}
	return urlBuilder.String()
}

type MediaType int

// Enum
const (
	Image MediaType = iota
	Sound
)

func checkUrl(url *string, media_type MediaType) error {
	resp, err := http.Head(*url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")

	switch media_type {
	case Image:
		if !strings.HasPrefix(contentType, "image/") {
			return fmt.Errorf("Invalid image file %s", *url)
		}
	case Sound:
		if !strings.HasPrefix(contentType, "audio/mpeg") {
			return fmt.Errorf("Invalid MP3 file %s", *url)
		}
	}

	return nil
}

func buildImgUrl(pokemon_name string) (string, error) {
	// image url
	// https://play.pokemonshowdown.com/sprites/dex/POKEMON_NAME.png
	img_url := buildUrl("https://play.pokemonshowdown.com/sprites/dex/", pokemon_name, ".png")

	log.Printf("Img url: %s", img_url)

	img_err := checkUrl(&img_url, Image)

	if img_err != nil {
		img_url = ""
	}
	return img_url, img_err
}

func buildCryUrl(pokemon_name string) (string, error) {
	// Sound url
	// https://play.pokemonshowdown.com/audio/cries/POKEMON_NAME.mp3
	sound_url := buildUrl("https://play.pokemonshowdown.com/audio/cries/", pokemon_name, ".mp3")
	sound_err := checkUrl(&sound_url, Sound)

	if sound_err != nil {
		sound_url = ""
	}
	return sound_url, sound_err
}

func getSmallInfo(id int) *PokemonSmallInfo {
	pokemon_info, db_err := GetSmallInfo(id)

	if db_err != nil {
		// Failed to get ID info: error msg
		log.Fatalf("Failed to get %d DB small info: %s", id, db_err)
	}

	lower_name := strings.ToLower(pokemon_info.Name)

	var img_err error
	pokemon_info.ImgUrl, img_err = buildImgUrl(lower_name)

	log.Printf("Img url: %s", pokemon_info.ImgUrl)

	if img_err != nil {
		log.Fatalf("Failed to get %s image: %s", pokemon_info.Name, img_err)
	}
	return pokemon_info
}

func capitalize(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func GetAllPokemonInfoByName(gin_ctx *gin.Context) {
	// Allow cors
	gin_ctx.Header("Access-Control-Allow-Origin", "*")
	gin_ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	// Get main info
	name := capitalize(gin_ctx.Param("name"))

	pokemon_info, db_err := getInfoByName(name)
	if db_err != nil {
		log.Fatalf("Failed to get %s DB info: %s", gin_ctx.Param("id"), db_err)
	}

	id := strconv.Itoa(pokemon_info.NationalID)

	// Get evolutions
	var db_evo_err error
	pokemon_info.Evolutions, db_evo_err = GetEvolutions(id)

	if db_evo_err != nil {
		log.Fatalf("Failed to get %s DB evolutions: %s", gin_ctx.Param("id"), db_evo_err)
	}

	// Get pre evolution if exists
	if pokemon_info.Pre_evolution_id != nil {
		pokemon_info.PreEvolution = getSmallInfo(*pokemon_info.Pre_evolution_id)
	}

	// Get media urls
	var sound_err error
	var img_err error

	lower_name := strings.ToLower(pokemon_info.Name)

	pokemon_info.SoundUrl, sound_err = buildCryUrl(lower_name)
	pokemon_info.ImgUrl, img_err = buildImgUrl(lower_name)

	if sound_err != nil {
		log.Fatalf("Failed to get %s sound: %s", pokemon_info.Name, sound_err)
	}
	if img_err != nil {
		log.Fatalf("Failed to get %s image: %s", pokemon_info.Name, img_err)
	}

	gin_ctx.JSON(http.StatusOK, pokemon_info)
}

func GetAllPokemonInfoById(gin_ctx *gin.Context) {
	// Allow cors
	gin_ctx.Header("Access-Control-Allow-Origin", "*")
	gin_ctx.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	// Get main info
	id := gin_ctx.Param("id")

	_, id_err := strconv.Atoi(id)

	if id_err != nil {
		gin_ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		log.Fatalf("Failed to convert %s to int: %s", gin_ctx.Param("id"), id_err)
		return
	}

	pokemon_info, db_err := getInfoById(id)
	if db_err != nil {
		log.Fatalf("Failed to get %s DB info: %s", gin_ctx.Param("id"), db_err)
	}

	// Get evolutions
	var db_evo_err error
	pokemon_info.Evolutions, db_evo_err = GetEvolutions(id)

	if db_evo_err != nil {
		log.Fatalf("Failed to get %s DB evolutions: %s", gin_ctx.Param("id"), db_evo_err)
	}

	// Get pre evolution if exists
	if pokemon_info.Pre_evolution_id != nil {
		pokemon_info.PreEvolution = getSmallInfo(*pokemon_info.Pre_evolution_id)
	}

	// Get media urls
	var sound_err error
	var img_err error

	lower_name := strings.ToLower(pokemon_info.Name)

	pokemon_info.SoundUrl, sound_err = buildCryUrl(lower_name)
	pokemon_info.ImgUrl, img_err = buildImgUrl(lower_name)

	if sound_err != nil {
		log.Fatalf("Failed to get %s sound: %s", pokemon_info.Name, sound_err)
	}
	if img_err != nil {
		log.Fatalf("Failed to get %s image: %s", pokemon_info.Name, img_err)
	}

	gin_ctx.JSON(http.StatusOK, pokemon_info)
}
