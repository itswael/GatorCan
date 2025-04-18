package services

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/adapters"
	"gatorcan-backend/config"
	"gatorcan-backend/errors"
	"gatorcan-backend/interfaces"
	"log"
)

type AiServiceServiceImpl struct {
	courseRepo interfaces.CourseRepository
	userRepo   interfaces.UserRepository
	config     *config.AppConfig
	httpClient interfaces.HTTPClient
}

func NewAIServiceService(courseRepo interfaces.CourseRepository, userRepo interfaces.UserRepository, config *config.AppConfig, httpClient interfaces.HTTPClient) interfaces.AiServiceService {
	return &AiServiceServiceImpl{courseRepo: courseRepo, userRepo: userRepo, config: config, httpClient: httpClient}
}

func (s *AiServiceServiceImpl) GetCourseRecommendations(ctx context.Context, logger *log.Logger, username string) ([]dtos.CourseRecommendationResponseDTO, error) {
	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Printf("user not found: %s %d", username, 404)
		return nil, errors.ErrUserNotFound
	}

	enrolled_courses, err := s.courseRepo.GetEnrolledCourses(ctx, int(user.ID))
	if err != nil {
		logger.Printf("failed to fetch enrolled courses: %s %d", username, 500)
		return nil, errors.ErrFailedToFetch
	}

	var enrolledIDs []int
	for _, enrollment := range enrolled_courses {
		enrolledIDs = append(enrolledIDs, int(enrollment.ActiveCourseID))
	}

	var keywords []string
	// add keywords from search criteria if any
	keywords = append(keywords, "db") // Example keywords

	// Call the AI service to get course recommendations
	courseRecommendations, err := adapters.GetRecommendedCourses(enrolledIDs, keywords, logger)
	if err != nil {
		logger.Printf("failed to call AI service: %v", err)
		return nil, errors.ErrMicroserviceError
	}

	return courseRecommendations, nil
}
