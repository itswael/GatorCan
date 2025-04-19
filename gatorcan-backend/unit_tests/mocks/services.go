package mocks

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/models"
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

func (m *MockCourseService) GetCourseByID(ctx context.Context, logger *log.Logger, courseID int) (dtos.CourseResponseDTO, error) {
	args := m.Called(ctx, courseID)
	return args.Get(0).(dtos.CourseResponseDTO), args.Error(1)
}

// MockUserService mocks the user service interface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(ctx context.Context, logger *log.Logger, dto *dtos.UserRequestDTO) (*dtos.UserResponseDTO, error) {
	args := m.Called(ctx, dto)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dtos.UserResponseDTO), args.Error(1)
}

func (m *MockUserService) Login(ctx context.Context, dto *dtos.LoginRequestDTO) (*dtos.LoginResponseDTO, error) {
	args := m.Called(ctx, dto)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dtos.LoginResponseDTO), args.Error(1)
}

func (m *MockUserService) GetUserDetails(ctx context.Context, username string) (*models.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) DeleteUser(ctx context.Context, username string) error {
	args := m.Called(ctx, username)
	return args.Error(0)
}

func (m *MockUserService) UpdateUser(ctx context.Context, username string, updateData *dtos.UpdateUserDTO) error {
	args := m.Called(ctx, username, updateData)
	return args.Error(0)
}

func (m *MockUserService) UpdateRoles(ctx context.Context, username string, roles []string) error {
	args := m.Called(ctx, username, roles)
	return args.Error(0)
}

// MockAssignmentService mocks the AssignmentService interface
type MockAssignmentService struct {
	mock.Mock
}

func (m *MockAssignmentService) GetAssignmentsByCourseID(ctx context.Context, courseID int) ([]dtos.AssignmentResponseDTO, error) {
	args := m.Called(ctx, courseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]dtos.AssignmentResponseDTO), args.Error(1)
}

func (m *MockAssignmentService) GetAssignmentByIDAndCourseID(ctx context.Context, assignmentID, courseID int) (dtos.AssignmentResponseDTO, error) {
	args := m.Called(ctx, assignmentID, courseID)
	return args.Get(0).(dtos.AssignmentResponseDTO), args.Error(1)
}

func (m *MockAssignmentService) UploadFileToAssignment(ctx context.Context, logger *log.Logger, username string, uploadData *dtos.UploadFileToAssignmentDTO) (*dtos.UploadFileToAssignmentResponseDTO, error) {
	args := m.Called(ctx, logger, username, uploadData)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dtos.UploadFileToAssignmentResponseDTO), args.Error(1)
}

// Add this to your existing services.go file after the MockAssignmentService:

// MockSubmissionService mocks the SubmissionService interface
type MockSubmissionService struct {
	mock.Mock
}

func (m *MockSubmissionService) GetSubmission(ctx context.Context, courseID, assignmentID int, userID uint) (*dtos.SubmissionResponseDTO, error) {
	args := m.Called(ctx, courseID, assignmentID, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dtos.SubmissionResponseDTO), args.Error(1)
}

func (m *MockSubmissionService) GradeSubmission(ctx context.Context, logger *log.Logger, instructorUsername string, gradeData *dtos.GradeSubmissionRequestDTO) (*dtos.GradeSubmissionResponseDTO, error) {
	args := m.Called(ctx, logger, instructorUsername, gradeData)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dtos.GradeSubmissionResponseDTO), args.Error(1)
}

func (m *MockAssignmentService) UpsertAssignment(ctx context.Context, logger *log.Logger, assignment *dtos.CreateAssignmentRequestDTO, existingassignment *dtos.UpdateAssignmentRequestDTO) (dtos.AssignmentResponseDTO, error) {
	args := m.Called(ctx, logger, assignment)
	return args.Get(0).(dtos.AssignmentResponseDTO), args.Error(1)
}
