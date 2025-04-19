package errors

import "errors"

var (
	ErrAssignmentNotFound           = errors.New("assignment not found")
	ErrFailedToCreateAssignment     = errors.New("failed to create assignment")
	ErrFailedToUpdateAssignment     = errors.New("failed to update assignment")
	ErrFailedToDeleteAssignment     = errors.New("failed to delete assignment")
	ErrFailedToUploadFile           = errors.New("failed to upload file to assignment")
	ErrFailedToSubmitAssignment     = errors.New("failed to submit assignment")
	ErrUnauthorizedAccessAssignment = errors.New("unauthorized access to assignment")
	ErrInvalidAssignmentID          = errors.New("invalid assignment ID")
	ErrSubmissionAlreadyExists      = errors.New("submission already exists")
	ErrFileNotFoundAssignment       = errors.New("file not found for the assignment")
	ErrFailedToLinkFileToUser       = errors.New("failed to link file to user")
	ErrGradingSubmissionFailed      = errors.New("failed to grade submission")
)
