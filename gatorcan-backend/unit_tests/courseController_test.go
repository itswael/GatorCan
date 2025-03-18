package unit_tests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/controllers"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

// Create domain error constants to match those used in the controller
var (
	ErrUserNotFound    = errors.New("user not found")
	ErrCourseNotFound  = errors.New("course not found")
	ErrAlreadyEnrolled = errors.New("enrollment request already exists")
	ErrCourseFull      = errors.New("course has reached maximum capacity")
	ErrFailedToEnroll  = errors.New("failed to request enrollment")
)

func TestGetEnrolledCourses(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	tests := []struct {
		name           string
		username       string
		setupMock      func(*MockCourseService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:     "Success with courses",
			username: "testuser",
			setupMock: func(m *MockCourseService) {
				courses := []dtos.EnrolledCoursesResponseDTO{
					{
						ID:              1,
						Name:            "Test Course",
						Description:     "Test Description",
						InstructorName:  "instructor",
						InstructorEmail: "instructor@example.com",
					},
				}
				m.On("GetEnrolledCourses", mock.Anything, mock.Anything, "testuser").
					Return(courses, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"name":"Test Course","description":"Test Description","instructor_name":"instructor","instructor_email":"instructor@example.com"}]`,
		},
		{
			name:     "Success with empty courses",
			username: "newuser",
			setupMock: func(m *MockCourseService) {
				m.On("GetEnrolledCourses", mock.Anything, mock.Anything, "newuser").
					Return([]dtos.EnrolledCoursesResponseDTO{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[]`,
		},
		{
			name:     "User not found",
			username: "nonexistent",
			setupMock: func(m *MockCourseService) {
				m.On("GetEnrolledCourses", mock.Anything, mock.Anything, "nonexistent").
					Return(nil, ErrUserNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"User not found"}`,
		},
		{
			name:     "Database error",
			username: "testuser",
			setupMock: func(m *MockCourseService) {
				m.On("GetEnrolledCourses", mock.Anything, mock.Anything, "testuser").
					Return(nil, errors.New("database connection failed"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Failed to fetch enrolled courses"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service
			mockService := new(MockCourseService)
			tc.setupMock(mockService)

			// Create controller
			courseController := controllers.NewCourseController(mockService, logger)

			// Setup HTTP request context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Set username in context
			c.Set("username", tc.username)

			// Create mock request with context
			req := httptest.NewRequest("GET", "/courses/enrolled", nil)
			c.Request = req

			// Execute controller method
			courseController.GetEnrolledCourses(c)

			// Assert response
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, w.Body.String())
			}

			// Verify all mock expectations were met
			mockService.AssertExpectations(t)
		})
	}
}

func TestGetCourses(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	tests := []struct {
		name           string
		username       string
		queryParams    map[string]string
		setupMock      func(*MockCourseService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "Success with default pagination",
			username:    "testuser",
			queryParams: map[string]string{},
			setupMock: func(m *MockCourseService) {
				courses := []dtos.CourseResponseDTO{
					{
						ID:          1,
						Name:        "Test Course",
						Description: "Test Description",
					},
				}
				m.On("GetCourses", mock.Anything, mock.Anything, "testuser", 1, 10).
					Return(courses, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"name":"Test Course","description":"Test Description"}]`,
		},
		{
			name:     "Success with custom pagination",
			username: "testuser",
			queryParams: map[string]string{
				"page":     "2",
				"pageSize": "5",
			},
			setupMock: func(m *MockCourseService) {
				m.On("GetCourses", mock.Anything, mock.Anything, "testuser", 2, 5).
					Return([]dtos.CourseResponseDTO{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[]`,
		},
		{
			name:     "Invalid page parameter",
			username: "testuser",
			queryParams: map[string]string{
				"page":     "invalid",
				"pageSize": "10",
			},
			setupMock: func(m *MockCourseService) {
				// Should use default page = 1
				m.On("GetCourses", mock.Anything, mock.Anything, "testuser", 1, 10).
					Return([]dtos.CourseResponseDTO{}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `[]`,
		},
		{
			name:        "Service error",
			username:    "testuser",
			queryParams: map[string]string{},
			setupMock: func(m *MockCourseService) {
				m.On("GetCourses", mock.Anything, mock.Anything, "testuser", 1, 10).
					Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Failed to fetch courses"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service
			mockService := new(MockCourseService)
			tc.setupMock(mockService)

			// Create controller
			courseController := controllers.NewCourseController(mockService, logger)

			// Setup HTTP request context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Set username in context
			c.Set("username", tc.username)

			// Create mock request with context and query parameters
			req := httptest.NewRequest("GET", "/courses", nil)
			q := req.URL.Query()
			for key, value := range tc.queryParams {
				q.Add(key, value)
			}
			req.URL.RawQuery = q.Encode()
			c.Request = req

			// Add query parameters to context
			for key, value := range tc.queryParams {
				c.Request.URL.Query().Set(key, value)
			}

			// Execute controller method
			courseController.GetCourses(c)

			// Assert response
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, w.Body.String())
			}

			// Verify all mock expectations were met
			mockService.AssertExpectations(t)
		})
	}
}

func TestEnrollInCourse(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	tests := []struct {
		name           string
		username       string
		requestBody    map[string]interface{}
		setupMock      func(*MockCourseService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:     "Success",
			username: "testuser",
			requestBody: map[string]interface{}{
				"CourseID": 1,
			},
			setupMock: func(m *MockCourseService) {
				m.On("EnrollUser", mock.Anything, mock.Anything, "testuser", 1).
					Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"message":"Enrollment requested successfully"}`,
		},
		{
			name:     "Invalid course ID",
			username: "testuser",
			requestBody: map[string]interface{}{
				"CourseID": 0, // Invalid ID
			},
			setupMock: func(m *MockCourseService) {
				// No mock needed - validation should catch it before service call
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid course ID"}`,
		},
		{
			name:        "Missing course ID",
			username:    "testuser",
			requestBody: map[string]interface{}{
				// Missing course_id field
			},
			setupMock: func(m *MockCourseService) {
				// No mock needed - validation should catch it before service call
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid request body"}`,
		},
		{
			name:     "User not found",
			username: "nonexistent",
			requestBody: map[string]interface{}{
				"CourseID": 1,
			},
			setupMock: func(m *MockCourseService) {
				m.On("EnrollUser", mock.Anything, mock.Anything, "nonexistent", 1).
					Return(ErrUserNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"User not found"}`,
		},
		{
			name:     "Course not found",
			username: "testuser",
			requestBody: map[string]interface{}{
				"CourseID": 999,
			},
			setupMock: func(m *MockCourseService) {
				m.On("EnrollUser", mock.Anything, mock.Anything, "testuser", 999).
					Return(ErrCourseNotFound)
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid course ID"}`,
		},
		{
			name:     "Already enrolled",
			username: "testuser",
			requestBody: map[string]interface{}{
				"CourseID": 1,
			},
			setupMock: func(m *MockCourseService) {
				m.On("EnrollUser", mock.Anything, mock.Anything, "testuser", 1).
					Return(ErrAlreadyEnrolled)
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid course ID"}`,
		},
		{
			name:     "Course full",
			username: "testuser",
			requestBody: map[string]interface{}{
				"CourseID": 2,
			},
			setupMock: func(m *MockCourseService) {
				m.On("EnrollUser", mock.Anything, mock.Anything, "testuser", 2).
					Return(ErrCourseFull)
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid course ID"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service
			mockService := new(MockCourseService)
			tc.setupMock(mockService)

			// Create controller
			courseController := controllers.NewCourseController(mockService, logger)

			// Setup HTTP request context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Set username in context
			c.Set("username", tc.username)

			// Create request body
			jsonData, _ := json.Marshal(tc.requestBody)

			// Create mock request
			req := httptest.NewRequest("POST", "/courses/enroll", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// Execute controller method
			courseController.EnrollInCourse(c)

			// Assert response
			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, w.Body.String())
			}

			// Verify all mock expectations were met
			mockService.AssertExpectations(t)
		})
	}
}
