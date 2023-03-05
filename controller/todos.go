package controller

import (
	"fmt"
	"learn/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type TodosController struct {
	db *gorm.DB
}

func NewTodosController(db *gorm.DB) *TodosController {
	return &TodosController{db}
}

func (todocontroller *TodosController) GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	if err := todocontroller.db.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No records found"})
		return
	}
	c.JSON(200, todos)
}

func (todocontroller *TodosController) CreateTodo(c *gin.Context) {
	// Get the user ID from the jwt token
	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Verify that the signing method is correct
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method: %v", t.Header["alg"])
		}

		// Get the secret key from environtment variable or configuration file
		secretKey := []byte("secretkey")

		return secretKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Type assertion to get user ID as float64
	userID := claims["user_id"].(float64)

	// Get the request body and create a new Todo object
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "the data is invalid"})
		return
	}

	// Set the user ID for the new todo
	newTodo.UserID = uint(userID)

	// Insert the new todo into the database
	todocontroller.db.Create(&newTodo)

	// Return the new todo in the response
	c.JSON(http.StatusCreated, newTodo)
}
