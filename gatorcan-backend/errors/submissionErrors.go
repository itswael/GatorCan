package errors

import "errors"

var (
	ErrSubmissionNotFound                  = errors.New("submission not found")
	ErrFailedToCreateSubmission            = errors.New("failed to create submission")
	ErrFailedToUpdateSubmission            = errors.New("failed to update submission")
	ErrFailedToDeleteSubmission            = errors.New("failed to delete submission")
	ErrFailedToCreateGradeSubmission       = errors.New("failed to create grade submission")
	ErrFailedToUpdateGradeSubmission       = errors.New("failed to update grade submission")
	ErrFailedToDeleteGradeSubmission       = errors.New("failed to delete grade submission")
	ErrUnauthorizedAccessToGradeSubmission = errors.New("unauthorized access to grade submission")
	ErrInvalidGradeSubmissionID            = errors.New("invalid grade submission ID")
	ErrFileNotFoundForGradeSubmission      = errors.New("file not found for the grade submission")
	ErrFailedToLinkFileToGradeSubmission   = errors.New("failed to link file to grade submission")
)
