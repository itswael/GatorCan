package mocks

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"log"

	"github.com/stretchr/testify/mock"
)

// MockCourseService mocks the CourseService interface
type MockCourseService struct {
	mock.Mock
}

func (m *MockCourseService) GetEnrolledCourses(ctx context.Context, logger *log.Logger, username string) ([]dtos.EnrolledCoursesResponseDTO, error) {
	args := m.Called(ctx, logger, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dtos.EnrolledCoursesResponseDTO), args.Error(1)
}

func (m *MockCourseService) GetCourses(ctx context.Context, logger *log.Logger, username string, page, pageSize int) ([]dtos.CourseResponseDTO, error) {
	args := m.Called(ctx, logger, username, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dtos.CourseResponseDTO), args.Error(1)
}

func (m *MockCourseService) EnrollUser(ctx context.Context, logger *log.Logger, username string, courseID int) error {
	args := m.Called(ctx, logger, username, courseID)
	return args.Error(0)
}
