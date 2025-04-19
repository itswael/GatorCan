package dtos

import "time"

// SubmissionResponseDTO represents the response structure for a submission.
type SubmissionResponseDTO struct {
	Grade       int    `json:"grade"`
	Feedback    string `json:"feedback"`
	MaxPoints   int    `json:"max_points"`
	SubmittedAt string `json:"submitted_at"`
	FileURL     string `json:"file_url"`
}

// SubmissionRequestDTO represents the request structure for a submission.
type SubmissionRequestDTO struct {
	AssignmentID uint   `json:"assignment_id"`
	CourseID     uint   `json:"course_id"`
	UserID       uint   `json:"user_id"`
	FileURL      string `json:"file_url"`
	FileName     string `json:"file_name"`
	FileType     string `json:"file_type"`
}

type GradeResponseDTO struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Grade       int       `json:"grade"`
	MaxPoints   int       `json:"max_points"`
	SubmittedAt time.Time `json:"date_of_submission"`
	Feedback    string    `json:"feedback"`
	Deadline    time.Time `json:"deadline"`
	Mean        float64   `json:"mean"`
	Median      float64   `json:"median"`
	Highest     float64   `json:"highest"`
	Lowest      float64   `json:"lowest"`
}

func NewSubmissionRequestDTO(assignmentID uint, fileURL, fileName, fileType string) *SubmissionRequestDTO {
	return &SubmissionRequestDTO{
		AssignmentID: assignmentID,
		FileURL:      fileURL,
		FileName:     fileName,
		FileType:     fileType,
	}
}
