package unit_tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/controllers"
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

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	tests := []struct {
		name           string
		input          dtos.UserRequestDTO
		mockResponse   *dtos.UserResponseDTO
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success",
			input: dtos.UserRequestDTO{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password123",
				Roles:    []string{"student"},
			},
			mockResponse: &dtos.UserResponseDTO{
				Code:    http.StatusCreated,
				Message: "User created successfully",
				Err:     false,
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedBody:   `{"message":"User created successfully"}`,
		},
		{
			name: "Invalid Email",
			input: dtos.UserRequestDTO{
				Username: "testuser",
				Email:    "invalid-email",
				Password: "password123",
				Roles:    []string{"student"},
			},
			mockResponse:   nil,
			mockError:      nil, // No mock needed as validation happens in controller
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid email format"}`,
		},
		{
			name: "User Already Exists",
			input: dtos.UserRequestDTO{
				Username: "existinguser",
				Email:    "existing@example.com",
				Password: "password123",
				Roles:    []string{"student"},
			},
			mockResponse: &dtos.UserResponseDTO{
				Code:    http.StatusBadRequest,
				Message: "User already exists",
				Err:     true,
			},
			mockError:      errors.New("user already exists"),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"User already exists"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service
			mockService := new(mocks.MockUserService)

			// Only set expectation if we're making it past validation
			if tc.input.Email != "invalid-email" {
				mockService.On("CreateUser", mock.Anything, &tc.input).Return(tc.mockResponse, tc.mockError)
			}

			// Create controller with mock
			userController := controllers.NewUserController(mockService, logger)

			// Setup HTTP context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Create request with JSON body
			jsonInput, _ := json.Marshal(tc.input)
			req, _ := http.NewRequest("POST", "/admin/add_user", bytes.NewBuffer(jsonInput))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// Call controller method
			userController.CreateUser(c, logger)

			// Assertions
			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())

			// Verify mock expectations
			mockService.AssertExpectations(t)
		})
	}
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	tests := []struct {
		name           string
		input          dtos.LoginRequestDTO
		mockResponse   *dtos.LoginResponseDTO
		mockError      error
		expectedStatus int
		expectedBody   string
		expectedHeader string
	}{
		{
			name: "Success",
			input: dtos.LoginRequestDTO{
				Username: "testuser",
				Password: "password123",
			},
			mockResponse: &dtos.LoginResponseDTO{
				Code:    http.StatusOK,
				Message: "Login successful",
				Token:   "jwt-token",
				Err:     false,
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"Login successful","token":"jwt-token"}`,
			expectedHeader: "Bearer jwt-token",
		},
		{
			name: "Invalid Credentials",
			input: dtos.LoginRequestDTO{
				Username: "testuser",
				Password: "wrongpassword",
			},
			mockResponse: &dtos.LoginResponseDTO{
				Code:    http.StatusUnauthorized,
				Message: "Invalid credentials",
				Err:     true,
			},
			mockError:      errors.New("invalid credentials"),
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Invalid credentials"}`,
			expectedHeader: "",
		},
		{
			name: "User Not Found",
			input: dtos.LoginRequestDTO{
				Username: "nonexistent",
				Password: "password123",
			},
			mockResponse: &dtos.LoginResponseDTO{
				Code:    http.StatusNotFound,
				Message: "Invalid_username",
				Err:     true,
			},
			mockError:      errors.New("user not found"),
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Invalid_username"}`,
			expectedHeader: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service
			mockService := new(mocks.MockUserService)
			mockService.On("Login", mock.Anything, &tc.input).Return(tc.mockResponse, tc.mockError)

			// Create controller with mock
			userController := controllers.NewUserController(mockService, logger)

			// Setup HTTP context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Create request with JSON body
			jsonInput, _ := json.Marshal(tc.input)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// Call controller method
			userController.Login(c, logger)

			// Assertions
			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())

			// Check authorization header for successful login
			if tc.expectedHeader != "" {
				assert.Equal(t, tc.expectedHeader, w.Header().Get("Authorization"))
			}

			// Verify mock expectations
			mockService.AssertExpectations(t)
		})
	}
}

func TestGetUserDetails(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	// now := time.Now()
	mockUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Roles:    []*models.Role{{Name: "student"}},
	}

	tests := []struct {
		name           string
		username       string
		mockUser       *models.User
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Success",
			username:       "testuser",
			mockUser:       mockUser,
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody: fmt.Sprintf(`{"created_at":"%s","email":"test@example.com","roles":["student"],"username":"testuser"}`,
				mockUser.CreatedAt.Format(time.RFC3339Nano)),
		},
		{
			name:           "User Not Found",
			username:       "nonexistent",
			mockUser:       nil,
			mockError:      errors.New("user not found"),
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"user not found"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service
			mockService := new(mocks.MockUserService)
			mockService.On("GetUserDetails", mock.Anything, tc.username).Return(tc.mockUser, tc.mockError)

			// Create controller with mock
			userController := controllers.NewUserController(mockService, logger)

			// Setup HTTP context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{{Key: "username", Value: tc.username}}
			c.Request, _ = http.NewRequest("GET", "/user/"+tc.username, nil)

			// Call controller method
			userController.GetUserDetails(c, logger)

			// Assertions
			assert.Equal(t, tc.expectedStatus, w.Code)

			if tc.expectedStatus == http.StatusNotFound {
				assert.JSONEq(t, tc.expectedBody, w.Body.String())
			} else {
				// Skip exact JSON comparison for success case as we have a dynamic timestamp
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "testuser", response["username"])
				assert.Equal(t, "test@example.com", response["email"])
				assert.Equal(t, []interface{}{"student"}, response["roles"])
				assert.Contains(t, response, "created_at")
			}

			// Verify mock expectations
			mockService.AssertExpectations(t)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	tests := []struct {
		name           string
		username       string
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Success",
			username:       "testuser",
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"User testuser has been deleted successfully"}`,
		},
		{
			name:           "User Not Found",
			username:       "nonexistent",
			mockError:      errors.New("user not found"),
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"user not found"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service
			mockService := new(mocks.MockUserService)
			mockService.On("DeleteUser", mock.Anything, tc.username).Return(tc.mockError)

			// Create controller with mock
			userController := controllers.NewUserController(mockService, logger)

			// Setup HTTP context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = []gin.Param{{Key: "username", Value: tc.username}}
			c.Request, _ = http.NewRequest("DELETE", "/admin/"+tc.username, nil)

			// Call controller method
			userController.DeleteUser(c, logger)

			// Assertions
			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())

			// Verify mock expectations
			mockService.AssertExpectations(t)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	tests := []struct {
		name           string
		username       string
		input          dtos.UpdateUserDTO
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:     "Success",
			username: "testuser",
			input: dtos.UpdateUserDTO{
				OldPassword: "oldpassword",
				NewPassword: "newpassword",
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"User updated successfully: testuser"}`,
		},
		{
			name:     "Invalid Password",
			username: "testuser",
			input: dtos.UpdateUserDTO{
				OldPassword: "wrongpassword",
				NewPassword: "newpassword",
			},
			mockError:      errors.New("invalid old password"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"invalid old password"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service
			mockService := new(mocks.MockUserService)
			mockService.On("UpdateUser", mock.Anything, tc.username, &tc.input).Return(tc.mockError)

			// Create controller with mock
			userController := controllers.NewUserController(mockService, logger)

			// Setup HTTP context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Set("username", tc.username)

			// Create request with JSON body
			jsonInput, _ := json.Marshal(tc.input)
			req, _ := http.NewRequest("PUT", "/user/update", bytes.NewBuffer(jsonInput))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// Call controller method
			userController.UpdateUser(c, logger)

			// Assertions
			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())

			// Verify mock expectations
			mockService.AssertExpectations(t)
		})
	}
}

func TestUpdateRoles(t *testing.T) {
	gin.SetMode(gin.TestMode)
	logger := log.New(io.Discard, "", 0)

	tests := []struct {
		name           string
		input          dtos.UpdateUserRolesDTO
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Success",
			input: dtos.UpdateUserRolesDTO{
				Username: "testuser",
				Roles:    []string{"admin", "instructor"},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"message":"User roles updated successfully for testuser"}`,
		},
		{
			name: "User Not Found",
			input: dtos.UpdateUserRolesDTO{
				Username: "nonexistent",
				Roles:    []string{"admin"},
			},
			mockError:      errors.New("user not found"),
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"user not found"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service
			mockService := new(mocks.MockUserService)
			mockService.On("UpdateRoles", mock.Anything, tc.input.Username, tc.input.Roles).Return(tc.mockError)

			// Create controller with mock
			userController := controllers.NewUserController(mockService, logger)

			// Setup HTTP context
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Create request with JSON body
			jsonInput, _ := json.Marshal(tc.input)
			req, _ := http.NewRequest("PUT", "/admin/update_role", bytes.NewBuffer(jsonInput))
			req.Header.Set("Content-Type", "application/json")
			c.Request = req

			// Call controller method
			userController.UpdateRoles(c, logger)

			// Assertions
			assert.Equal(t, tc.expectedStatus, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())

			// Verify mock expectations
			mockService.AssertExpectations(t)
		})
	}
}
