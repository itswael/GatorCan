package errors

import "errors"

var (
	ErrAssignmentNotFound      = errors.New("assignment not found")
	ErrFailedToCreate          = errors.New("failed to create assignment")
	ErrFailedToUpdate          = errors.New("failed to update assignment")
	ErrFailedToDelete          = errors.New("failed to delete assignment")
	ErrFailedToUploadFile      = errors.New("failed to upload file to assignment")
	ErrFailedToSubmit          = errors.New("failed to submit assignment")
	ErrUnauthorizedAccess      = errors.New("unauthorized access to assignment")
	ErrInvalidAssignmentID     = errors.New("invalid assignment ID")
	ErrSubmissionAlreadyExists = errors.New("submission already exists")
	ErrFileNotFound            = errors.New("file not found for the assignment")
	ErrFailedToLinkFileToUser  = errors.New("failed to link file to user")
)
