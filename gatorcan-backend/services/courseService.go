package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/config"
	"gatorcan-backend/errors"
	"gatorcan-backend/interfaces"
	"gatorcan-backend/models"
	"log"
	"net/http"
	"time"
)

// type CourseRepository interface {
// 	GetCourses(page int, pageSize int) ([]models.Course, error)
// }

type CourseServiceImpl struct {
	courseRepo interfaces.CourseRepository
	userRepo   interfaces.UserRepository
	config     *config.AppConfig
	httpClient interfaces.HTTPClient
}

func NewCourseService(
	courseRepo interfaces.CourseRepository,
	userRepo interfaces.UserRepository,
	config *config.AppConfig,
	httpClient interfaces.HTTPClient,
) interfaces.CourseService {
	return &CourseServiceImpl{
		courseRepo: courseRepo,
		userRepo:   userRepo,
		config:     config,
		httpClient: httpClient,
	}
}

func (s *CourseServiceImpl) GetEnrolledCourses(ctx context.Context, logger *log.Logger, username string) ([]dtos.EnrolledCoursesResponseDTO, error) {

	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Printf("user not found: %s %d", username, 404)
		return nil, errors.ErrUserNotFound
	}

	enrollments, err := s.courseRepo.GetEnrolledCourses(ctx, int(user.ID))
	if err != nil {
		logger.Printf("failed to fetch enrolled courses: %s %d", username, 500)
		return nil, errors.ErrFailedToFetch
	}

	var enrolledCourses []dtos.EnrolledCoursesResponseDTO
	for _, enrollment := range enrollments {
		var course dtos.EnrolledCoursesResponseDTO
		course.ID = enrollment.ActiveCourse.Course.ID
		course.Name = enrollment.ActiveCourse.Course.Name
		course.Description = enrollment.ActiveCourse.Course.Description
		course.StartDate = enrollment.ActiveCourse.StartDate
		course.EndDate = enrollment.ActiveCourse.EndDate
		var instructor *models.User
		instructor, err = s.userRepo.GetUserByID(ctx, enrollment.ActiveCourse.InstructorID)
		if err != nil {
			return nil, errors.ErrFailedToFetch
		}
		course.InstructorName = instructor.Username
		course.InstructorEmail = instructor.Email
		enrolledCourses = append(enrolledCourses, course)
	}

	return enrolledCourses, nil
}

func (s *CourseServiceImpl) GetCourses(ctx context.Context, logger *log.Logger, username string, page int, pageSize int) ([]dtos.CourseResponseDTO, error) {
	_, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Printf("user not found: %s %d", username, 404)
		return nil, errors.ErrUserNotFound
	}

	// Fetch courses using pagination
	courses, err := s.courseRepo.GetCourses(ctx, page, pageSize)
	if err != nil {
		logger.Printf("Failed to fetch courses for page %d with pageSize %d: %v", page, pageSize, err)
		return nil, errors.ErrFailedToFetch
	}

	// Convert courses using DTO function
	return dtos.ConvertToCourseResponseDTOs(courses), nil
}

func (s *CourseServiceImpl) GetCourseByID(ctx context.Context, logger *log.Logger, courseID int) (dtos.CourseResponseDTO, error) {
	activeCourse, err := s.courseRepo.GetCourseByID(ctx, courseID)
	if err != nil {
		logger.Printf("course not found: %d %d", courseID, 404)
		return dtos.CourseResponseDTO{}, errors.ErrCourseNotFound
	}

	// get course details and instructor details
	course, err := s.courseRepo.GetCourseDetails(ctx, activeCourse.CourseID)
	if err != nil {
		logger.Printf("course not found: %d %d", courseID, 404)
		return dtos.CourseResponseDTO{}, errors.ErrCourseNotFound
	}

	instructor, err := s.userRepo.GetUserByID(ctx, activeCourse.InstructorID)
	if err != nil {
		logger.Printf("instructor not found: %d %d", activeCourse.InstructorID, 404)
		return dtos.CourseResponseDTO{}, errors.ErrUserNotFound
	}

	courseDTO := dtos.CourseResponseDTO{
		ID:              course.ID,
		Name:            course.Name,
		Description:     course.Description,
		InstructorName:  instructor.Username,
		InstructorEmail: instructor.Email,
	}

	return courseDTO, nil
}

func (s *CourseServiceImpl) EnrollUser(ctx context.Context, logger *log.Logger, username string, courseID int) error {

	user, err := s.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		logger.Printf("user not found: %s %d", username, 404)
		return errors.ErrUserNotFound
	}
	course, err := s.courseRepo.GetCourseByID(ctx, courseID)
	if err != nil {
		logger.Printf("course not found: %d %d", courseID, 404)
		return errors.ErrCourseNotFound
	}

	// Check if user is already enrolled
	enrollments, err := s.courseRepo.GetEnrolledCourses(ctx, int(user.ID))
	if err != nil {
		logger.Printf("failed to fetch enrolled courses: %s %d", username, 500)
		return errors.ErrFailedToFetch
	}
	for _, enrollment := range enrollments {
		if enrollment.ActiveCourse.CourseID == uint(courseID) {
			logger.Printf("user already enrolled: %s %d", username, 400)
			return errors.ErrAlreadyEnrolled
		}
	}

	// Check if the course is active
	if course.StartDate.After(time.Now()) {
		logger.Printf("course is not active: %d %d", courseID, 400)
		return errors.ErrCourseInactive
	}

	// Check if the course is full
	if course.Capacity == course.Enrolled {
		logger.Printf("course has reached maximum capacity: %d %d", courseID, 400)
		return errors.ErrCourseFull
	}

	err = s.courseRepo.RequestEnrollment(ctx, user.ID, course.ID)
	if err != nil {
		logger.Printf("failed to request enrollment: %s %d", username, 500)
		return errors.ErrFailedToEnroll
	}

	err = sendDiscordWebhook(user.ID, course.ID)
	if err != nil {
		logger.Printf("failed to send Discord webhook: %s %d", username, 500)
	}

	return nil
}

func sendDiscordWebhook(userID, courseID uint) error {
	const discordWebhookURL = "https://discord.com/api/webhooks/1345719796234453063/ToWh9shTfyqtSJtAwmgyz9rjw6W05E6pvvfMe5FqIql6v5T-hv0zIp3OUUQfMg62YcYw"
	const roleID = "<@&1345719467585310720>"
	message := fmt.Sprintf("%s A new enrollment request has been made!\nUser ID: `%d`\nCourse ID: `%d`\nRequesting permission to enroll.", roleID, userID, courseID)

	payload := map[string]string{
		"content": message,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", discordWebhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("discord webhook returned non-200 status: %d", resp.StatusCode)
	}

	return nil
}
