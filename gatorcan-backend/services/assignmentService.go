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

func NewAssignmentService(assignmentRepo interfaces.AssignmentRepository, config *config.AppConfig, httpClient interfaces.HTTPClient) interfaces.AssignmentService {
	return &AssignmentService{assignmentRepo: assignmentRepo, config: config, httpClient: httpClient}
}

func (s *AssignmentService) GetAssignmentsByCourseID(ctx context.Context, courseID int) ([]dtos.AssignmentResponseDTO, error) {
	panic("not implemented")
}
func (s *AssignmentService) GetAssignmentByIDAndCourseID(ctx context.Context, assignmentID int, courseID int) (dtos.AssignmentResponseDTO, error) {
	panic("not implemented")
}

func (s *AssignmentService) UploadFileToAssignment(ctx context.Context, logger *log.Logger, username string, uploadData *dtos.UploadFileToAssignmentDTO) (*dtos.UploadFileToAssignmentResponseDTO, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Printf("user not found: %s %d", username, 404)
		return nil, errors.ErrUserNotFound
	}

}
