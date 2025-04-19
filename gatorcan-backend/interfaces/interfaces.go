package interfaces

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/models"
	"log"
	"net/http"
)

type CourseRepository interface {
	GetEnrolledCourses(ctx context.Context, userID int) ([]models.Enrollment, error)
	GetCourses(ctx context.Context, page, pageSize int) ([]models.Course, error)
	GetCourseByID(ctx context.Context, courseID int) (models.ActiveCourse, error)
	RequestEnrollment(ctx context.Context, userID, activeCourseID uint) error
	ApproveEnrollment(ctx context.Context, enrollmentID uint) error
	RejectEnrollment(ctx context.Context, enrollmentID uint) error
	GetPendingEnrollments(ctx context.Context) ([]models.Enrollment, error)
	GetCourseDetails(ctx context.Context, courseID uint) (models.Course, error)
}

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByUsernameorEmail(ctx context.Context, username string, email string) (*models.User, error)
	CreateNewUser(ctx context.Context, userDTO *dtos.UserCreateDTO) (*models.User, error)
	DeleteUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	UpdateUserRoles(ctx context.Context, user *models.User, roles []*models.Role) error
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type CourseService interface {
	GetEnrolledCourses(ctx context.Context, logger *log.Logger, username string) ([]dtos.EnrolledCoursesResponseDTO, error)
	GetCourses(ctx context.Context, logger *log.Logger, username string, page, pageSize int) ([]dtos.CourseResponseDTO, error)
	EnrollUser(ctx context.Context, logger *log.Logger, username string, courseID int) error
	GetCourseByID(ctx context.Context, logger *log.Logger, courseID int) (dtos.CourseResponseDTO, error)
}

type UserService interface {
	Login(ctx context.Context, loginData *dtos.LoginRequestDTO) (*dtos.LoginResponseDTO, error)
	GetUserDetails(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, username string, user *dtos.UpdateUserDTO) error
	CreateUser(ctx context.Context, logger *log.Logger, user *dtos.UserRequestDTO) (*dtos.UserResponseDTO, error)
	DeleteUser(ctx context.Context, username string) error
	UpdateRoles(ctx context.Context, username string, roles []string) error
}

type RoleRepository interface {
	GetRolesByName(ctx context.Context, roleNames []string) ([]models.Role, error)
}

type AssignmentRepository interface {
	GetAssignmentsByCourseID(ctx context.Context, courseID int) ([]models.Assignment, error)
	GetAssignmentByIDAndCourseID(ctx context.Context, assignmentID int, courseID int) (models.Assignment, error)
	UploadFileToAssignment(ctx context.Context, logger *log.Logger, username string, uploadData *dtos.UploadFileToAssignmentDTO) (*dtos.UploadFileToAssignmentResponseDTO, error)
	CreateAssignmentFile(ctx context.Context, assignmentFile *models.AssignmentFile) error
	LinkUserToAssignmentFile(ctx context.Context, userAssignmentFile *models.UserAssignmentFile) error
	UpsertAssignment(ctx context.Context, assignment *models.Assignment) error
}

type AssignmentService interface {
	GetAssignmentsByCourseID(ctx context.Context, courseID int) ([]dtos.AssignmentResponseDTO, error)
	GetAssignmentByIDAndCourseID(ctx context.Context, assignmentID int, courseID int) (dtos.AssignmentResponseDTO, error)
	UploadFileToAssignment(ctx context.Context, logger *log.Logger, username string, uploadData *dtos.UploadFileToAssignmentDTO) (*dtos.UploadFileToAssignmentResponseDTO, error)
	UpsertAssignment(ctx context.Context, logger *log.Logger, assignmentData *dtos.CreateOrUpdateAssignmentRequestDTO) (dtos.AssignmentResponseDTO, error)
}

type SubmissionRepository interface {
	GradeSubmission(ctx context.Context, assignmentID uint, courseID uint, userID uint, grade float64, feedback string) error
	GetSubmission(ctx context.Context, courseID int, assignmentID int, userID uint) (*models.Submission, error)
	GetGrades(ctx context.Context, courseID int, userID uint, count int) ([]dtos.GradeResponseDTO, error)
}

type SubmissionService interface {
	GradeSubmission(ctx context.Context, logger *log.Logger, username string, submissionData *dtos.GradeSubmissionRequestDTO) (*dtos.GradeSubmissionResponseDTO, error)
	GetSubmission(ctx context.Context, courseID int, assignmentID int, userID uint) (*dtos.SubmissionResponseDTO, error)
	GetGrades(ctx context.Context, logger *log.Logger, courseID int, userID uint) ([]dtos.GradeResponseDTO, error)
}

type AiServiceService interface {
	GetCourseRecommendations(ctx context.Context, logger *log.Logger, username string) ([]dtos.CourseRecommendationResponseDTO, error)
	GetTextSummary(ctx context.Context, logger *log.Logger, textSummaryRequest *dtos.TextSummaryRequestDTO) (*dtos.TextSummaryResponseDTO, error)
}

type AWSService interface {
	PushNotificationToSNS(ctx context.Context, logger *log.Logger, message string) error
}
