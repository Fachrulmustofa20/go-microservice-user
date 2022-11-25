package constants

import (
	"errors"
)

var (
	AppPort = "APP_PORT"

	AppJSON = "application/json"

	ErrResponse = errors.New("token not found. please sign in first")
	ErrExpired  = errors.New("token expired. please login again")
	ErrLogin    = errors.New("invalid email/password")
)
