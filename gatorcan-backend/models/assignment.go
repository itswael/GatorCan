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
	ID                  uint                 `gorm:"primaryKey" json:"id"`
	AssignmentID        uint                 `gorm:"not null" json:"assignment_id"`
	Assignment          Assignment           `gorm:"foreignKey:AssignmentID"`
	FileName            string               `gorm:"not null" json:"file_name"`
	FileURL             string               `gorm:"not null" json:"file_url"`
	FileType            string               `gorm:"not null" json:"file_type"`
	CreatedAt           time.Time            `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt           time.Time            `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	UserAssignmentFiles []UserAssignmentFile `gorm:"foreignKey:AssignmentFileID" json:"user_assignment_files"`
}

type UserAssignmentFile struct {
	gorm.Model
	ID               uint           `gorm:"primaryKey" json:"id"`
	UserID           uint           `gorm:"not null;uniqueIndex:idx_user_assignment_file" json:"user_id"`
	User             User           `gorm:"foreignKey:UserID"`
	AssignmentFileID uint           `gorm:"not null;uniqueIndex:idx_user_assignment_file" json:"assignment_file_id"`
	AssignmentFile   AssignmentFile `gorm:"foreignKey:AssignmentFileID"`
}
