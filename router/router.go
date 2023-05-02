package router

import (
	"api-go/controller"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()
	router.GET("/albums", controller.GetAlbums)
}