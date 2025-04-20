package unit_tests

import (
	"fmt"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/controllers"
	"gatorcan-backend/errors"
	"gatorcan-backend/models"
	"gatorcan-backend/unit_tests/mocks"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
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
			MockAWSService := new(mocks.MockAWSService)

			tc.setupUserMock(mockUserService)
			tc.setupSubmissionMock(mockSubmissionService)

			// Create controller
			submissionController := controllers.NewSubmissionController(mockSubmissionService, mockUserService, MockAWSService, logger)

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

func TestGradeSubmission(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	tests := []struct {
		name                string
		username            string
		courseID            string
		gradeRequestBody    string
		setupUserMock       func(*mocks.MockUserService)
		setupSubmissionMock func(*mocks.MockSubmissionService)
		setupAwsMock        func(*mocks.MockAWSService)
		expectedStatus      int
		expectedBody        string
		checkBody           bool
	}{
		{
			name:     "Success",
			username: "instructor1",
			courseID: "1",
			gradeRequestBody: `{
				"assignment_id": 1,
				"user_id": 1,
				"course_id": 1,
				"grade": 90,
				"feedback": "Well done!"
			}`,
			setupUserMock: func(m *mocks.MockUserService) {
				// No need to mock any user service methods for GradeSubmission
			},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {
				// Create expected dto that matches the JSON in the request
				expectedDTO := &dtos.GradeSubmissionRequestDTO{
					AssignmentID: 1,
					UserID:       1,
					CourseID:     1,
					Grade:        90,
					Feedback:     "Well done!",
				}

				m.On("GradeSubmission", mock.Anything, mock.Anything, "instructor1", mock.MatchedBy(func(dto *dtos.GradeSubmissionRequestDTO) bool {
					return dto.AssignmentID == expectedDTO.AssignmentID &&
						dto.UserID == expectedDTO.UserID &&
						dto.CourseID == expectedDTO.CourseID &&
						dto.Grade == expectedDTO.Grade &&
						dto.Feedback == expectedDTO.Feedback
				})).Return(&dtos.GradeSubmissionResponseDTO{
					AssignmentID: 1,
					CourseID:     1,
					UserID:       1,
					Grade:        90,
					Feedback:     "Well done!",
				}, nil)
			},
			setupAwsMock: func(m *mocks.MockAWSService) {
				// Mock AWS SNS notification call
				expectedMessage := "Submission graded: 1 for user: instructor1"
				m.On("PushNotificationToSNS", mock.Anything, mock.Anything, expectedMessage).Return(nil)
			},
			expectedStatus: http.StatusOK,
			checkBody:      false,
		},
		{
			name:                "Unauthorized",
			username:            "",
			courseID:            "1",
			gradeRequestBody:    `{"assignment_id": 1, "user_id": 1, "course_id": 1, "grade": 90, "feedback": "Well done!"}`,
			setupUserMock:       func(m *mocks.MockUserService) {},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {},
			setupAwsMock:        func(m *mocks.MockAWSService) {},
			expectedStatus:      http.StatusUnauthorized,
			expectedBody:        `{"error":"Unauthorized"}`,
			checkBody:           true,
		},
		{
			name:     "Invalid Course ID",
			username: "instructor1",
			courseID: "invalid",
			gradeRequestBody: `{
				"assignment_id": 1,
				"user_id": 1,
				"course_id": 1,
				"grade": 90,
				"feedback": "Well done!"
			}`,
			setupUserMock:       func(m *mocks.MockUserService) {},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {},
			setupAwsMock:        func(m *mocks.MockAWSService) {},
			expectedStatus:      http.StatusBadRequest,
			expectedBody:        `{"error":"Invalid course ID"}`,
			checkBody:           true,
		},
		{
			name:     "Invalid Request Body",
			username: "instructor1",
			courseID: "1",
			gradeRequestBody: `{
				"invalid_json":
			}`,
			setupUserMock:       func(m *mocks.MockUserService) {},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {},
			setupAwsMock:        func(m *mocks.MockAWSService) {},
			expectedStatus:      http.StatusBadRequest,
			expectedBody:        `{"error":"Invalid request body"}`,
			checkBody:           true,
		},
		{
			name:     "Submission Not Found",
			username: "instructor1",
			courseID: "1",
			gradeRequestBody: `{
				"assignment_id": 99,
				"user_id": 1,
				"course_id": 1,
				"grade": 90,
				"feedback": "Well done!"
			}`,
			setupUserMock: func(m *mocks.MockUserService) {
				// No need to mock any user service methods
			},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {
				m.On("GradeSubmission", mock.Anything, mock.Anything, "instructor1", mock.Anything).Return(nil, errors.ErrSubmissionNotFound)
			},
			setupAwsMock: func(m *mocks.MockAWSService) {
				// No need to mock AWS service since the service function returns error before AWS call
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Submission not found"}`,
			checkBody:      true,
		},
		{
			name:     "Server Error",
			username: "instructor1",
			courseID: "1",
			gradeRequestBody: `{
				"assignment_id": 1,
				"user_id": 1,
				"course_id": 1,
				"grade": 90,
				"feedback": "Well done!"
			}`,
			setupUserMock: func(m *mocks.MockUserService) {
				// No need to mock any user service methods
			},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {
				m.On("GradeSubmission", mock.Anything, mock.Anything, "instructor1", mock.Anything).Return(nil, errors.ErrGradingSubmissionFailed)
			},
			setupAwsMock: func(m *mocks.MockAWSService) {
				// No need to mock AWS service since the service function returns error before AWS call
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Error grading submission"}`,
			checkBody:      true,
		},
		{
			name:     "SNS Notification Failed",
			username: "instructor1",
			courseID: "1",
			gradeRequestBody: `{
				"assignment_id": 1,
				"user_id": 1,
				"course_id": 1,
				"grade": 90,
				"feedback": "Well done!"
			}`,
			setupUserMock: func(m *mocks.MockUserService) {
				// No need to mock any user service methods
			},
			setupSubmissionMock: func(m *mocks.MockSubmissionService) {
				m.On("GradeSubmission", mock.Anything, mock.Anything, "instructor1", mock.Anything).Return(&dtos.GradeSubmissionResponseDTO{
					AssignmentID: 1,
					CourseID:     1,
					UserID:       1,
					Grade:        90,
					Feedback:     "Well done!",
				}, nil)
			},
			setupAwsMock: func(m *mocks.MockAWSService) {
				// Mock AWS SNS notification call to fail
				expectedMessage := "Submission graded: 1 for user: instructor1"
				m.On("PushNotificationToSNS", mock.Anything, mock.Anything, expectedMessage).Return(errors.ErrSNSNotificationFailed)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"failed to send SNS notification"}`,
			checkBody:      true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock services
			mockSubmissionService := new(mocks.MockSubmissionService)
			mockUserService := new(mocks.MockUserService)
			mockAwsService := new(mocks.MockAWSService)

			tc.setupUserMock(mockUserService)
			tc.setupSubmissionMock(mockSubmissionService)
			tc.setupAwsMock(mockAwsService)

			// Create controller
			submissionController := controllers.NewSubmissionController(mockSubmissionService, mockUserService, mockAwsService, logger)

			// Setup router with middleware to set username
			router := gin.New()
			router.Use(func(c *gin.Context) {
				if tc.username != "" {
					c.Set("username", tc.username)
				}
				c.Next()
			})

			// Add the route
			router.POST("/courses/:cid/submissions/grade", submissionController.GradeSubmission)

			// Setup the recorder
			w := httptest.NewRecorder()

			// Create the test URL with the course ID parameter
			url := fmt.Sprintf("/courses/%s/submissions/grade", tc.courseID)

			// Create request with JSON body
			req := httptest.NewRequest("POST", url, strings.NewReader(tc.gradeRequestBody))
			req.Header.Set("Content-Type", "application/json")

			// Execute the request
			router.ServeHTTP(w, req)

			// Assert status code
			assert.Equal(t, tc.expectedStatus, w.Code, "Expected status %d but got %d", tc.expectedStatus, w.Code)

			// Assert response body if expected and checkBody is true
			if tc.checkBody && tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, w.Body.String(), "Expected body %s but got %s", tc.expectedBody, w.Body.String())
			}

			// Verify mock expectations were met
			mockUserService.AssertExpectations(t)
			mockSubmissionService.AssertExpectations(t)
			mockAwsService.AssertExpectations(t)
		})
	}
}
