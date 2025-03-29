package dtos

// SubmissionResponseDTO represents the response structure for a submission.
type SubmissionResponseDTO struct {
	ID           uint   `json:"id"`
	AssignmentID uint   `json:"assignment_id"`
	CourseID     uint   `json:"course_id"`
	UploaderID   uint   `json:"uploader_id"`
	Grade        int    `json:"grade"`
	Feedback     string `json:"feedback"`
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

func NewSubmissionRequestDTO(assignmentID uint, fileURL, fileName, fileType string) *SubmissionRequestDTO {
	return &SubmissionRequestDTO{
		AssignmentID: assignmentID,
		FileURL:      fileURL,
		FileName:     fileName,
		FileType:     fileType,
	}
}
