package dtos

import "time"

type UploadFileToAssignmentDTO struct {
	AssignmentID uint   `json:"assignment_id" binding:"required"`
	FileURL      string `json:"file_url" binding:"required"`
	FileName     string `json:"filename" binding:"required"`
	FileType     string `json:"file_type" binding:"required"`
	CourseID     uint   `json:"course_id" binding:"required"`
}

type UpdateAssignmentDTO struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Deadline    *string `json:"deadline,omitempty"`
	MaxPoints   *int    `json:"max_points,omitempty"`
}

type UploadFileToAssignmentResponseDTO struct {
	AssignmentID uint      `json:"assignment_id"`
	FileName     string    `json:"filename"`
	FileType     string    `json:"file_type"`
	FileURL      string    `json:"file_url"`
	UploadedAt   time.Time `json:"uploaded_at"`
	UploaderID   uint      `json:"uploader_id"`
	CourseID     uint      `json:"course_id"`
}

type GradeAssignmentDTO struct {
	AssignmentID uint    `json:"assignment_id" binding:"required"`
	UserID       uint    `json:"user_id" binding:"required"`
	Grade        float64 `json:"grade" binding:"required"`
	Feedback     string  `json:"feedback" binding:"required"`
}
