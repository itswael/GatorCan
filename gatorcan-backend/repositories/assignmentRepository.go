package repositories

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/errors"
	"gatorcan-backend/models"
	"log"

	"gorm.io/gorm"
)

type AssignmentRepository interface {
	GetAssignmentsByCourseID(courseID int) ([]models.Assignment, error)
	GetAssignmentByIDAndCourseID(ctx context.Context, assignmentID int, courseID int) (models.Assignment, error)
	UploadFileToAssignment(ctx context.Context, logger *log.Logger, username string, uploadData *dtos.UploadFileToAssignmentDTO) (*dtos.UploadFileToAssignmentResponseDTO, error)
}

type assignmentRepository struct {
	db *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) AssignmentRepository {
	return &assignmentRepository{db: db}
}

// GetAssignmentByIDAndCourseID implements AssignmentRepository.
func (a *assignmentRepository) GetAssignmentByIDAndCourseID(ctx context.Context, assignmentID int, courseID int) (models.Assignment, error) {
	panic("unimplemented")
}

// GetAssignmentsByCourseID implements AssignmentRepository.
func (a *assignmentRepository) GetAssignmentsByCourseID(courseID int) ([]models.Assignment, error) {
	panic("unimplemented")
}

func (a *assignmentRepository) UploadFileToAssignment(ctx context.Context, logger *log.Logger, username string, uploadData *dtos.UploadFileToAssignmentDTO) (*dtos.UploadFileToAssignmentResponseDTO, error) {
	// Create the assignment file record.
	assignmentFile := models.AssignmentFile{
		AssignmentID: uploadData.AssignmentID,
		FileName:     uploadData.FileName,
		FileURL:      uploadData.FileURL,
		FileType:     uploadData.FileType,
	}

	// To do: Implement a check to see if the assignment exists. If it does, then overwrite the assignment file.

	if err := a.db.WithContext(ctx).Create(&assignmentFile).Error; err != nil {
		logger.Printf("Error uploading file to assignment: %v", err)
		return nil, errors.ErrFailedToUploadFile
	}

	// Retrieve the user record. Adjust the field (e.g., "username") if needed.
	var user models.User
	if err := a.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		logger.Printf("Error fetching user with username %s: %v", username, err)
		return nil, errors.ErrUserNotFound
	}

	// Create the join record linking the user with the assignment file.
	userAssignmentFile := models.UserAssignmentFile{
		UserID:           user.ID,
		AssignmentFileID: assignmentFile.ID,
	}
	if err := a.db.WithContext(ctx).Create(&userAssignmentFile).Error; err != nil {
		logger.Printf("Error linking file to user: %v", err)
		return nil, errors.ErrFailedToLinkFileToUser
	}

	// Prepare the response DTO.
	response := &dtos.UploadFileToAssignmentResponseDTO{
		AssignmentID: assignmentFile.AssignmentID,
		FileName:     assignmentFile.FileName,
		FileURL:      assignmentFile.FileURL,
		FileType:     assignmentFile.FileType,
		UploadedAt:   assignmentFile.CreatedAt,
		UploaderID:   user.ID,
		CourseID:     uploadData.CourseID,
	}

	return response, nil
}
