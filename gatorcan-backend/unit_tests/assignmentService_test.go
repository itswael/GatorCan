package unit_tests

import (
	"context"
	"errors"
	"gatorcan-backend/config"
	domainErrors "gatorcan-backend/errors"
	"gatorcan-backend/models"
	"gatorcan-backend/services"
	"gatorcan-backend/unit_tests/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAssignmentsByCourseID(t *testing.T) {
	// Setup mocks
	mockAssignmentRepo := new(mocks.MockAssignmentRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create service with mocks
	assignmentService := services.NewAssignmentService(
		mockAssignmentRepo,
		mockUserRepo,
		mockCourseRepo,
		appConfig,
		mockHTTPClient,
	)

	// Setup context
	ctx := context.Background()

	// Test data
	now := time.Now()
	testAssignments := []models.Assignment{
		{
			ID:             1,
			Title:          "Assignment 1",
			Description:    "Description 1",
			Deadline:       now.Add(24 * time.Hour),
			ActiveCourseID: 101,
			MaxPoints:      100,
		},
		{
			ID:             2,
			Title:          "Assignment 2",
			Description:    "Description 2",
			Deadline:       now.Add(48 * time.Hour),
			ActiveCourseID: 101,
			MaxPoints:      50,
		},
	}

	tests := []struct {
		name          string
		courseID      int
		mockSetup     func()
		expectedCount int
		expectError   bool
		errorType     error
	}{
		{
			name:     "Success",
			courseID: 101,
			mockSetup: func() {
				mockAssignmentRepo.On("GetAssignmentsByCourseID", ctx, 101).Return(testAssignments, nil).Once()
			},
			expectedCount: 2,
			expectError:   false,
		},
		{
			name:     "No Assignments Found",
			courseID: 999,
			mockSetup: func() {
				mockAssignmentRepo.On("GetAssignmentsByCourseID", ctx, 999).Return(nil, errors.New("not found")).Once()
			},
			expectedCount: 0,
			expectError:   true,
			errorType:     domainErrors.ErrAssignmentNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mocks
			tc.mockSetup()

			// Call the service
			assignments, err := assignmentService.GetAssignmentsByCourseID(ctx, tc.courseID)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
				if tc.errorType != nil {
					assert.Equal(t, tc.errorType, err)
				}
				assert.Empty(t, assignments)
			} else {
				assert.NoError(t, err)
				assert.Len(t, assignments, tc.expectedCount)

				// Verify DTO conversion
				for i, assignment := range assignments {
					assert.Equal(t, testAssignments[i].ID, assignment.ID)
					assert.Equal(t, testAssignments[i].Title, assignment.Title)
					assert.Equal(t, testAssignments[i].Description, assignment.Description)
					assert.Equal(t, testAssignments[i].Deadline, assignment.Deadline)
					assert.Equal(t, testAssignments[i].ActiveCourseID, assignment.ActiveCourseID)
					assert.Equal(t, testAssignments[i].MaxPoints, assignment.MaxPoints)
				}
			}

			// Verify mock expectations
			mockAssignmentRepo.AssertExpectations(t)
		})
	}
}

func TestGetAssignmentByIDAndCourseID(t *testing.T) {
	// Setup mocks
	mockAssignmentRepo := new(mocks.MockAssignmentRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create service with mocks
	assignmentService := services.NewAssignmentService(
		mockAssignmentRepo,
		mockUserRepo,
		mockCourseRepo,
		appConfig,
		mockHTTPClient,
	)

	// Setup context
	ctx := context.Background()

	// Test data
	now := time.Now()
	testAssignment := models.Assignment{
		ID:             1,
		Title:          "Test Assignment",
		Description:    "Test Description",
		Deadline:       now.Add(24 * time.Hour),
		ActiveCourseID: 101,
		MaxPoints:      100,
	}

	testCourse := models.ActiveCourse{
		ID:           101,
		CourseID:     201,
		InstructorID: 301,
	}

	tests := []struct {
		name         string
		assignmentID int
		courseID     int
		mockSetup    func()
		expectError  bool
		errorType    error
	}{
		{
			name:         "Success",
			assignmentID: 1,
			courseID:     101,
			mockSetup: func() {
				mockAssignmentRepo.On("GetAssignmentByIDAndCourseID", ctx, 1, 101).Return(testAssignment, nil).Once()
				mockCourseRepo.On("GetCourseByID", ctx, 101).Return(testCourse, nil).Once()
			},
			expectError: false,
		},
		{
			name:         "Assignment Not Found",
			assignmentID: 999,
			courseID:     101,
			mockSetup: func() {
				mockAssignmentRepo.On("GetAssignmentByIDAndCourseID", ctx, 999, 101).Return(models.Assignment{}, errors.New("not found")).Once()
			},
			expectError: true,
			errorType:   domainErrors.ErrAssignmentNotFound,
		},
		{
			name:         "Course Not Found",
			assignmentID: 1,
			courseID:     999,
			mockSetup: func() {
				mockAssignmentRepo.On("GetAssignmentByIDAndCourseID", ctx, 1, 999).Return(testAssignment, nil).Once()
				mockCourseRepo.On("GetCourseByID", ctx, 999).Return(models.ActiveCourse{}, errors.New("not found")).Once()
			},
			expectError: true,
			errorType:   domainErrors.ErrCourseNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mocks
			tc.mockSetup()

			// Call the service
			assignment, err := assignmentService.GetAssignmentByIDAndCourseID(ctx, tc.assignmentID, tc.courseID)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
				if tc.errorType != nil {
					assert.Equal(t, tc.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testAssignment.ID, assignment.ID)
				assert.Equal(t, testAssignment.Title, assignment.Title)
				assert.Equal(t, testAssignment.Description, assignment.Description)
				assert.Equal(t, testAssignment.Deadline, assignment.Deadline)
				assert.Equal(t, testAssignment.ActiveCourseID, assignment.ActiveCourseID)
				assert.Equal(t, testAssignment.MaxPoints, assignment.MaxPoints)
			}

			// Verify mock expectations
			mockAssignmentRepo.AssertExpectations(t)
			mockCourseRepo.AssertExpectations(t)
		})
	}
}
