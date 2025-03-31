package unit_tests

import (
	"context"
	"errors"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/config"
	domainErrors "gatorcan-backend/errors"
	"gatorcan-backend/models"
	"gatorcan-backend/services"
	"gatorcan-backend/unit_tests/mocks"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGradeSubmission_service(t *testing.T) {
	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create logger
	logger := log.New(os.Stdout, "TEST: ", log.LstdFlags)

	// Setup context (common for all tests)
	ctx := context.Background()

	// Define our test cases with a mockSetup that accepts fresh mocks
	tests := []struct {
		name           string
		username       string
		submissionData *dtos.GradeSubmissionRequestDTO
		mockSetup      func(
			userRepo *mocks.MockUserRepository,
			courseRepo *mocks.MockCourseRepository,
			submissionRepo *mocks.MockSubmissionRepository,
			httpClient *mocks.MockHTTPClient,
		)
		expectError      bool
		expectedError    error
		expectedResponse *dtos.GradeSubmissionResponseDTO
	}{
		{
			name:     "Valid submission grading",
			username: "testuser",
			submissionData: &dtos.GradeSubmissionRequestDTO{
				AssignmentID: 1,
				CourseID:     1,
				UserID:       1,
				Grade:        85,
				Feedback:     "Good work!",
			},
			mockSetup: func(
				userRepo *mocks.MockUserRepository,
				courseRepo *mocks.MockCourseRepository,
				submissionRepo *mocks.MockSubmissionRepository,
				httpClient *mocks.MockHTTPClient,
			) {
				userRepo.On("GetUserByUsername", ctx, "testuser").
					Return(&models.User{Model: gorm.Model{ID: 1}}, nil)
				// Return an ActiveCourse value (not a pointer) as expected by the service
				courseRepo.On("GetCourseByID", ctx, 1).
					Return(models.ActiveCourse{ID: 1, CourseID: 1, InstructorID: 1}, nil)
				submissionRepo.On("GradeSubmission", ctx, uint(1), uint(1), uint(1), float64(85), "Good work!").
					Return(nil)
			},
			expectError: false,
			expectedResponse: &dtos.GradeSubmissionResponseDTO{
				AssignmentID: 1,
				CourseID:     1,
				UserID:       1,
				Grade:        85,
				Feedback:     "Good work!",
			},
		},
		{
			name:     "User not found",
			username: "nonexistentuser",
			submissionData: &dtos.GradeSubmissionRequestDTO{
				AssignmentID: 1,
				CourseID:     1,
				UserID:       1,
				Grade:        85,
				Feedback:     "Good work!",
			},
			mockSetup: func(
				userRepo *mocks.MockUserRepository,
				courseRepo *mocks.MockCourseRepository,
				submissionRepo *mocks.MockSubmissionRepository,
				httpClient *mocks.MockHTTPClient,
			) {
				userRepo.On("GetUserByUsername", ctx, "nonexistentuser").
					Return(nil, errors.New("not found"))
			},
			expectError:      true,
			expectedError:    domainErrors.ErrUserNotFound,
			expectedResponse: nil,
		},
		{
			name:     "Course not found",
			username: "testuser",
			submissionData: &dtos.GradeSubmissionRequestDTO{
				AssignmentID: 1,
				CourseID:     1,
				UserID:       1,
				Grade:        85,
				Feedback:     "Good work!",
			},
			mockSetup: func(
				userRepo *mocks.MockUserRepository,
				courseRepo *mocks.MockCourseRepository,
				submissionRepo *mocks.MockSubmissionRepository,
				httpClient *mocks.MockHTTPClient,
			) {
				userRepo.On("GetUserByUsername", ctx, "testuser").
					Return(&models.User{Model: gorm.Model{ID: 1}}, nil)
				// Return zero value for ActiveCourse and an error to simulate course not found
				courseRepo.On("GetCourseByID", ctx, 1).
					Return(models.ActiveCourse{}, errors.New("not found"))
			},
			expectError:      true,
			expectedError:    domainErrors.ErrCourseNotFound,
			expectedResponse: nil,
		},
		{
			name:     "Grading submission failed",
			username: "testuser",
			submissionData: &dtos.GradeSubmissionRequestDTO{
				AssignmentID: 1,
				CourseID:     1,
				UserID:       1,
				Grade:        85,
				Feedback:     "Good work!",
			},
			mockSetup: func(
				userRepo *mocks.MockUserRepository,
				courseRepo *mocks.MockCourseRepository,
				submissionRepo *mocks.MockSubmissionRepository,
				httpClient *mocks.MockHTTPClient,
			) {
				userRepo.On("GetUserByUsername", ctx, "testuser").
					Return(&models.User{Model: gorm.Model{ID: 1}}, nil)
				courseRepo.On("GetCourseByID", ctx, 1).
					Return(models.ActiveCourse{ID: 1, CourseID: 1, InstructorID: 1}, nil)
				submissionRepo.On("GradeSubmission", ctx, uint(1), uint(1), uint(1), float64(85), "Good work!").
					Return(errors.New("grading failed"))
			},
			expectError:      true,
			expectedError:    domainErrors.ErrGradingSubmissionFailed,
			expectedResponse: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create fresh mocks for this test case
			mockUserRepo := new(mocks.MockUserRepository)
			mockCourseRepo := new(mocks.MockCourseRepository)
			mockSubmissionRepo := new(mocks.MockSubmissionRepository)
			mockHTTPClient := new(mocks.MockHTTPClient)

			// Create a fresh service instance with these mocks
			submissionService := services.NewSubmissionService(
				mockSubmissionRepo,
				nil, // assignment repo is not used
				mockUserRepo,
				mockCourseRepo,
				appConfig,
				mockHTTPClient,
			)

			// Set up mocks for this test case
			tc.mockSetup(mockUserRepo, mockCourseRepo, mockSubmissionRepo, mockHTTPClient)

			// Call the service
			response, err := submissionService.GradeSubmission(ctx, logger, tc.username, tc.submissionData)

			// Assert on the error and response
			if tc.expectError {
				assert.Error(t, err)
				if tc.expectedError != nil {
					assert.Equal(t, tc.expectedError, err)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResponse, response)
			}

			// Assert that all expectations were met for this test case
			mockUserRepo.AssertExpectations(t)
			mockCourseRepo.AssertExpectations(t)
			mockSubmissionRepo.AssertExpectations(t)
			mockHTTPClient.AssertExpectations(t)
		})
	}
}
