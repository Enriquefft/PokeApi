package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func buildImageURL(name string) string {
	var imageURLBuilder strings.Builder
	imageURLBuilder.WriteString("https://github.com/Edgar5377/Pokedex/blob/main/Pokemon%20Dataset/")
	imageURLBuilder.WriteString(name)
	imageURLBuilder.WriteString(".png?raw=true")
	return imageURLBuilder.String()
}

func GetImage(gin_ctx *gin.Context) {
	name := gin_ctx.Params.ByName("name")
	id := gin_ctx.Params.ByName("id")

	// If name is present, then use name
	if name != "" {
		imageURL := buildImageURL(name)
		gin_ctx.JSON(http.StatusOK, imageURL)

	} else if id != "" {
		// Else if id is present, then get the name through the DB

		gin_ctx.JSON(http.StatusOK, "id")
	} else {
		// Else, return error
		gin_ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "No name or id provided"})
	}
}
