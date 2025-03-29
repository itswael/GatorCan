package services

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/config"
	"gatorcan-backend/errors"
	"gatorcan-backend/interfaces"
	"log"
)

type SubmissionServiceImpl struct {
	submissionRepo interfaces.SubmissionRepository
	userRepo       interfaces.UserRepository
	courseRepo     interfaces.CourseRepository
	config         *config.AppConfig
	httpClient     interfaces.HTTPClient
}

func NewSubmissionService(submissionRepo interfaces.SubmissionRepository, userRepo interfaces.UserRepository, courseRepo interfaces.CourseRepository, config *config.AppConfig, httpClient interfaces.HTTPClient) interfaces.SubmissionService {
	return &SubmissionServiceImpl{submissionRepo: submissionRepo, userRepo: userRepo, courseRepo: courseRepo, config: config, httpClient: httpClient}
}

func (s *SubmissionServiceImpl) GradeSubmission(ctx context.Context, logger *log.Logger, username string, submissionData *dtos.GradeSubmissionRequestDTO) (*dtos.GradeSubmissionResponseDTO, error) {
	// Validate user existence
	_, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Printf("User not found: %s", username)
		return nil, errors.ErrUserNotFound
	}

	// Validate course existence
	_, err = s.courseRepo.GetCourseByID(ctx, int(submissionData.CourseID))
	if err != nil {
		logger.Printf("Course not found: %d", submissionData.CourseID)
		return nil, errors.ErrCourseNotFound
	}

	// Grade the submission
	if err := s.submissionRepo.GradeSubmission(ctx, submissionData.AssignmentID, submissionData.CourseID, submissionData.UserID, float64(submissionData.Grade), submissionData.Feedback); err != nil {
		logger.Printf("Error grading submission: %v", err)
		return nil, errors.ErrGradingSubmissionFailed
	}

	// Create the response
	response := &dtos.GradeSubmissionResponseDTO{
		AssignmentID: submissionData.AssignmentID,
		CourseID:     submissionData.CourseID,
		UserID:       submissionData.UserID,
		Grade:        submissionData.Grade,
		Feedback:     submissionData.Feedback,
	}

	return response, nil
}
