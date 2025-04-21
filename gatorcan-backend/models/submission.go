package models

import (
	"time"

	"gorm.io/gorm"
)

// created_at can be used as submission timestamp
// updated_at can be used as latest submission timestamp

type Submission struct {
	gorm.Model
	ID           uint       `gorm:"primaryKey" json:"id"`
	AssignmentID uint       `json:"assignment_id"`
	CourseID     uint       `json:"course_id"`
	UserID       uint       `json:"user_id"`
	File_url     string     `json:"file_url"`
	File_name    string     `json:"file_name"`
	File_type    string     `json:"file_type"`
	Grade        int        `json:"grade" gorm:"default:NULL"`
	Feedback     string     `json:"feedback"`
	Created_at   time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Updated_at   time.Time  `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	Deleted_at   time.Time  `json:"deleted_at" gorm:"default:NULL"`
	Assignment   Assignment `gorm:"foreignKey:AssignmentID"`
	User         User       `gorm:"foreignKey:UserID"`
	Course       Course     `gorm:"foreignKey:CourseID"`
}
