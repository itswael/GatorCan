package unit_tests

import (
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/controllers"
	"gatorcan-backend/errors"
	"gatorcan-backend/models"
	"gatorcan-backend/unit_tests/mocks"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetSubmission(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	// Create mock user response
	testUser := &models.User{
		Username: "student1",
		Email:    "student1@example.com",
	}

	// Create mock submission response
	testSubmission := &dtos.SubmissionResponseDTO{
		FileURL:     "https://example.com/submission.pdf",
		SubmittedAt: time.Now().Format(time.RFC3339),
		Grade:       85,
		Feedback:    "Good work!",
		MaxPoints:   100,
	}

	tests := []struct {
		name                string
		username            string
		courseIDParam       string
		assignmentIDParam   string
		setupUserMock       func(*mocks.MockUserService)
		setupSubmissionMock func(*mocks.MockSubmissionService)
		expectedStatus      int
		expectedBody        string
		checkBody           bool
	}{
		{
			name:              "Success",
			username:          "student1",
			courseIDParam:     "1",
			assignmentIDParam: "2",
			setupUserMock: func(m *mocks.MockUserService) {
				m.On("GetUserDetails", mock.Anything, "student1").Return(testUser, nil)
			},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {
				m.On("GetSubmission", mock.Anything, 1, 2, uint(0)).Return(testSubmission, nil)
			},
			expectedStatus: http.StatusOK,
			checkBody:      false, // Cannot test exact JSON due to timestamps
		},
		{
			name:                "Unauthorized",
			username:            "",
			courseIDParam:       "1",
			assignmentIDParam:   "2",
			setupUserMock:       func(m *mocks.MockUserService) {},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {},
			expectedStatus:      http.StatusUnauthorized,
			expectedBody:        `{"error":"Unauthorized"}`,
			checkBody:           true,
		},
		{
			name:              "User Not Found",
			username:          "nonexistent",
			courseIDParam:     "1",
			assignmentIDParam: "2",
			setupUserMock: func(m *mocks.MockUserService) {
				m.On("GetUserDetails", mock.Anything, "nonexistent").Return(nil, errors.ErrUserNotFound)
			},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {},
			expectedStatus:      http.StatusInternalServerError,
			expectedBody:        `{"error":"Error fetching user ID"}`,
			checkBody:           true,
		},
		{
			name:              "Invalid Course ID",
			username:          "student1",
			courseIDParam:     "invalid",
			assignmentIDParam: "2",
			setupUserMock: func(m *mocks.MockUserService) {
				m.On("GetUserDetails", mock.Anything, "student1").Return(testUser, nil)
			},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {},
			expectedStatus:      http.StatusBadRequest,
			expectedBody:        `{"error":"Invalid course ID"}`,
			checkBody:           true,
		},
		{
			name:              "Invalid Assignment ID",
			username:          "student1",
			courseIDParam:     "1",
			assignmentIDParam: "invalid",
			setupUserMock: func(m *mocks.MockUserService) {
				m.On("GetUserDetails", mock.Anything, "student1").Return(testUser, nil)
			},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {},
			expectedStatus:      http.StatusBadRequest,
			expectedBody:        `{"error":"Invalid assignment ID"}`,
			checkBody:           true,
		},
		{
			name:              "Submission Not Found",
			username:          "student1",
			courseIDParam:     "1",
			assignmentIDParam: "999",
			setupUserMock: func(m *mocks.MockUserService) {
				m.On("GetUserDetails", mock.Anything, "student1").Return(testUser, nil)
			},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {
				m.On("GetSubmission", mock.Anything, 1, 999, uint(0)).Return(nil, errors.ErrSubmissionNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Submission not found"}`,
			checkBody:      true,
		},
		{
			name:              "Server Error",
			username:          "student1",
			courseIDParam:     "1",
			assignmentIDParam: "2",
			setupUserMock: func(m *mocks.MockUserService) {
				m.On("GetUserDetails", mock.Anything, "student1").Return(testUser, nil)
			},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {
				m.On("GetSubmission", mock.Anything, 1, 2, uint(0)).Return(nil, errors.ErrDatabaseError)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Error fetching submission"}`,
			checkBody:      true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock services
			mockSubmissionService := new(mocks.MockSubmissionService)
			mockUserService := new(mocks.MockUserService)

			tc.setupUserMock(mockUserService)
			tc.setupSubmissionMock(mockSubmissionService)

			// Create controller
			submissionController := controllers.NewSubmissionController(mockSubmissionService, mockUserService, logger)

			// Setup HTTP request context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Set username if provided
			if tc.username != "" {
				c.Set("username", tc.username)
			}

			// Add URL parameters
			c.Params = gin.Params{
				{Key: "cid", Value: tc.courseIDParam},
				{Key: "aid", Value: tc.assignmentIDParam},
			}

			// Create request
			req := httptest.NewRequest("GET", "/courses/"+tc.courseIDParam+"/assignments/"+tc.assignmentIDParam+"/submissions", nil)
			c.Request = req

			// Execute controller method
			submissionController.GetSubmission(c)

			// Assert status code
			assert.Equal(t, tc.expectedStatus, w.Code)

			// Assert response body if expected and checkBody is true
			if tc.checkBody && tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, w.Body.String())
			}

			// Verify mock expectations were met
			mockUserService.AssertExpectations(t)
			mockSubmissionService.AssertExpectations(t)
		})
	}
}
