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
	assignmentRepo interfaces.AssignmentRepository
	config         *config.AppConfig
	httpClient     interfaces.HTTPClient
}

func NewSubmissionService(submissionRepo interfaces.SubmissionRepository, assignmentRepo interfaces.AssignmentRepository, userRepo interfaces.UserRepository, courseRepo interfaces.CourseRepository, config *config.AppConfig, httpClient interfaces.HTTPClient) interfaces.SubmissionService {
	return &SubmissionServiceImpl{submissionRepo: submissionRepo, assignmentRepo: assignmentRepo, userRepo: userRepo, courseRepo: courseRepo, config: config, httpClient: httpClient}
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

func (s *SubmissionServiceImpl) GetSubmission(ctx context.Context, courseID int, assignmentID int, userID uint) (*dtos.SubmissionResponseDTO, error) {
	submission, err := s.submissionRepo.GetSubmission(ctx, courseID, assignmentID, userID)
	if err != nil {
		return &dtos.SubmissionResponseDTO{}, errors.ErrSubmissionNotFound
	}

	//find max point for assignment
	assignment, err := s.assignmentRepo.GetAssignmentByIDAndCourseID(ctx, assignmentID, courseID)
	if err != nil {
		return &dtos.SubmissionResponseDTO{}, errors.ErrAssignmentNotFound
	}

	response := dtos.SubmissionResponseDTO{
		FileURL:     submission.File_url,
		Grade:       submission.Grade,
		Feedback:    submission.Feedback,
		SubmittedAt: submission.Updated_at.Format("2006-01-02 15:04:05"),
		MaxPoints:   assignment.MaxPoints,
	}
	return &response, nil
}
