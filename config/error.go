package config

import "errors"

var (
	ErrDataNotFound  = errors.New("the requested resource doesn't exists")
	ErrFailedCasting = errors.New("data failed to process")
)
