package controller

import (
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
		c.AbortWithStatus(500)
		return
	}
	c.JSON(200, todos)
}

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

func (todocontroller *TodosController) CreateTodo(c *gin.Context) {
	var request CreateTodoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secretkey"), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		return
	}

	userID := claims["sub"].(float64)

	todo := models.Todo{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		UserID:      uint(userID),
	}

	if err := todocontroller.db.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}
