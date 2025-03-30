package unit_tests

import (
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/controllers"
	"gatorcan-backend/errors"
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

func TestGetAssignments(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	tests := []struct {
		name           string
		courseIDParam  string
		setupMock      func(*mocks.MockAssignmentService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:          "Success",
			courseIDParam: "1",
			setupMock: func(m *mocks.MockAssignmentService) {
				assignments := []dtos.AssignmentResponseDTO{
					{
						ID:             1,
						Title:          "Assignment 1",
						Description:    "Test Description",
						Deadline:       time.Date(2025, 4, 15, 0, 0, 0, 0, time.UTC),
						ActiveCourseID: 1,
					},
					{
						ID:             2,
						Title:          "Assignment 2",
						Description:    "Another Test",
						Deadline:       time.Date(2025, 4, 30, 0, 0, 0, 0, time.UTC),
						ActiveCourseID: 1,
					},
				}
				m.On("GetAssignmentsByCourseID", mock.Anything, 1).Return(assignments, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"assignments":[{"id":1,"title":"Assignment 1","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","max_points":0,"description":"Test Description","deadline":"2025-04-15T00:00:00Z","course_id":1},{"id":2,"title":"Assignment 2","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","max_points":0,"description":"Another Test","deadline":"2025-04-30T00:00:00Z","course_id":1}]}`,
		},
		{
			name:           "Invalid Course ID",
			courseIDParam:  "invalid",
			setupMock:      func(m *mocks.MockAssignmentService) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid course ID"}`,
		},
		{
			name:          "No Assignments Found",
			courseIDParam: "999",
			setupMock: func(m *mocks.MockAssignmentService) {
				m.On("GetAssignmentsByCourseID", mock.Anything, 999).Return(nil, errors.ErrAssignmentNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Assignments not found"}`,
		},
		{
			name:          "Server Error",
			courseIDParam: "1",
			setupMock: func(m *mocks.MockAssignmentService) {
				m.On("GetAssignmentsByCourseID", mock.Anything, 1).Return(nil, errors.ErrDatabaseError)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Error getting assignments"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service
			mockService := new(mocks.MockAssignmentService)
			tc.setupMock(mockService)

			// Create controller
			assignmentController := controllers.NewAssignmentController(mockService, logger)

			// Setup HTTP request context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Add course ID as URL parameter
			c.Params = gin.Params{
				{Key: "cid", Value: tc.courseIDParam},
			}

			// Create mock request
			req := httptest.NewRequest("GET", "/courses/"+tc.courseIDParam+"/assignments", nil)
			c.Request = req

			// Execute controller method
			assignmentController.GetAssignments(c)

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

func TestGetAssignment(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	tests := []struct {
		name              string
		courseIDParam     string
		assignmentIDParam string
		setupMock         func(*mocks.MockAssignmentService)
		expectedStatus    int
		expectedBody      string
	}{
		{
			name:              "Success",
			courseIDParam:     "1",
			assignmentIDParam: "2",
			setupMock: func(m *mocks.MockAssignmentService) {
				assignment := dtos.AssignmentResponseDTO{
					ID:             2,
					Title:          "Assignment 2",
					Description:    "Test Description",
					Deadline:       time.Date(2025, 4, 15, 0, 0, 0, 0, time.UTC),
					ActiveCourseID: 1,
				}
				m.On("GetAssignmentByIDAndCourseID", mock.Anything, 2, 1).Return(assignment, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"assignments":{"id":2,"title":"Assignment 2","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","max_points":0,"description":"Test Description","deadline":"2025-04-15T00:00:00Z","course_id":1}}`,
		},
		{
			name:              "Invalid Course ID",
			courseIDParam:     "invalid",
			assignmentIDParam: "2",
			setupMock:         func(m *mocks.MockAssignmentService) {},
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      `{"error":"Invalid course ID"}`,
		},
		{
			name:              "Assignment Not Found",
			courseIDParam:     "1",
			assignmentIDParam: "999",
			setupMock: func(m *mocks.MockAssignmentService) {
				m.On("GetAssignmentByIDAndCourseID", mock.Anything, 999, 1).Return(dtos.AssignmentResponseDTO{}, errors.ErrAssignmentNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Assignment not found"}`,
		},
		{
			name:              "Server Error",
			courseIDParam:     "1",
			assignmentIDParam: "2",
			setupMock: func(m *mocks.MockAssignmentService) {
				m.On("GetAssignmentByIDAndCourseID", mock.Anything, 2, 1).Return(dtos.AssignmentResponseDTO{}, errors.ErrDatabaseError)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Error getting assignments"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service
			mockService := new(mocks.MockAssignmentService)
			tc.setupMock(mockService)

			// Create controller
			assignmentController := controllers.NewAssignmentController(mockService, logger)

			// Setup HTTP request context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Add URL parameters
			c.Params = gin.Params{
				{Key: "cid", Value: tc.courseIDParam},
				{Key: "aid", Value: tc.assignmentIDParam},
			}

			// Create mock request
			req := httptest.NewRequest("GET", "/courses/"+tc.courseIDParam+"/assignments/"+tc.assignmentIDParam, nil)
			c.Request = req

			// Execute controller method
			assignmentController.GetAssignment(c)

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
