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
}

type submissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	return &submissionRepository{db: db}
}

func (r *submissionRepository) GradeSubmission(ctx context.Context, assignmentID uint, courseID uint, userID uint, grade float64, feedback string) error {
	var submission models.Submission
	if err := r.db.WithContext(ctx).
		Where("assignment_id = ? AND course_id = ? AND user_id = ?", assignmentID, courseID, userID).
		First(&submission).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.ErrSubmissionNotFound
		}
		log.Printf("Error fetching submission: %v", err)
		return errors.ErrDatabaseError
	}

	// Use Updates() to modify specific fields efficiently, including `updated_at`
	if err := r.db.WithContext(ctx).
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
