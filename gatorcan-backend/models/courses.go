package models

import (
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseName     string `gorm:"not null"`        // Name of the course (e.g., "Data Science")
	CourseID       string `gorm:"unique;not null"` // Unique identifier for the course (e.g., "CS101")
	AvailableSeats int    `gorm:"not null"`        // Seats available for enrollment
	TotalSeats     int    `gorm:"not null"`        // Total seats in the course
}
