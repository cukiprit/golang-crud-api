package controller

import (
	"learn/api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUsersController(db *gorm.DB) *UserController {
	return &UserController{db}
}

// Register creates a new user in the database
func (usercontroller *UserController) Register(c *gin.Context) {
	// Get the user data from the request body
	var users models.User
	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatus(400)
		return
	}

	// Hash the user's password before storing in the database
	if err := users.HashPassword(); err != nil {
		c.AbortWithStatus(500)
		return
	}

	// Create a new user in the database
	if err := usercontroller.db.Create(&users).Error; err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}

// Login authenticates a user
func (usercontroller *UserController) Login(c *gin.Context) {
	// Get the user data from the request body
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.AbortWithStatus(400)
		return
	}

	// check if the user exists in the database
	var user models.User
	if err := usercontroller.db.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		c.AbortWithStatus(401)
		return
	}

	// Compare the password hash with the user's input
	if !user.CheckPasswordHash(loginData.Password) {
		c.AbortWithStatus(401)
		return
	}

	c.JSON(200, gin.H{
		"message": "User authenticated successfully",
		"user_id": user.ID,
	})
}
