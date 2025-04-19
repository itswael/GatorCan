package services

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/config"
	"gatorcan-backend/errors"
	"gatorcan-backend/interfaces"
	"gatorcan-backend/models"
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
	assignments, err := s.assignmentRepo.GetAssignmentsByCourseID(ctx, courseID)
	if err != nil {
		return nil, errors.ErrAssignmentNotFound
	}
	assignmentsResponse := make([]dtos.AssignmentResponseDTO, len(assignments))
	for i, assignment := range assignments {
		assignmentsResponse[i] = dtos.AssignmentResponseDTO{
			ID:             assignment.ID,
			Title:          assignment.Title,
			Description:    assignment.Description,
			Deadline:       assignment.Deadline,
			ActiveCourseID: assignment.ActiveCourseID,
			MaxPoints:      assignment.MaxPoints,
		}
	}
	return assignmentsResponse, nil
}
func (s *AssignmentService) GetAssignmentByIDAndCourseID(ctx context.Context, assignmentID int, courseID int) (dtos.AssignmentResponseDTO, error) {
	assignment, err := s.assignmentRepo.GetAssignmentByIDAndCourseID(ctx, assignmentID, courseID)
	if err != nil {
		return dtos.AssignmentResponseDTO{}, errors.ErrAssignmentNotFound
	}

	_, err = s.courseRepo.GetCourseByID(ctx, courseID)
	if err != nil {
		return dtos.AssignmentResponseDTO{}, errors.ErrCourseNotFound
	}

	assignmentResponse := dtos.AssignmentResponseDTO{
		ID:             assignment.ID,
		Title:          assignment.Title,
		Description:    assignment.Description,
		Deadline:       assignment.Deadline,
		ActiveCourseID: assignment.ActiveCourseID,
		MaxPoints:      assignment.MaxPoints,
	}

	return assignmentResponse, nil
}

func (s *AssignmentService) UploadFileToAssignment(ctx context.Context, logger *log.Logger, username string, uploadData *dtos.UploadFileToAssignmentDTO) (*dtos.UploadFileToAssignmentResponseDTO, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
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

	_, err = s.assignmentRepo.GetAssignmentByIDAndCourseID(ctx, int(uploadData.AssignmentID), int(uploadData.CourseID))
	if err != nil {
		logger.Printf("assignment not found: %d %d", uploadData.AssignmentID, 404)
		return nil, errors.ErrAssignmentNotFound
	}

	// Call the repository to handle the database logic
	assignmentFile := models.AssignmentFile{
		AssignmentID: uploadData.AssignmentID,
		FileName:     uploadData.FileName,
		FileURL:      uploadData.FileURL,
		FileType:     uploadData.FileType,
	}

	if err := s.assignmentRepo.CreateAssignmentFile(ctx, &assignmentFile); err != nil {
		logger.Printf("Error uploading file to assignment: %v", err)
		return nil, errors.ErrFailedToUploadFile
	}

	// Link user to the uploaded file
	userAssignmentFile := models.UserAssignmentFile{
		UserID:           user.ID,
		AssignmentFileID: assignmentFile.ID,
	}

	if err := s.assignmentRepo.LinkUserToAssignmentFile(ctx, &userAssignmentFile); err != nil {
		logger.Printf("Error linking file to user: %v", err)
		return nil, errors.ErrFailedToLinkFileToUser
	}

	// Convert to response DTO
	response := dtos.NewUploadFileToAssignmentResponseDTO(&assignmentFile, user.ID, uploadData.CourseID)

	return response, nil
}

func (s *AssignmentService) UpsertAssignment(ctx context.Context, logger *log.Logger, assignment *dtos.CreateAssignmentRequestDTO) (dtos.AssignmentResponseDTO, error) {
	// Check if the course exists
	_, err := s.courseRepo.GetCourseByID(ctx, int(assignment.CourseID))
	if err != nil {
		logger.Printf("course not found: %d %d", assignment.CourseID, 404)
		return dtos.AssignmentResponseDTO{}, errors.ErrCourseNotFound
	}

	// Create the assignment model
	assignmentModel := models.Assignment{
		Title:          assignment.Title,
		Description:    assignment.Description,
		Deadline:       assignment.Deadline,
		ActiveCourseID: assignment.CourseID,
		MaxPoints:      assignment.MaxPoints,
	}

	// Call the repository to create the assignment
	if err := s.assignmentRepo.UpsertAssignment(ctx, &assignmentModel); err != nil {
		logger.Printf("failed to create assignment: %v", err)
		return dtos.AssignmentResponseDTO{}, errors.ErrFailedToCreateAssignment
	}

	// Convert to response DTO
	assignmentResponse := dtos.AssignmentResponseDTO{
		ID:             assignmentModel.ID,
		Title:          assignmentModel.Title,
		Description:    assignmentModel.Description,
		Deadline:       assignmentModel.Deadline,
		ActiveCourseID: assignmentModel.ActiveCourseID,
		MaxPoints:      assignmentModel.MaxPoints,
	}

	return assignmentResponse, nil
}
