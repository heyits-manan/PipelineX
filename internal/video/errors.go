package video

import "errors"

// TODO: Keep shared domain errors here.
// TODO: Add more errors only when they support real control flow.

var (
	ErrNotFound     = errors.New("not found")
	ErrConflict     = errors.New("conflict")
	ErrInvalidInput = errors.New("invalid input")
)

