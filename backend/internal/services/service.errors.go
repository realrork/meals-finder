package services

import "errors"

var (
	ErrUnauthorizedUser = errors.New("wrong login or password")
	ErrInternalFailure  = errors.New("internal failure")
)
