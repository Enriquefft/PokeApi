package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // Driver for sqlite3
)

var db *sqlx.DB

func InitDB() error {
	var err error
	db, err = sqlx.Open("sqlite3", "Test.db")
	if err != nil {
		log.Fatal("Couldn't connect to database", err)
		return err
	}
	log.Println("Connected to database")
	return nil
}
