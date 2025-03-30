package unit_tests

import (
	"context"
	"errors"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/config"
	domainErrors "gatorcan-backend/errors"
	"gatorcan-backend/models"
	"gatorcan-backend/services"
	"io"
	"log"
	"testing"
	"time"

	"gatorcan-backend/unit_tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetCoursesService(t *testing.T) {
	// Create mocks
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create service with mocks
	courseService := services.NewCourseService(mockCourseRepo, mockUserRepo, appConfig, mockHTTPClient)

	// Setup context and logger
	ctx := context.Background()
	logger := log.New(io.Discard, "", 0)

	// Create test user
	testUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
	}

	tests := []struct {
		name          string
		username      string
		page          int
		pageSize      int
		mockCourses   []models.Course
		mockError     error
		expectedCount int
		expectError   bool
		errorType     error
	}{
		{
			name:     "Success - Full Page",
			username: "testuser",
			page:     1,
			pageSize: 20,
			mockCourses: []models.Course{
				{ID: 1, Name: "Course 1", Description: "Description 1"},
				{ID: 2, Name: "Course 2", Description: "Description 2"},
			},
			mockError:     nil,
			expectedCount: 2,
			expectError:   false,
		},
		{
			name:          "Success - Empty Page",
			username:      "testuser",
			page:          2,
			pageSize:      20,
			mockCourses:   []models.Course{},
			mockError:     nil,
			expectedCount: 0,
			expectError:   false,
		},
		{
			name:          "User Not Found",
			username:      "nonexistent",
			page:          1,
			pageSize:      20,
			mockCourses:   nil,
			mockError:     errors.New("user not found"),
			expectedCount: 0,
			expectError:   true,
			errorType:     domainErrors.ErrUserNotFound,
		},
		{
			name:          "Database Error",
			username:      "testuser",
			page:          1,
			pageSize:      20,
			mockCourses:   nil,
			mockError:     errors.New("database error"),
			expectedCount: 0,
			expectError:   true,
			errorType:     domainErrors.ErrFailedToFetch,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock expectations based on test case
			if tc.username == "nonexistent" {
				mockUserRepo.On("GetUserByUsername", ctx, tc.username).Return(nil, tc.mockError).Once()
			} else {
				mockUserRepo.On("GetUserByUsername", ctx, tc.username).Return(testUser, nil).Once()
				mockCourseRepo.On("GetCourses", ctx, tc.page, tc.pageSize).Return(tc.mockCourses, tc.mockError).Once()
			}

			// Call the service
			courses, err := courseService.GetCourses(ctx, logger, tc.username, tc.page, tc.pageSize)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
				if tc.errorType != nil {
					assert.Equal(t, tc.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedCount, len(courses))

				// Verify the DTO conversion
				if len(courses) > 0 {
					assert.Equal(t, tc.mockCourses[0].Name, courses[0].Name)
					assert.Equal(t, tc.mockCourses[0].Description, courses[0].Description)
				}
			}

			// Verify that mock expectations were met
			mockUserRepo.AssertExpectations(t)
			mockCourseRepo.AssertExpectations(t)
		})
	}
}

func TestGetEnrolledCourses_service(t *testing.T) {
	// Create mocks
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create service with mocks
	courseService := services.NewCourseService(mockCourseRepo, mockUserRepo, appConfig, mockHTTPClient)

	// Setup context and logger
	ctx := context.Background()
	logger := log.New(io.Discard, "", 0)

	// Create test data
	now := time.Now()
	testUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
	}

	testInstructor := &models.User{
		Username: "instructor",
		Email:    "instructor@example.com",
	}

	testCourse := models.Course{
		ID:          1,
		Name:        "Test Course",
		Description: "Test Description",
	}

	testActiveCourse := models.ActiveCourse{
		ID:           1,
		CourseID:     1,
		InstructorID: 2,
		StartDate:    now,
		EndDate:      now.AddDate(0, 4, 0), // 4 months later
		Course:       testCourse,
	}

	testEnrollments := []models.Enrollment{
		{
			ID:             1,
			UserID:         1,
			ActiveCourseID: 1,
			ActiveCourse:   testActiveCourse,
			Status:         models.Approved,
		},
	}

	tests := []struct {
		name          string
		username      string
		mockUser      *models.User
		mockError     error
		enrollments   []models.Enrollment
		enrollmentErr error
		expectedCount int
		expectError   bool
		errorType     error
	}{
		{
			name:          "Success With Enrolled Courses",
			username:      "testuser",
			mockUser:      testUser,
			mockError:     nil,
			enrollments:   testEnrollments,
			enrollmentErr: nil,
			expectedCount: 1,
			expectError:   false,
		},
		{
			name:          "Success With No Courses",
			username:      "newuser",
			mockUser:      testUser,
			mockError:     nil,
			enrollments:   []models.Enrollment{},
			enrollmentErr: nil,
			expectedCount: 0,
			expectError:   false,
		},
		{
			name:          "User Not Found",
			username:      "nonexistent",
			mockUser:      nil,
			mockError:     errors.New("user not found"),
			enrollments:   nil,
			enrollmentErr: nil,
			expectedCount: 0,
			expectError:   true,
			errorType:     domainErrors.ErrUserNotFound,
		},
		{
			name:          "Error Fetching Enrollments",
			username:      "testuser",
			mockUser:      testUser,
			mockError:     nil,
			enrollments:   nil,
			enrollmentErr: errors.New("database error"),
			expectedCount: 0,
			expectError:   true,
			errorType:     domainErrors.ErrFailedToFetch,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock expectations based on test case
			mockUserRepo.On("GetUserByUsername", ctx, tc.username).Return(tc.mockUser, tc.mockError).Once()

			if tc.mockUser != nil && tc.mockError == nil {
				mockCourseRepo.On("GetEnrolledCourses", ctx, int(tc.mockUser.ID)).
					Return(tc.enrollments, tc.enrollmentErr).Once()

				if tc.enrollments != nil && len(tc.enrollments) > 0 {
					mockUserRepo.On("GetUserByID", ctx, uint(2)).Return(testInstructor, nil).Once()
				}
			}

			// Call the service
			enrolledCourses, err := courseService.GetEnrolledCourses(ctx, logger, tc.username)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
				if tc.errorType != nil {
					assert.Equal(t, tc.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedCount, len(enrolledCourses))

				// Verify data in result
				if len(enrolledCourses) > 0 {
					assert.Equal(t, testCourse.Name, enrolledCourses[0].Name)
					assert.Equal(t, testCourse.Description, enrolledCourses[0].Description)
					assert.Equal(t, testInstructor.Username, enrolledCourses[0].InstructorName)
					assert.Equal(t, testInstructor.Email, enrolledCourses[0].InstructorEmail)
				}
			}

			// Verify that mock expectations were met
			mockUserRepo.AssertExpectations(t)
			mockCourseRepo.AssertExpectations(t)
		})
	}
}

func TestEnrollUser(t *testing.T) {
	// Create mocks
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test", // Using test environment to skip notifications
	}

	// Create service with mocks
	courseService := services.NewCourseService(mockCourseRepo, mockUserRepo, appConfig, mockHTTPClient)

	// Setup context and logger
	ctx := context.Background()
	logger := log.New(io.Discard, "", 0)

	// Create test data
	now := time.Now()
	testUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
	}

	testActiveCourse := models.ActiveCourse{
		ID:           1,
		CourseID:     1,
		InstructorID: 2,
		StartDate:    now.AddDate(0, -1, 0), // 1 month ago
		EndDate:      now.AddDate(0, 3, 0),  // 3 months from now
		Capacity:     30,
		Enrolled:     15,
	}

	fullCourse := models.ActiveCourse{
		ID:           2,
		CourseID:     2,
		InstructorID: 2,
		StartDate:    now.AddDate(0, -1, 0), // 1 month ago
		EndDate:      now.AddDate(0, 3, 0),  // 3 months from now
		Capacity:     20,
		Enrolled:     20, // Full
	}

	inactiveCourse := models.ActiveCourse{
		ID:           3,
		CourseID:     3,
		InstructorID: 2,
		StartDate:    now.AddDate(0, 1, 0), // 1 month from now
		EndDate:      now.AddDate(0, 5, 0), // 5 months from now
		Capacity:     30,
		Enrolled:     0,
	}

	// Setup existing enrollments
	existingEnrollments := []models.Enrollment{
		{
			ID:             1,
			UserID:         1,
			ActiveCourseID: 4, // Already enrolled in course 4
			Status:         models.Approved,
		},
	}

	tests := []struct {
		name        string
		username    string
		courseID    int
		mockUser    *models.User
		mockCourse  models.ActiveCourse
		userError   error
		courseError error
		enrollError error
		expectError bool
		errorType   error
	}{
		{
			name:        "Success",
			username:    "testuser",
			courseID:    1,
			mockUser:    testUser,
			mockCourse:  testActiveCourse,
			userError:   nil,
			courseError: nil,
			enrollError: nil,
			expectError: false,
		},
		{
			name:        "User Not Found",
			username:    "nonexistent",
			courseID:    1,
			mockUser:    nil,
			userError:   errors.New("user not found"),
			expectError: true,
			errorType:   domainErrors.ErrUserNotFound,
		},
		{
			name:        "Course Not Found",
			username:    "testuser",
			courseID:    999,
			mockUser:    testUser,
			userError:   nil,
			courseError: errors.New("course not found"),
			expectError: true,
			errorType:   domainErrors.ErrCourseNotFound,
		},
		{
			name:        "Course Full",
			username:    "testuser",
			courseID:    2,
			mockUser:    testUser,
			mockCourse:  fullCourse,
			userError:   nil,
			courseError: nil,
			expectError: true,
			errorType:   domainErrors.ErrCourseFull,
		},
		{
			name:        "Course Not Active Yet",
			username:    "testuser",
			courseID:    3,
			mockUser:    testUser,
			mockCourse:  inactiveCourse,
			userError:   nil,
			courseError: nil,
			expectError: true,
			errorType:   domainErrors.ErrCourseInactive,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock expectations based on test case
			mockUserRepo.On("GetUserByUsername", ctx, tc.username).Return(tc.mockUser, tc.userError).Once()

			if tc.mockUser != nil && tc.userError == nil {
				mockCourseRepo.On("GetCourseByID", ctx, tc.courseID).Return(tc.mockCourse, tc.courseError).Once()

				if tc.courseError == nil {
					// Add this expectation for all test cases where we successfully find the course
					// This is important because the service checks for existing enrollments for all cases
					mockCourseRepo.On("GetEnrolledCourses", ctx, int(tc.mockUser.ID)).
						Return(existingEnrollments, nil).Once()

					// Only add the RequestEnrollment expectation for the success case
					if !tc.expectError {
						mockCourseRepo.On("RequestEnrollment", ctx, tc.mockUser.ID, tc.mockCourse.ID).
							Return(tc.enrollError).Once()
					}
				}
			}

			// Call the service
			err := courseService.EnrollUser(ctx, logger, tc.username, tc.courseID)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
				if tc.errorType != nil {
					assert.Equal(t, tc.errorType, err)
				}
			} else {
				assert.NoError(t, err)
			}

			// Verify that mock expectations were met
			mockUserRepo.AssertExpectations(t)
			mockCourseRepo.AssertExpectations(t)
		})
	}
}

func TestCourseResponseDTOConversion(t *testing.T) {
	// Create test data
	now := time.Now()
	courses := []models.Course{
		{
			ID:          1,
			Name:        "Test Course",
			Description: "Test Description",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{
			ID:          2,
			Name:        "Another Course",
			Description: "Another Description",
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	// Call the conversion function
	dtos := dtos.ConvertToCourseResponseDTOs(courses)

	// Verify results
	assert.Equal(t, 2, len(dtos))
	assert.Equal(t, uint(1), dtos[0].ID)
	assert.Equal(t, "Test Course", dtos[0].Name)
	assert.Equal(t, "Test Description", dtos[0].Description)

	assert.Equal(t, uint(2), dtos[1].ID)
	assert.Equal(t, "Another Course", dtos[1].Name)
	assert.Equal(t, "Another Description", dtos[1].Description)
}

func TestGetCourseByID(t *testing.T) {
	// Create mocks
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create service with mocks
	courseService := services.NewCourseService(mockCourseRepo, mockUserRepo, appConfig, mockHTTPClient)

	// Setup context and logger
	ctx := context.Background()
	logger := log.New(io.Discard, "", 0)

	// Create test data
	testActiveCourse := models.ActiveCourse{
		ID:           1,
		CourseID:     101,
		InstructorID: 201,
	}

	testCourse := models.Course{
		ID:          101,
		Name:        "Test Course",
		Description: "Test Description",
	}

	testInstructor := &models.User{
		Username: "instructor",
		Email:    "instructor@example.com",
	}

	tests := []struct {
		name            string
		courseID        int
		activeCourse    models.ActiveCourse
		activeCourseErr error
		course          models.Course
		courseErr       error
		instructor      *models.User
		instructorErr   error
		expectError     bool
		errorType       error
	}{
		{
			name:            "Success",
			courseID:        1,
			activeCourse:    testActiveCourse,
			activeCourseErr: nil,
			course:          testCourse,
			courseErr:       nil,
			instructor:      testInstructor,
			instructorErr:   nil,
			expectError:     false,
		},
		{
			name:            "Active Course Not Found",
			courseID:        999,
			activeCourse:    models.ActiveCourse{},
			activeCourseErr: errors.New("course not found"),
			course:          models.Course{},
			courseErr:       nil,
			instructor:      nil,
			instructorErr:   nil,
			expectError:     true,
			errorType:       domainErrors.ErrCourseNotFound,
		},
		{
			name:            "Course Details Not Found",
			courseID:        1,
			activeCourse:    testActiveCourse,
			activeCourseErr: nil,
			course:          models.Course{},
			courseErr:       errors.New("course details not found"),
			instructor:      nil,
			instructorErr:   nil,
			expectError:     true,
			errorType:       domainErrors.ErrCourseNotFound,
		},
		{
			name:            "Instructor Not Found",
			courseID:        1,
			activeCourse:    testActiveCourse,
			activeCourseErr: nil,
			course:          testCourse,
			courseErr:       nil,
			instructor:      nil,
			instructorErr:   errors.New("instructor not found"),
			expectError:     true,
			errorType:       domainErrors.ErrUserNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock expectations based on test case
			mockCourseRepo.On("GetCourseByID", ctx, tc.courseID).Return(tc.activeCourse, tc.activeCourseErr).Once()

			// Only set up the following expectations if we expect them to be called
			if tc.activeCourseErr == nil {
				mockCourseRepo.On("GetCourseDetails", ctx, tc.activeCourse.CourseID).Return(tc.course, tc.courseErr).Once()

				if tc.courseErr == nil {
					mockUserRepo.On("GetUserByID", ctx, tc.activeCourse.InstructorID).Return(tc.instructor, tc.instructorErr).Once()
				}
			}

			// Call the service
			course, err := courseService.GetCourseByID(ctx, logger, tc.courseID)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
				if tc.errorType != nil {
					assert.Equal(t, tc.errorType, err)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.course.ID, course.ID)
				assert.Equal(t, tc.course.Name, course.Name)
				assert.Equal(t, tc.course.Description, course.Description)
				assert.Equal(t, tc.instructor.Username, course.InstructorName)
				assert.Equal(t, tc.instructor.Email, course.InstructorEmail)
			}

			// Verify that mock expectations were met
			mockCourseRepo.AssertExpectations(t)
			mockUserRepo.AssertExpectations(t)
		})
	}
}
