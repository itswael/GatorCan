package errors

import "errors"

var (
	ErrSNSNotificationFailed = errors.New("failed to send SNS notification")
)
