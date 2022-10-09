package util

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrGoalAlreadyExists  = errors.New("goal already exists")
	ErrGoalNotFound       = errors.New("goal not found")
	ErrUnableToCreate     = errors.New("unable to create")
	ErrInvalidAuthToken   = errors.New("invalid auth-token")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrNotImplemented     = errors.New("not implemented")
	ErrUnableToUpdate     = errors.New("unable to update")
	ErrNotFound           = errors.New("not found")
)
