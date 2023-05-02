package controller

import (
	"api-go/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "Freaks", Artist: "Surf Course", Price: 36.99},
	{ID: "2", Title: "Notion", Artist: "The Rare Occasions", Price: 27.99},
	{ID: "3", Title: "Island In The Sun", Artist: "Weezer", Price: 29.99},
	{ID: "4", Title: "The Adults Are Talking", Artist: "The Strokes", Price: 20.99},
	{ID: "5", Title: "Welcome To The Jungle", Artist: "Guns Roses", Price: 40.99},
	{ID: "6", Title: "Californication", Artist: "Red Hot Chilli Peppers", Price: 42.99},
	{ID: "7", Title: "Creep", Artist: "RadioHead", Price: 38.99},
	{ID: "8", Title: "Dark Necessities", Artist: "Red Hot Chilli Peppers", Price: 39.99},
	{ID: "9", Title: "Dream On", Artist: "Aerosmith", Price: 37.99},
	{ID: "10", Title: "Seven Nation Army", Artist: "The White Stripes", Price: 32.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumsByID(c *gin.Context) {

	id := c.Param("id")

	for _, x := range albums {
		if x.ID == id {
			c.IndentedJSON(http.StatusOK, x)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func PostAlbums(c *gin.Context) {

	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	jsonFile, err := os.Open("input.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &newAlbum)

	//add lines to the album slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
