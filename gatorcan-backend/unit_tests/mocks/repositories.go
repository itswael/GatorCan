package mocks

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/models"
	"log"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockRoleRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) GetRolesByName(ctx context.Context, roles []string) ([]models.Role, error) {
	args := m.Called(ctx, roles)
	return args.Get(0).([]models.Role), args.Error(1)
}

// MockCourseRepository is a mock implementation of the course repository
type MockCourseRepository struct {
	mock.Mock
}

func (m *MockCourseRepository) GetEnrolledCourses(ctx context.Context, userID int) ([]models.Enrollment, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Enrollment), args.Error(1)
}

func (m *MockCourseRepository) GetCourses(ctx context.Context, page, pageSize int) ([]models.Course, error) {
	args := m.Called(ctx, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Course), args.Error(1)
}

func (m *MockCourseRepository) GetCourseByID(ctx context.Context, courseID int) (models.ActiveCourse, error) {
	args := m.Called(ctx, courseID)
	return args.Get(0).(models.ActiveCourse), args.Error(1)
}

func (m *MockCourseRepository) RequestEnrollment(ctx context.Context, userID, activeCourseID uint) error {
	args := m.Called(ctx, userID, activeCourseID)
	return args.Error(0)
}

func (m *MockCourseRepository) ApproveEnrollment(ctx context.Context, enrollmentID uint) error {
	args := m.Called(ctx, enrollmentID)
	return args.Error(0)
}

func (m *MockCourseRepository) RejectEnrollment(ctx context.Context, enrollmentID uint) error {
	args := m.Called(ctx, enrollmentID)
	return args.Error(0)
}

func (m *MockCourseRepository) GetPendingEnrollments(ctx context.Context) ([]models.Enrollment, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Enrollment), args.Error(1)
}

func (m *MockCourseRepository) GetCourseDetails(ctx context.Context, courseID uint) (models.Course, error) {
	args := m.Called(ctx, courseID)
	return args.Get(0).(models.Course), args.Error(1)
}

// MockUserRepository mocks the user repository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	args := m.Called(ctx, username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUserByUsernameorEmail(ctx context.Context, username, email string) (*models.User, error) {
	args := m.Called(ctx, username, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) CreateNewUser(ctx context.Context, userDTO *dtos.UserCreateDTO) (*models.User, error) {
	args := m.Called(ctx, userDTO)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdatePassword(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateEmail(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) UpdateUserRoles(ctx context.Context, user *models.User, roles []*models.Role) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// MockHTTPClient mocks the HTTP client
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*http.Response), args.Error(1)
}

// MockAssignmentRepository mocks the assignment repository
type MockAssignmentRepository struct {
	mock.Mock
}

func (m *MockAssignmentRepository) GetAssignmentsByCourseID(ctx context.Context, courseID int) ([]models.Assignment, error) {
	args := m.Called(ctx, courseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Assignment), args.Error(1)
}

func (m *MockAssignmentRepository) GetAssignmentByIDAndCourseID(ctx context.Context, assignmentID int, courseID int) (models.Assignment, error) {
	args := m.Called(ctx, assignmentID, courseID)
	return args.Get(0).(models.Assignment), args.Error(1)
}

func (m *MockAssignmentRepository) CreateAssignmentFile(ctx context.Context, file *models.AssignmentFile) error {
	args := m.Called(ctx, file)
	// Set the ID on the file to simulate DB creation
	if args.Error(0) == nil && file.ID == 0 {
		file.ID = 1
	}
	return args.Error(0)
}

func (m *MockAssignmentRepository) LinkUserToAssignmentFile(ctx context.Context, userFile *models.UserAssignmentFile) error {
	args := m.Called(ctx, userFile)
	// Set the ID on the userFile to simulate DB creation
	if args.Error(0) == nil && userFile.ID == 0 {
		userFile.ID = 1
	}
	return args.Error(0)
}

func (m *MockAssignmentRepository) UploadFileToAssignment(ctx context.Context, logger *log.Logger, username string, uploadData *dtos.UploadFileToAssignmentDTO) (*dtos.UploadFileToAssignmentResponseDTO, error) {
	args := m.Called(ctx, logger, username, uploadData)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dtos.UploadFileToAssignmentResponseDTO), args.Error(1)
}

func (m *MockAssignmentRepository) UpsertAssignment(ctx context.Context, assignment *models.Assignment) error {
	args := m.Called(ctx, assignment)
	// Set the ID on the assignment to simulate DB creation
	if args.Error(0) == nil && assignment.ID == 0 {
		assignment.ID = 1
	}
	return args.Error(0)
}

// MockSubmissionRepository mocks the SubmissionRepository interface
type MockSubmissionRepository struct {
	mock.Mock
}

func (m *MockSubmissionRepository) GetSubmission(ctx context.Context, courseID, assignmentID int, userID uint) (*models.Submission, error) {
	args := m.Called(ctx, courseID, assignmentID, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Submission), args.Error(1)
}

func (m *MockSubmissionRepository) GradeSubmission(ctx context.Context, assignmentID uint, courseID uint, userID uint, grade float64, feedback string) error {
	args := m.Called(ctx, assignmentID, courseID, userID, grade, feedback)
	return args.Error(0)
}
