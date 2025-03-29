package dtos

// SubmissionResponseDTO represents the response structure for a submission.
type GradeSubmissionResponseDTO struct {
	AssignmentID uint   `json:"assignment_id"`
	CourseID     uint   `json:"course_id"`
	UserID       uint   `json:"user_id"`
	Grade        int    `json:"grade"`
	Feedback     string `json:"feedback"`
}

// GradeSubmissionRequestDTO represents the request structure for a submission.
type GradeSubmissionRequestDTO struct {
	AssignmentID uint   `json:"assignment_id"`
	CourseID     uint   `json:"course_id"`
	UserID       uint   `json:"user_id"`
	Grade        int    `json:"grade"`
	Feedback     string `json:"feedback"`
}

func NewGradeSubmissionRequestDTO(assignmentID, courseID, userID uint, grade int, feedback string) *GradeSubmissionRequestDTO {
	return &GradeSubmissionRequestDTO{
		AssignmentID: assignmentID,
		CourseID:     courseID,
		UserID:       userID,
		Grade:        grade,
		Feedback:     feedback,
	}
}
