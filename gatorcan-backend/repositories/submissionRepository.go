package repositories

import (
	"context"
	"gatorcan-backend/errors"
	"gatorcan-backend/models"
	"log"
	"time"

	"gorm.io/gorm"
)

type SubmissionRepository interface {
	GradeSubmission(ctx context.Context, assignmentID uint, courseID uint, userID uint, grade float64, feedback string) error
	GetSubmission(ctx context.Context, courseID int, assignmentID int, userID uint) (*models.Submission, error)
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
