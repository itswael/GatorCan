package dtos

import "time"

type AssignmentRequestDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Deadline    string `json:"deadline" binding:"required"`
	MaxPoints   int    `json:"max_points"`
	FileURL     string `json:"url"`
	Grade       int    `json:"grade"`
	Feedback    string `json:"feedback"`
}

type AssignmentResponseDTO struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Deadline       time.Time `json:"deadline"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	ActiveCourseID uint      `json:"course_id"`
	MaxPoints      int       `json:"max_points"`
}

type CreateAssignmentRequestDTO struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	MaxPoints   int       `json:"max_points"`
	CourseID    uint      `json:"course_id"`
}

type UpdateAssignmentRequestDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	MaxPoints   int       `json:"max_points"`
	CourseID    uint      `json:"course_id"`
}
