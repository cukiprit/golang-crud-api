package seeder

import (
	"learn/api/models"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	// Initialize a users variable based on User model
	var users []models.User
	db.Find(&users)
	// Check if the User model have 0 data
	if len(users) < 3 {
		// Generate 2 user data
		for i := 0; i < 2; i++ {
			user := models.User{
				Username: gofakeit.Username(),
				Password: gofakeit.Password(true, true, true, false, false, 10),
			}
			if err := db.Create(&user).Error; err != nil {
				return err
			}
			// Initialize a todos variable based on Todo model
			var todos []models.Todo
			db.Find(&todos)
			// Check if the Todo model have 0 data
			if len(todos) < 4 {
				// Generate 3 todo data
				for j := 0; j < 3; j++ {
					todo := models.Todo{
						Title:       gofakeit.BS(),
						Description: gofakeit.Sentence(10),
						Status:      gofakeit.RandomString([]string{"Finished", "On-Going", "Unfinished"}),
						UserID:      user.ID,
					}
					if err := db.Create(&todo).Error; err != nil {
						return err
					}
				}
			} else {
				log.Printf("Found %d todos\n", len(todos))
			}
		}
	} else {
		log.Printf("Found %d users\n", len(users))
	}
	return nil
}
