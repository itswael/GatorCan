package dtos

import "time"

type UploadFileToAssignmentDTO struct {
	AssignmentID uint   `json:"assignment_id" binding:"required"`
	FileURL      string `json:"file_url" binding:"required"`
}

type UploadFileToAssignmentResponseDTO struct {
	AssignmentID uint      `json:"assignment_id"`
	FileURL      string    `json:"file_url"`
	UploadedAt   time.Time `json:"uploaded_at"`
	UploaderID   uint      `json:"uploader_id"`
	UploaderName string    `json:"uploader_name"`
}

type SubmitAssignmentDTO struct {
	AssignmentID  uint   `json:"assignment_id" binding:"required"`
	UserID        uint   `json:"user_id" binding:"required"`
	SubmissionURL string `json:"submission_url" binding:"required"`
}

type GradeAssignmentDTO struct {
	AssignmentID uint    `json:"assignment_id" binding:"required"`
	UserID       uint    `json:"user_id" binding:"required"`
	Grade        float64 `json:"grade" binding:"required"`
	Feedback     string  `json:"feedback" binding:"required"`
}
