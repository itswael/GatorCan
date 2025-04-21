package repositories

import (
	"context"
	dtos "gatorcan-backend/DTOs"
	"gatorcan-backend/errors"
	"gatorcan-backend/models"
	"log"
	"time"

	"gorm.io/gorm"
)

type SubmissionRepository interface {
	GradeSubmission(ctx context.Context, assignmentID uint, courseID uint, userID uint, grade float64, feedback string) error
	GetSubmission(ctx context.Context, courseID int, assignmentID int, userID uint) (*models.Submission, error)
	GetSubmissions(ctx context.Context, courseID int, assignmentID int) ([]dtos.SubmissionsResponseDTO, error)
	GetGrades(ctx context.Context, courseID int, userID uint, count int) ([]dtos.GradeResponseDTO, error)
	Submit(ctx context.Context, userID uint, assignmentID uint, courseID uint) error
}

type submissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	return &submissionRepository{db: db}
}

func (s *submissionRepository) GradeSubmission(ctx context.Context, assignmentID uint, courseID uint, userID uint, grade float64, feedback string) error {
	var submission models.Submission
	if err := s.db.WithContext(ctx).
		Where("assignment_id = ? AND course_id = ? AND user_id = ?", assignmentID, courseID, userID).
		First(&submission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrSubmissionNotFound
		}
		log.Printf("Error fetching submission: %v", err)
		return errors.ErrDatabaseError
	}

	// Use Updates() to modify specific fields efficiently, including `updated_at`
	if err := s.db.WithContext(ctx).
		Model(&submission).
		Updates(map[string]interface{}{
			"grade":      int(grade),
			"feedback":   feedback,
			"updated_at": time.Now(), // Ensure updated_at is refreshed
		}).Error; err != nil {
		log.Printf("Error updating submission grade: %v", err)
		return errors.ErrGradingSubmissionFailed
	}

	return nil
}

// getsubmission
func (s *submissionRepository) GetSubmission(ctx context.Context, course_id int, assignmentID int, userID uint) (*models.Submission, error) {
	submission := models.Submission{}
	if err := s.db.WithContext(ctx).
		Where("assignment_id = ? AND course_id = ? AND user_id = ?", assignmentID, course_id, userID).
		First(&submission).Error; err != nil {
		return &models.Submission{}, errors.ErrSubmissionNotFound
	}
	return &submission, nil
}

func (s *submissionRepository) GetGrades(ctx context.Context, courseID int, userID uint, count int) ([]dtos.GradeResponseDTO, error) {
	var grades []dtos.GradeResponseDTO
	if err := s.db.WithContext(ctx).
		Raw(`
				SELECT 
					s.assignment_id, 
					a.title, 
					a.max_points, 
					a.deadline, 
					s.grade, 
					s.updated_at, 
					s.feedback,
					AVG(all_s.grade)/? AS mean,
					MIN(all_s.grade) AS min,
					MAX(all_s.grade) AS max
				FROM 
					assignments a 
					INNER JOIN submissions s ON a.id = s.assignment_id
					LEFT JOIN submissions all_s ON a.id = all_s.assignment_id AND all_s.course_id = s.course_id
				WHERE 
					s.user_id = ? 
					AND s.course_id = ?
				GROUP BY 
					s.assignment_id, a.title, a.max_points, a.deadline, s.grade, s.updated_at, s.feedback
			`, count, userID, courseID).Scan(&grades).Error; err != nil {
		log.Printf("Error fetching grades: %v", err)
		return nil, errors.ErrFetchingGrades
	}

	return grades, nil

}

func (s *submissionRepository) GetSubmissions(ctx context.Context, courseID int, assignmentID int) ([]dtos.SubmissionsResponseDTO, error) {
	var submissions []dtos.SubmissionsResponseDTO
	if err := s.db.WithContext(ctx).
		Raw(`
				SELECT 
					s.assignment_id, 
					s.user_id, 
					u.username, 
					s.grade, 
					s.feedback, 
					s.file_name, 
					s.file_type, 
					s.file_url, 
					s.updated_at
				FROM 
					submissions s 
					INNER JOIN users u ON s.user_id = u.id
				WHERE 
					s.assignment_id = ? 
					AND s.course_id = ?
			`, assignmentID, courseID).Scan(&submissions).Error; err != nil {
		log.Printf("Error fetching submissions: %v", err)
		return nil, errors.ErrFetchingSubmissions
	}

	return submissions, nil
}

func (s *submissionRepository) Submit(ctx context.Context, userID uint, assignmentID uint, courseID uint) error {
	// select a.file_name, a.file_url, a.file_type
	// from assignment_files a
	// where a.id in (select assignment_file_id from user_assignment_files where user_id = ?)
	type assignmentData struct {
		FileName string `json:"file_name"`
		FileURL  string `json:"file_url"`
		FileType string `json:"file_type"`
	}
	var data []assignmentData
	if err := s.db.WithContext(ctx).
		Raw(`
				SELECT 
					a.file_name, 
					a.file_url, 
					a.file_type
				FROM 
					assignment_files a 
					INNER JOIN user_assignment_files uaf ON a.id = uaf.assignment_file_id
				WHERE 
					uaf.user_id = ?
			`, userID).Scan(&data).Error; err != nil {
		log.Printf("Error submitting assignment: %v", err)
		return errors.ErrSubmittingAssignment
	}

	// Check if data is empty
	if len(data) == 0 {
		return errors.ErrAssignmentFileNotFound
	}

	// Save data to database
	for _, d := range data {
		submission := models.Submission{
			AssignmentID: assignmentID,
			CourseID:     courseID,
			UserID:       userID,
			File_name:    d.FileName,
			File_url:     d.FileURL,
			File_type:    d.FileType,
			Updated_at:   time.Now(),
		}
		if err := s.db.WithContext(ctx).Create(&submission).Error; err != nil {
			log.Printf("Error submitting assignment: %v", err)
			return errors.ErrSubmittingAssignment
		}
	}
	return nil
}
