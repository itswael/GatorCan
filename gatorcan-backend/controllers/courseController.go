package controllers

import (
	"gatorcan-backend/database"
	"gatorcan-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCourse allows an admin to create a new course
func CreateCourse(c *gin.Context) {
	var request struct {
		CourseName     string `json:"course_name" binding:"required"`
		CourseID       string `json:"course_id" binding:"required"`
		AvailableSeats int    `json:"available_seats" binding:"required"`
		TotalSeats     int    `json:"total_seats" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if user has admin role
	roles, _ := c.Get("roles")
	userRoles := roles.([]string)
	isAdmin := false
	for _, role := range userRoles {
		if role == "admin" {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can create courses"})
		return
	}

	// Create the course
	course := models.Course{
		CourseName:     request.CourseName,
		CourseID:       request.CourseID,
		AvailableSeats: request.AvailableSeats,
		TotalSeats:     request.TotalSeats,
	}

	if err := database.DB.Create(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course created successfully", "course": course})
}
