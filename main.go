package main

import (
	"learn/api/config"
	"learn/api/controller"
	"learn/api/middleware"
	"learn/api/seeder"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Print("Failed to connect database")
	}

	if err := seeder.Seed(db); err != nil {
		log.Print("Failed to seed database")
	}

	router := gin.Default()
	tc := controller.NewTodosController(db)
	uc := controller.NewUsersController(db)

	router.GET("/", tc.GetAllTodos)
	router.POST("/todos", middleware.AuthMiddleware(), tc.CreateTodo)
	router.POST("/login", uc.Login)
	router.POST("/register", uc.Register)

	router.Run("localhost:8080")
}
