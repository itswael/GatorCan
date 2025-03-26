package dtos

import "time"

type UploadFileToAssignmentDTO struct {
	AssignmentID uint   `json:"assignment_id" binding:"required"`
	FileURL      string `json:"file_url" binding:"required"`
	FileName     string `json:"filename" binding:"required"`
	FileType     string `json:"file_type" binding:"required"`
	CourseID     uint   `json:"course_id" binding:"required"`
}

type UploadFileToAssignmentResponseDTO struct {
	AssignmentID uint      `json:"assignment_id"`
	FileID       uint      `json:"file_id"`
	FileName     string    `json:"filename"`
	FileType     string    `json:"file_type"`
	FileURL      string    `json:"file_url"`
	UploadedAt   time.Time `json:"uploaded_at"`
	UploaderID   uint      `json:"uploader_id"`
	CourseID     uint      `json:"course_id"`
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
