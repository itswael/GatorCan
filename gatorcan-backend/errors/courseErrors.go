package errors

import "errors"

var (
	ErrCourseNotFound  = errors.New("course not found")
	ErrAlreadyEnrolled = errors.New("enrollment request already exists")
	ErrCourseFull      = errors.New("course has reached maximum capacity")
	ErrFailedToEnroll  = errors.New("failed to request enrollment")
	ErrCourseInactive  = errors.New("course is not active")
	ErrFailedToFetch   = errors.New("failed to fetch courses")
)
