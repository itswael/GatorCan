package repositories

import (
	"gatorcan-backend/models"

	"gorm.io/gorm"
)

type AssignmentRepository interface {
	GetAssignmentsByCourseID(courseID int) ([]models.Assignment, error)
	GetAssignmentByIDAndCourseID(assignmentID int, courseID int) (models.Assignment, error)
}

type assignmentRepository struct {
	db *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) AssignmentRepository {
	return &assignmentRepository{db: db}
}

// GetAssignmentByIDAndCourseID implements AssignmentRepository.
func (a *assignmentRepository) GetAssignmentByIDAndCourseID(assignmentID int, courseID int) (models.Assignment, error) {
	panic("unimplemented")
}

// GetAssignmentsByCourseID implements AssignmentRepository.
func (a *assignmentRepository) GetAssignmentsByCourseID(courseID int) ([]models.Assignment, error) {
	panic("unimplemented")
}
