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

type AssignmentFile struct {
	gorm.Model
	ID             uint         `gorm:"primaryKey" json:"id"`
	AssignmentID   uint         `json:"assignment_id"`
	Assignment     Assignment   `gorm:"foreignKey:AssignmentID"`
	FileName       string       `gorm:"not null" json:"file_name"`
	FileURL        string       `gorm:"not null" json:"file_url"`
	FileType       string       `gorm:"not null" json:"file_type"`
	CreatedAt      time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	UploadedBy     string       `gorm:"not null" json:"uploaded_by"`
	ActiveCourseID uint         `json:"course_id"`
	ActiveCourse   ActiveCourse `gorm:"foreignKey:ActiveCourseID"`
}
