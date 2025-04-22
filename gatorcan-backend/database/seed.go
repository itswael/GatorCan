package database

import (
	"gatorcan-backend/models"
	"log"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	// Ensure roles exist
	roleNames := []string{"admin", "instructor", "student"}
	rolesMap := make(map[string]*models.Role)

	for _, roleName := range roleNames {
		var role models.Role
		if err := db.Where("name = ?", roleName).First(&role).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				role = models.Role{Name: roleName}
				if err := db.Create(&role).Error; err != nil {
					log.Printf("Failed to create role %s: %v", roleName, err)
					continue
				}
				log.Printf("Created role: %s", roleName)
			} else {
				return err
			}
		}
		rolesMap[roleName] = &role
	}

	// Users with their corresponding role names
	users := []struct {
		User models.User
		Role string
	}{
		{
			User: models.User{
				Username: "admin",
				Email:    "admin@gmail.com",
				Password: "$2a$10$qBNeqBt9zETe3idk5UAxWeBa3e36Z5uIQ5Xq8uQlQ6LKVV7TXNK5S", // bcrypt hash of "admin"
			},
			Role: "admin",
		},
		{
			User: models.User{
				Username: "instructor",
				Email:    "instructor@gmail.com",
				Password: "$2a$10$iZxEoLP4fqM7DD6kelGyZu9TjyxTzvWwymtbd3TW2Ko7DXkI92dTG", // bcrypt hash of "instructor"
			},
			Role: "instructor",
		},
		{
			User: models.User{
				Username: "student",
				Email:    "student@gmail.com",
				Password: "$2a$10$e6uRj7lKz1DA.lA6ZzMbXeA2k6v1vZyV3FgAedAiQK/lsmMDmX5l6", // bcrypt hash of "student"
			},
			Role: "student",
		},
	}

	for _, u := range users {
		var existing models.User
		result := db.Preload("Roles").Where("email = ?", u.User.Email).First(&existing)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				u.User.Roles = []*models.Role{rolesMap[u.Role]}
				if err := db.Create(&u.User).Error; err != nil {
					log.Printf("Failed to create user %s: %v", u.User.Email, err)
				} else {
					log.Printf("Created user: %s with role: %s", u.User.Email, u.Role)
				}
			} else {
				return result.Error
			}
		} else {
			// User exists but maybe role is missing
			if !hasRole(existing.Roles, u.Role) {
				if err := db.Model(&existing).Association("Roles").Append(rolesMap[u.Role]); err != nil {
					log.Printf("Failed to assign role %s to user %s: %v", u.Role, u.User.Email, err)
				} else {
					log.Printf("Assigned role %s to existing user %s", u.Role, u.User.Email)
				}
			}
		}
	}

	return nil
}

func hasRole(roles []*models.Role, name string) bool {
	for _, r := range roles {
		if r.Name == name {
			return true
		}
	}
	return false
}
