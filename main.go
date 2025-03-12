package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shivanshi1/golang-rest-api/database"
)

// Load environment variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
	fmt.Println("✅ Environment variables loaded")
}

func main() {
	// Load environment variables
	LoadEnv()

	// Initialize database
	database.ConnectDatabase()

	// Set up Gin router
	router := gin.Default()

	// API Routes
	router.POST("/users", CreateUser)
	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)
	router.PUT("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}

// Create User
func CreateUser(c *gin.Context) {
	var user database.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// Get All Users
func GetUsers(c *gin.Context) {
	users := database.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

// Get Single User
func GetUser(c *gin.Context) {
	var user database.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// Update User
func UpdateUser(c *gin.Context) {
	var user database.User
	id := c.Param("id")

	// Check if user exists
	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updatedUser database.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the primary key (ID) does not change
	updatedUser.ID = user.ID

	database.DB.Save(&updatedUser)
	c.JSON(http.StatusOK, updatedUser)
}

// Delete User
func DeleteUser(c *gin.Context) {
	var user database.User
	if err := database.DB.First(&user, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
