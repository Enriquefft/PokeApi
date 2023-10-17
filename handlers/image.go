package image_route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetImage(gin_ctx *gin.Context) {
	name := gin_ctx.Params.ByName("name")
	id := gin_ctx.Params.ByName("id")

	// If name is present, then use name
	if name != "" {
		gin_ctx.JSON(http.StatusOK, "name")
	} else if id != "" {
		// If id is present, then use id
		gin_ctx.JSON(http.StatusOK, "id")
	} else {
		// Else, return error
		gin_ctx.JSON(http.StatusNotFound, gin.H{"status": "404", "message": "No name or id provided"})
	}
}
