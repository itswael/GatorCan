package models

import (
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	ID             uint         `gorm:"primaryKey" json:"id"`
	Title          string       `gorm:"not null" json:"title"`
	Description    string       `json:"description"`
	Deadline       time.Time    `json:"deadline"`
	CreatedAt      time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	ActiveCourseID uint         `json:"course_id"`
	ActiveCourse   ActiveCourse `gorm:"foreignKey:ActiveCourseID"`
	MaxPoints      int          `json:"max_points"`
}
