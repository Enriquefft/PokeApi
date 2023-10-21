package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	db_err := InitDB()
	if db_err != nil {
		panic(db_err)
	}
	router := gin.Default()
	router.GET("/info/id", GetAllPokemonInfoById)

	return router
}

func main() {
	r := setupRouter()

	// Listen and Server in Utec network port 3001
	// r.Run("10.100.226.35:3001")
	r.Run(":3001")
}
