package controllers

import (
	"net/http"

	"github.com/akmalcm/team-soccer-go/configs"
	"github.com/akmalcm/team-soccer-go/models"
	"github.com/gin-gonic/gin"
)

type CreatePlayerInput struct {
	Forename string `json:"forename" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	ImageURL string `json:"image_url" binding:"required"`
}

type UpdatePlayerInput struct {
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
	ImageURL string `json:"image_url"`
}

// GET /players
// Find all players
func FindPlayers(c *gin.Context) {
	var players []models.Player
	configs.DB.Scopes(configs.Paginate(c)).Find(&players)

	c.JSON(http.StatusOK, gin.H{"data": players})
}

// GET /players/:id
// Find a player
func FindPlayer(c *gin.Context) {
	// Get model if exist
	var player models.Player
	if err := configs.DB.Where("id = ?", c.Query("id")).First(&player).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": player})
}

// POST /players
// Create new player
func CreatePlayer(c *gin.Context) {
	// Validate input
	var input CreatePlayerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create player
	player := models.Player{Forename: input.Forename, Surname: input.Surname}
	configs.DB.Create(&player)

	c.JSON(http.StatusOK, gin.H{"data": player})
}

// PATCH /players/:id
// Update a player
func UpdatePlayer(c *gin.Context) {
	// Get model if exist
	var player models.Player
	if err := configs.DB.Where("id = ?", c.Param("id")).First(&player).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdatePlayerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	configs.DB.Model(&player).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": player})
}

// DELETE /players/:id
// Delete a player
func DeletePlayer(c *gin.Context) {
	// Get model if exist
	var player models.Player
	if err := configs.DB.Where("id = ?", c.Param("id")).First(&player).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	configs.DB.Delete(&player)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
