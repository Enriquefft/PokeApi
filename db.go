package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // Driver for sqlite3
)

var db *sqlx.DB

func InitDB() error {
	DB_PATH := "DB/Test.db"

	var err error
	db, err = sqlx.Open("sqlite3", DB_PATH)
	if err != nil {
		log.Fatal("Couldn't connect to database", err)
		return err
	}
	log.Println("Connected to database")
	return nil
}

type PokemonInfo struct {
	NationalID       int     `db:"national_id" json:"national_id"`
	Name             string  `db:"name" json:"name"`
	Speed            int     `db:"speed" json:"speed"`
	SpecialDefense   int     `db:"s_def" json:"s_def"`
	SpecialAttack    int     `db:"s_attack" json:"s_attack"`
	Defense          int     `db:"def" json:"def"`
	Attack           int     `db:"attack" json:"attack"`
	HP               int     `db:"hp" json:"hp"`
	Type1            string  `db:"type1" json:"type1"`
	Type2            *string `db:"type2" json:"type2,omitempty"`
	Height           float64 `db:"height" json:"height"`
	Weight           float64 `db:"weight" json:"weight"`
	Ability1         string  `db:"ability1" json:"ability1"`
	Ability2         *string `db:"ability2" json:"ability2,omitempty"`
	HiddenAbility    *string `db:"hidden_ability" json:"hidden_ability,omitempty"`
	Pre_evolution_id *int    `db:"pre_evo" json:"-"`

	PreEvolution *PokemonSmallInfo  `json:"pre_evo,omitempty"`
	Evolutions   []PokemonSmallInfo `json:"evolutions,omitempty"`
	ImgUrl       string             `json:"img_url,omitempty"`
	SoundUrl     string             `json:"sound_url,omitempty"`
}

type PokemonSmallInfo struct {
	NationalID int    `db:"national_id" json:"national_id"`
	Name       string `db:"name" json:"name"`
	ImgUrl     string `json:"img_url,omitempty"`
}

func GetSmallInfo(id int) (*PokemonSmallInfo, error) {
	var pokemon_info PokemonSmallInfo

	data_err := db.Get(&pokemon_info, "SELECT national_id, name FROM pokemons WHERE national_id = ?", id)
	if data_err != nil {
		// Failed to get ID info: error msg
		log.Fatalf("Failed to get %d DB small info: %s", id, data_err)
		return nil, data_err
	}

	return &pokemon_info, nil
}

func getInfoById(id string) (*PokemonInfo, error) {
	var pokemon_info PokemonInfo

	data_err := db.Get(&pokemon_info, "SELECT * FROM pokemons WHERE national_id = ?", id)
	if data_err != nil {
		log.Fatalf("Failed to get %s DB info: %s", id, data_err)
	}
	return &pokemon_info, nil
}

func getInfoByName(name string) (*PokemonInfo, error) {
	var pokemon_info PokemonInfo

	data_err := db.Get(&pokemon_info, "SELECT * FROM pokemons WHERE name = ?", name)
	if data_err != nil {
		log.Fatalf("Failed to get %s DB info: %s", name, data_err)
	}
	return &pokemon_info, nil
}

func GetEvolutions(id string) ([]PokemonSmallInfo, error) {
	var evolutions []PokemonSmallInfo

	rows, evo_err := db.Query("SELECT evolution_id FROM pokemon_evolutions WHERE pokemon_id = ?", id)

	if evo_err != nil {
		log.Fatalf("Failed to get %s evolutions: %s", id, evo_err)
	}

	for rows.Next() {
		var evolutionID int
		if err := rows.Scan(&evolutionID); err != nil {
			log.Fatalf("Failed to scan evolution id: %s", err)
		}

		small_info, err := GetSmallInfo(evolutionID)
		if err != nil {
			log.Fatalf("Failed to get small info for evolution id: %s", err)
		}

		var name_err error
		small_info.ImgUrl, name_err = buildImgUrl(small_info.Name)

		if name_err != nil {
		}

		evolutions = append(evolutions, *small_info)
	}

	return evolutions, nil
}
