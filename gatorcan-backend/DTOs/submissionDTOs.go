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
	AssignmentID uint      `gorm:"primaryKey" json:"assignment_id"`
	Title        string    `gorm:"not null" json:"title"`
	Grade        int       `json:"grade"`
	MaxPoints    int       `json:"max_points"`
	UpdatedAt    time.Time `json:"updated_at"`
	Feedback     string    `json:"feedback"`
	Deadline     time.Time `json:"deadline"`
	Max          int       `json:"max"`
	Min          int       `json:"min"`
	Mean         float64   `json:"mean"`
}

func NewSubmissionRequestDTO(assignmentID uint, fileURL, fileName, fileType string) *SubmissionRequestDTO {
	return &SubmissionRequestDTO{
		AssignmentID: assignmentID,
		FileURL:      fileURL,
		FileName:     fileName,
		FileType:     fileType,
	}
}
