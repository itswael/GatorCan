package services

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/config"
	"gatorcan-backend/errors"
	"gatorcan-backend/interfaces"
	"log"
)

type AssignmentService struct {
	assignmentRepo interfaces.AssignmentRepository
	userRepo       interfaces.UserRepository
	courseRepo     interfaces.CourseRepository
	config         *config.AppConfig
	httpClient     interfaces.HTTPClient
}

func NewAssignmentService(assignmentRepo interfaces.AssignmentRepository, userRepo interfaces.UserRepository, courseRepo interfaces.CourseRepository, config *config.AppConfig, httpClient interfaces.HTTPClient) interfaces.AssignmentService {
	return &AssignmentService{assignmentRepo: assignmentRepo, userRepo: userRepo, courseRepo: courseRepo, config: config, httpClient: httpClient}
}

func (s *AssignmentService) GetAssignmentsByCourseID(ctx context.Context, courseID int) ([]dtos.AssignmentResponseDTO, error) {
	panic("not implemented")
}
func (s *AssignmentService) GetAssignmentByIDAndCourseID(ctx context.Context, assignmentID int, courseID int) (dtos.AssignmentResponseDTO, error) {
	panic("not implemented")
}

func (s *AssignmentService) UploadFileToAssignment(ctx context.Context, logger *log.Logger, username string, uploadData *dtos.UploadFileToAssignmentDTO) (*dtos.UploadFileToAssignmentResponseDTO, error) {
	_, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Printf("user not found: %s %d", username, 404)
		return nil, errors.ErrUserNotFound
	}
	_, err = s.courseRepo.GetCourseByID(ctx, int(uploadData.CourseID))
	if err != nil {
		logger.Printf("course not found: %d %d", uploadData.CourseID, 404)
		return nil, errors.ErrCourseNotFound
	}

	//To do: Implement "GetAssignmentByIDAndCourseID" method in the assignment repository.

	// _, err = s.assignmentRepo.GetAssignmentByIDAndCourseID(ctx, int(uploadData.AssignmentID), int(uploadData.CourseID))
	// if err != nil {
	// 	logger.Printf("assignment not found: %d %d", uploadData.AssignmentID, 404)
	// 	return nil, errors.ErrAssignmentNotFound
	// }

	// Call the repository to handle the database logic
	uploadResponse, err := s.assignmentRepo.UploadFileToAssignment(ctx, logger, username, &dtos.UploadFileToAssignmentDTO{
		AssignmentID: uploadData.AssignmentID,
		CourseID:     uploadData.CourseID,
		FileName:     uploadData.FileName,
		FileURL:      uploadData.FileURL,
		FileType:     uploadData.FileType,
	})
	if err != nil {
		logger.Printf("error uploading file to assignment: %v", err)
		return nil, err
	}

	return uploadResponse, nil
}
