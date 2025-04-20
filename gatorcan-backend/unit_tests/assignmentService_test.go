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
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
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

func TestUploadFileToAssignment_service(t *testing.T) {
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
	testFileName := "testfile.txt"
	testFileURL := "http://example.com/testfile.txt"
	testFileType := "text/plain"
	now := time.Now()

	// Mock user data (with proper ID set)
	mockUser := &models.User{
		Model:    gorm.Model{ID: 1},
		Username: "testuser",
		Roles:    []*models.Role{{Name: "student"}},
	}

	mockInstructor := &models.User{
		Username: "instructor",
		Roles:    []*models.Role{{Name: "instructor"}},
	}

	mockCourse := &models.Course{
		ID:          201,
		Name:        "Test Course",
		Description: "Test Description",
	}

	// Mock active course data
	mockActiveCourse := models.ActiveCourse{
		ID:           101,
		CourseID:     201,
		InstructorID: 301,
		StartDate:    now.Add(-24 * time.Hour),
		EndDate:      now.Add(24 * time.Hour),
		IsActive:     true,
		Instructor:   *mockInstructor,
		Course:       *mockCourse,
		Capacity:     30,
	}

	// Mock assignment data
	mockAssignment := &models.Assignment{
		ID:             1,
		Title:          "Test Assignment",
		Description:    "Test Description",
		Deadline:       now.Add(24 * time.Hour),
		ActiveCourseID: 101,
		MaxPoints:      100,
	}

	// We don't rely on a pre-built expected file object for timestamps.
	// Instead, we’ll use 'now' to compare within an acceptable range.

	tests := []struct {
		name         string
		courseID     int
		assignmentID int
		username     string
		fileName     string
		fileURL      string
		FileType     string
		mockSetup    func()
		expectError  bool
		errorType    error
	}{
		{
			name:         "Success",
			courseID:     101,
			assignmentID: 1,
			username:     "testuser",
			fileName:     testFileName,
			fileURL:      testFileURL,
			FileType:     testFileType,
			mockSetup: func() {
				// Expect GetUserByUsername with "testuser"
				mockUserRepo.On("GetUserByUsername", ctx, "testuser").
					Return(mockUser, nil).Once()

				// Expect GetCourseByID with 101 to return the active course
				mockCourseRepo.On("GetCourseByID", ctx, 101).
					Return(mockActiveCourse, nil).Once()

				// Expect GetAssignmentByIDAndCourseID with (1, 101) to return the assignment
				mockAssignmentRepo.On("GetAssignmentByIDAndCourseID", ctx, 1, 101).
					Return(*mockAssignment, nil).Once()

				// Expect CreateAssignmentFile to update the passed assignment file and simulate DB insertion by setting its ID.
				mockAssignmentRepo.On("CreateAssignmentFile", ctx, mock.AnythingOfType("*models.AssignmentFile")).
					Run(func(args mock.Arguments) {
						fileArg := args.Get(1).(*models.AssignmentFile)
						fileArg.FileName = testFileName
						fileArg.FileURL = testFileURL
						fileArg.FileType = testFileType
						fileArg.AssignmentID = 1
						fileArg.CreatedAt = now
						fileArg.UpdatedAt = now
						// Simulate DB-generated ID:
						fileArg.ID = 1
					}).
					Return(nil).Once()

				// Expect LinkUserToAssignmentFile to be called and return nil.
				mockAssignmentRepo.On("LinkUserToAssignmentFile", ctx, mock.AnythingOfType("*models.UserAssignmentFile")).
					Return(nil).Once()
			},
			expectError: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			dto := &dtos.UploadFileToAssignmentDTO{
				AssignmentID: uint(tc.assignmentID),
				CourseID:     uint(tc.courseID),
				FileURL:      tc.fileURL,
				FileName:     tc.fileName,
				FileType:     tc.FileType,
			}
			logger := log.New(os.Stdout, "test: ", log.LstdFlags)
			// Call the service with the provided username ("testuser")
			success, err := assignmentService.UploadFileToAssignment(ctx, logger, tc.username, dto)

			if tc.expectError {
				assert.Error(t, err)
				if tc.errorType != nil {
					assert.Equal(t, tc.errorType, err)
				}
				assert.Nil(t, success)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, success)
				assert.Equal(t, tc.fileName, success.FileName)
				assert.Equal(t, tc.fileURL, success.FileURL)
				assert.Equal(t, tc.FileType, success.FileType)
				// Instead of direct equality, assert that UploadedAt is within one second of 'now'.
				assert.WithinDuration(t, now, success.UploadedAt, time.Second)
			}

			mockAssignmentRepo.AssertExpectations(t)
			mockUserRepo.AssertExpectations(t)
			mockCourseRepo.AssertExpectations(t)
			mockHTTPClient.AssertExpectations(t)
		})
	}
}

func TestCreateOrUpdateAssignment(t *testing.T) {
	// Setup mocks
	mockAssignmentRepo := new(mocks.MockAssignmentRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create the service
	assignmentService := services.NewAssignmentService(
		mockAssignmentRepo,
		mockUserRepo,
		mockCourseRepo,
		appConfig,
		mockHTTPClient,
	)

	ctx := context.Background()
	now := time.Now()

	tests := []struct {
		name         string
		courseID     int
		username     string
		dto          *dtos.CreateOrUpdateAssignmentRequestDTO
		mockSetup    func()
		expectError  bool
		errorType    error
		expectedResp dtos.AssignmentResponseDTO
	}{
		{
			name:     "Create Assignment Success",
			courseID: 101,
			username: "testuser",
			dto: &dtos.CreateOrUpdateAssignmentRequestDTO{
				ID:          0,
				CourseID:    101,
				MaxPoints:   100,
				Description: "Test Description",
				Title:       "Test Assignment",
				Deadline:    now.Add(48 * time.Hour),
			},
			mockSetup: func() {
				mockCourseRepo.On("GetCourseByID", ctx, 101).
					Return(models.ActiveCourse{ID: 101}, nil).Once()

				mockAssignmentRepo.On("UpsertAssignment", ctx, mock.MatchedBy(func(a *models.Assignment) bool {
					return a.ActiveCourseID == 101 && a.Title == "Test Assignment"
				})).Run(func(args mock.Arguments) {
					arg := args.Get(1).(*models.Assignment)
					arg.ID = 1
					arg.CreatedAt = now
					arg.UpdatedAt = now
				}).Return(nil).Once()
			},
			expectError: false,
			expectedResp: dtos.AssignmentResponseDTO{
				ID:             1,
				Title:          "Test Assignment",
				Description:    "Test Description",
				Deadline:       now.Add(48 * time.Hour),
				ActiveCourseID: 101,
				MaxPoints:      100,
			},
		},
		{
			name:     "Course Not Found",
			courseID: 999,
			username: "testuser",
			dto: &dtos.CreateOrUpdateAssignmentRequestDTO{
				ID:          0,
				CourseID:    999,
				Title:       "Some Assignment",
				Description: "Blah",
				Deadline:    now.Add(24 * time.Hour),
			},
			mockSetup: func() {
				mockCourseRepo.On("GetCourseByID", ctx, 999).
					Return(models.ActiveCourse{}, domainErrors.ErrCourseNotFound).Once()
			},
			expectError:  true,
			errorType:    domainErrors.ErrCourseNotFound,
			expectedResp: dtos.AssignmentResponseDTO{},
		},
		{
			name:     "Upsert Error",
			courseID: 101,
			username: "testuser",
			dto: &dtos.CreateOrUpdateAssignmentRequestDTO{
				ID:          0,
				CourseID:    101,
				Title:       "Err Assignment",
				Description: "Fail here",
				Deadline:    now.Add(24 * time.Hour),
			},
			mockSetup: func() {
				mockCourseRepo.On("GetCourseByID", ctx, 101).
					Return(models.ActiveCourse{ID: 101}, nil).Once()

				mockAssignmentRepo.On("UpsertAssignment", ctx, mock.Anything).
					Return(domainErrors.ErrFailedToCreateAssignment).Once()
			},
			expectError:  true,
			errorType:    domainErrors.ErrFailedToCreateAssignment,
			expectedResp: dtos.AssignmentResponseDTO{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			logger := log.New(io.Discard, "", log.LstdFlags)

			resp, err := assignmentService.UpsertAssignment(ctx, logger, tc.dto)

			if tc.expectError {
				assert.Error(t, err)
				if tc.errorType != nil {
					assert.Equal(t, tc.errorType, err)
				}
				assert.Equal(t, tc.expectedResp, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResp.Title, resp.Title)
				assert.Equal(t, tc.expectedResp.Description, resp.Description)
				assert.Equal(t, tc.expectedResp.MaxPoints, resp.MaxPoints)
				assert.Equal(t, tc.expectedResp.ActiveCourseID, resp.ActiveCourseID)
			}

			mockAssignmentRepo.AssertExpectations(t)
			mockCourseRepo.AssertExpectations(t)
		})
	}
}
