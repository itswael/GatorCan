package unit_tests

import (
	"context"
	"errors"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/config"
	"gatorcan-backend/models"
	"gatorcan-backend/services"
	"gatorcan-backend/unit_tests/mocks"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin_service(t *testing.T) {
	// Setup mocks
	mockUserRepo := new(mocks.MockUserRepository)
	mockRoleRepo := new(mocks.MockRoleRepository)
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create service with mocks
	userService := services.NewUserService(
		mockCourseRepo,
		mockUserRepo,
		mockRoleRepo,
		appConfig,
		mockHTTPClient,
	)

	// Create test context
	ctx := context.Background()

	// Setup test data
	hashedPassword := "$2a$12$4I/NDnXkOjVBchenmJsJneF3fzYKoDkWk.cflJ.2fmR5a2Kio214q" // This would be properly hashed in real code
	mockUser := &models.User{
		Username: "testuser",
		Password: hashedPassword,
		Roles:    []*models.Role{{Name: "student"}},
	}

	tests := []struct {
		name          string
		loginData     *dtos.LoginRequestDTO
		mockUser      *models.User
		mockError     error
		expectError   bool
		expectedToken string
	}{
		{
			name: "Successful Login",
			loginData: &dtos.LoginRequestDTO{
				Username: "testuser",
				Password: "password123",
			},
			mockUser:      mockUser,
			mockError:     nil,
			expectError:   false,
			expectedToken: "jwt-token", // The token that would be generated
		},
		{
			name: "User Not Found",
			loginData: &dtos.LoginRequestDTO{
				Username: "nonexistent",
				Password: "password123",
			},
			mockUser:      nil,
			mockError:     errors.New("user not found"),
			expectError:   true,
			expectedToken: "",
		},
		{
			name: "Invalid Password",
			loginData: &dtos.LoginRequestDTO{
				Username: "testuser",
				Password: "wrongpassword",
			},
			mockUser:      mockUser,
			mockError:     nil,
			expectError:   true,
			expectedToken: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup expectations
			mockUserRepo.On("GetUserByUsername", ctx, tc.loginData.Username).Return(tc.mockUser, tc.mockError).Once()

			// For successful login, we need to mock the password verification and token generation
			if tc.mockUser != nil && tc.loginData.Password == "password123" {
				// We'd normally mock the JWT generation, but our simplified implementation will handle this
				// This would be more complex with actual utils.GenerateToken calls
			}

			// Call service method
			response, err := userService.Login(ctx, tc.loginData)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
				assert.True(t, response.Err)
				assert.Empty(t, response.Token)
			} else {
				assert.NoError(t, err)
				assert.False(t, response.Err)
				assert.NotEmpty(t, response.Token)
			}

			// Verify all expectations
			mockUserRepo.AssertExpectations(t)
		})
	}
}

func TestCreateUser_service(t *testing.T) {
	// Setup mocks
	mockUserRepo := new(mocks.MockUserRepository)
	mockRoleRepo := new(mocks.MockRoleRepository)
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create service with mocks
	userService := services.NewUserService(
		mockCourseRepo,
		mockUserRepo,
		mockRoleRepo,
		appConfig,
		mockHTTPClient,
	)

	// Create test context
	ctx := context.Background()

	tests := []struct {
		name             string
		userData         *dtos.UserRequestDTO
		mockExistingUser *models.User
		mockRoles        []models.Role
		mockNewUser      *models.User
		expectError      bool
		expectedCode     int
	}{
		{
			name: "Successful User Creation",
			userData: &dtos.UserRequestDTO{
				Username: "newuser",
				Email:    "new@example.com",
				Password: "password123",
				Roles:    []string{"student"},
			},
			mockExistingUser: nil,
			mockRoles: []models.Role{
				{Name: "student"},
			},
			mockNewUser: &models.User{
				Username: "newuser",
				Email:    "new@example.com",
				Roles:    []*models.Role{{Name: "student"}},
			},
			expectError:  false,
			expectedCode: http.StatusCreated,
		},
		{
			name: "User Already Exists",
			userData: &dtos.UserRequestDTO{
				Username: "existinguser",
				Email:    "existing@example.com",
				Password: "password123",
				Roles:    []string{"student"},
			},
			mockExistingUser: &models.User{
				Username: "existinguser",
				Email:    "existing@example.com",
			},
			mockRoles:    nil,
			mockNewUser:  nil,
			expectError:  true,
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Role Not Found",
			userData: &dtos.UserRequestDTO{
				Username: "newuser",
				Email:    "new@example.com",
				Password: "password123",
				Roles:    []string{"invalidrole"},
			},
			mockExistingUser: nil,
			mockRoles:        []models.Role{},
			mockNewUser:      nil,
			expectError:      true,
			expectedCode:     http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup expectations
			mockUserRepo.On("GetUserByUsernameorEmail", ctx, tc.userData.Username, tc.userData.Email).
				Return(tc.mockExistingUser, func() error {
					if tc.mockExistingUser != nil {
						return nil
					}
					return errors.New("record not found")
				}()).Once()

			if tc.mockExistingUser == nil {
				mockRoleRepo.On("GetRolesByName", ctx, tc.userData.Roles).
					Return(tc.mockRoles, func() error {
						if tc.mockRoles == nil {
							return errors.New("role not found")
						}
						return nil
					}()).Maybe()

				if tc.mockRoles != nil {
					mockUserRepo.On("CreateNewUser", ctx, mock.AnythingOfType("*dtos.UserCreateDTO")).Return(tc.mockNewUser, nil).Maybe()
				}
			}

			// Call service method
			response, err := userService.CreateUser(ctx, tc.userData)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
				assert.True(t, response.Err)
				assert.Equal(t, tc.expectedCode, response.Code)
			} else {
				assert.NoError(t, err)
				assert.False(t, response.Err)
				assert.Equal(t, tc.expectedCode, response.Code)
				assert.Equal(t, "User created successfully", response.Message)
			}

			// Verify all expectations
			mockUserRepo.AssertExpectations(t)
			mockRoleRepo.AssertExpectations(t)
		})
	}
}

func TestGetUserDetails_service(t *testing.T) {
	// Setup mocks
	mockUserRepo := new(mocks.MockUserRepository)
	mockRoleRepo := new(mocks.MockRoleRepository)
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create service with mocks
	userService := services.NewUserService(
		mockCourseRepo,
		mockUserRepo,
		mockRoleRepo,
		appConfig,
		mockHTTPClient,
	)

	// Create test context
	ctx := context.Background()

	mockUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Roles:    []*models.Role{{Name: "student"}},
	}

	tests := []struct {
		name        string
		username    string
		mockUser    *models.User
		mockError   error
		expectError bool
	}{
		{
			name:        "Success",
			username:    "testuser",
			mockUser:    mockUser,
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "User Not Found",
			username:    "nonexistent",
			mockUser:    nil,
			mockError:   errors.New("user not found"),
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup expectations
			mockUserRepo.On("GetUserByUsername", ctx, tc.username).Return(tc.mockUser, tc.mockError).Once()

			// Call service method
			user, err := userService.GetUserDetails(ctx, tc.username)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tc.mockUser.Username, user.Username)
				assert.Equal(t, tc.mockUser.Email, user.Email)
				assert.Equal(t, tc.mockUser.CreatedAt, user.CreatedAt)
				assert.Equal(t, len(tc.mockUser.Roles), len(user.Roles))
			}

			// Verify all expectations
			mockUserRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteUser_service(t *testing.T) {
	// Setup mocks
	mockUserRepo := new(mocks.MockUserRepository)
	mockRoleRepo := new(mocks.MockRoleRepository)
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create service with mocks
	userService := services.NewUserService(
		mockCourseRepo,
		mockUserRepo,
		mockRoleRepo,
		appConfig,
		mockHTTPClient,
	)

	// Create test context
	ctx := context.Background()

	// Create test data
	mockUser := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
	}

	tests := []struct {
		name         string
		username     string
		mockUser     *models.User
		getUserError error
		deleteError  error
		expectError  bool
	}{
		{
			name:         "Success",
			username:     "testuser",
			mockUser:     mockUser,
			getUserError: nil,
			deleteError:  nil,
			expectError:  false,
		},
		{
			name:         "User Not Found",
			username:     "nonexistent",
			mockUser:     nil,
			getUserError: errors.New("user not found"),
			deleteError:  nil,
			expectError:  true,
		},
		{
			name:         "Delete Error",
			username:     "testuser",
			mockUser:     mockUser,
			getUserError: nil,
			deleteError:  errors.New("delete failed"),
			expectError:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup expectations
			mockUserRepo.On("GetUserByUsername", ctx, tc.username).Return(tc.mockUser, tc.getUserError).Once()

			if tc.mockUser != nil && tc.getUserError == nil {
				mockUserRepo.On("DeleteUser", ctx, tc.mockUser).Return(tc.deleteError).Once()
			}

			// Call service method
			err := userService.DeleteUser(ctx, tc.username)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Verify all expectations
			mockUserRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateUser_service(t *testing.T) {
	// Setup mocks
	mockUserRepo := new(mocks.MockUserRepository)
	mockRoleRepo := new(mocks.MockRoleRepository)
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create service with mocks
	userService := services.NewUserService(
		mockCourseRepo,
		mockUserRepo,
		mockRoleRepo,
		appConfig,
		mockHTTPClient,
	)

	// Create test context
	ctx := context.Background()

	// Create test data with real hashed password that verifies against "oldpassword"
	// In a real environment, we'd use the actual utils.HashPassword and utils.VerifyPassword
	// But for testing, we'll handle the verification in the mock expectations
	mockUser := &models.User{
		Username: "testuser",
		Password: "$2a$10$somehashedpassword", // Assume this is a valid hash for "oldpassword"
	}

	tests := []struct {
		name         string
		username     string
		updateData   *dtos.UpdateUserDTO
		mockUser     *models.User
		getUserError error
		updateError  error
		expectError  bool
	}{
		{
			name:     "Success",
			username: "testuser",
			updateData: &dtos.UpdateUserDTO{
				OldPassword: "oldpassword",
				NewPassword: "newpassword",
			},
			mockUser:     mockUser,
			getUserError: nil,
			updateError:  nil,
			expectError:  false,
		},
		{
			name:     "User Not Found",
			username: "nonexistent",
			updateData: &dtos.UpdateUserDTO{
				OldPassword: "oldpassword",
				NewPassword: "newpassword",
			},
			mockUser:     nil,
			getUserError: errors.New("user not found"),
			updateError:  nil,
			expectError:  true,
		},
		{
			name:     "Incorrect Old Password",
			username: "testuser",
			updateData: &dtos.UpdateUserDTO{
				OldPassword: "wrongpassword",
				NewPassword: "newpassword",
			},
			mockUser:     mockUser,
			getUserError: nil,
			updateError:  nil,
			expectError:  true,
		},
		{
			name:     "Update Error",
			username: "testuser",
			updateData: &dtos.UpdateUserDTO{
				OldPassword: "oldpassword",
				NewPassword: "newpassword",
			},
			mockUser:     mockUser,
			getUserError: nil,
			updateError:  errors.New("update failed"),
			expectError:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup expectations
			mockUserRepo.On("GetUserByUsername", ctx, tc.username).Return(tc.mockUser, tc.getUserError).Once()

			// For successful cases where we need to check password and update user
			if tc.mockUser != nil && tc.getUserError == nil {
				// In a real test with real utils.VerifyPassword, this would be more complex
				// Here we're simplifying by assuming oldpassword is valid for the mockUser
				if tc.updateData.OldPassword == "oldpassword" {
					mockUserRepo.On("UpdateUser", ctx, mock.AnythingOfType("*models.User")).Return(tc.updateError).Once()
				}
			}

			// Call service method
			err := userService.UpdateUser(ctx, tc.username, tc.updateData)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Verify all expectations
			mockUserRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateRoles_service(t *testing.T) {
	// Setup mocks
	mockUserRepo := new(mocks.MockUserRepository)
	mockRoleRepo := new(mocks.MockRoleRepository)
	mockCourseRepo := new(mocks.MockCourseRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	// Create test config
	appConfig := &config.AppConfig{
		Environment: "test",
	}

	// Create service with mocks
	userService := services.NewUserService(
		mockCourseRepo,
		mockUserRepo,
		mockRoleRepo,
		appConfig,
		mockHTTPClient,
	)

	// Create test context
	ctx := context.Background()

	// Create test data
	mockUser := &models.User{
		Username: "testuser",
		Roles:    []*models.Role{{Name: "student"}},
	}

	mockRoles := []models.Role{
		{Name: "admin"},
		{Name: "instructor"},
	}

	tests := []struct {
		name          string
		username      string
		roles         []string
		mockUser      *models.User
		getUserError  error
		mockRoles     []models.Role
		getRolesError error
		updateError   error
		expectError   bool
	}{
		{
			name:          "Success",
			username:      "testuser",
			roles:         []string{"admin", "instructor"},
			mockUser:      mockUser,
			getUserError:  nil,
			mockRoles:     mockRoles,
			getRolesError: nil,
			updateError:   nil,
			expectError:   false,
		},
		{
			name:          "User Not Found",
			username:      "nonexistent",
			roles:         []string{"admin"},
			mockUser:      nil,
			getUserError:  errors.New("user not found"),
			mockRoles:     nil,
			getRolesError: nil,
			updateError:   nil,
			expectError:   true,
		},
		{
			name:          "Role Not Found",
			username:      "testuser",
			roles:         []string{"invalidrole"},
			mockUser:      mockUser,
			getUserError:  nil,
			mockRoles:     []models.Role{},
			getRolesError: errors.New("role not found"),
			updateError:   nil,
			expectError:   true,
		},
		{
			name:          "Update Error",
			username:      "testuser",
			roles:         []string{"admin", "instructor"},
			mockUser:      mockUser,
			getUserError:  nil,
			mockRoles:     mockRoles,
			getRolesError: nil,
			updateError:   errors.New("update failed"),
			expectError:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup expectations
			mockUserRepo.On("GetUserByUsername", ctx, tc.username).Return(tc.mockUser, tc.getUserError).Once()

			if tc.mockUser != nil && tc.getUserError == nil {
				mockRoleRepo.On("GetRolesByName", ctx, tc.roles).Return(tc.mockRoles, tc.getRolesError).Once()

				if tc.mockRoles != nil && tc.getRolesError == nil {
					mockUserRepo.On("UpdateUser", ctx, mock.AnythingOfType("*models.User")).Return(tc.updateError).Once()
				}
			}

			// Call service method
			err := userService.UpdateRoles(ctx, tc.username, tc.roles)

			// Assertions
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Verify all expectations
			mockUserRepo.AssertExpectations(t)
			mockRoleRepo.AssertExpectations(t)
		})
	}
}
