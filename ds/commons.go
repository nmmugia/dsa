package ds

import "errors"

var (
	ErrNoSuchElement = errors.New("ErrNoSuchElement")
	ErrCapacityLimit = errors.New("ErrCapacityLimit")
)
