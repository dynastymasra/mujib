package config

import "errors"

var (
	ErrDataNotFound  = errors.New("the requested resource doesn't exists")
	ErrNotAuthorized = errors.New("don't have access")
)
