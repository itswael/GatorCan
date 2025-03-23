package services

import (
	"context"
	"gatorcan-backend/config"
	"gatorcan-backend/interfaces"
)

type AssignmentService struct {
	assignmentRepo interfaces.AssignmentRepository
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
