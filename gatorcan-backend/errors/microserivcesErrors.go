package errors

import "errors"

var (
	ErrMicroserviceNotFound = errors.New("microservice not found")
	ErrMicroserviceDown     = errors.New("microservice is down")
	ErrMicroserviceTimeout  = errors.New("microservice request timed out")
	ErrMicroserviceError    = errors.New("microservice returned an error")
)
