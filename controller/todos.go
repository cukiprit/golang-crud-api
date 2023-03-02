package controller

import (
	"learn/api/models"

	"github.com/gin-gonic/gin"
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
