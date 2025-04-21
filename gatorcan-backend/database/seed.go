package database

import (
	"gatorcan-backend/models"
	"log"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	users := []models.User{
		{
			Username: "admin",
			Email:    "admin@gmail.com",
			Password: "$2a$10$qBNeqBt9zETe3idk5UAxWeBa3e36Z5uIQ5Xq8uQlQ6LKVV7TXNK5S", // You should hash this in real-world cases
		},
		{
			Username: "instructor",
			Email:    "instructor@gmail.com",
			Password: "$2a$10$iZxEoLP4fqM7DD6kelGyZu9TjyxTzvWwymtbd3TW2Ko7DXkI92dTG",
		},
	}

	for _, user := range users {
		var existing models.User
		result := db.Where("email = ?", user.Email).First(&existing)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				if err := db.Create(&user).Error; err != nil {
					log.Printf("Failed to create user %s: %v", user.Email, err)
				} else {
					log.Printf("Created user: %s", user.Email)
				}
			} else {
				return result.Error
			}
		}
	}

	return nil
}
