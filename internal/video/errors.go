package video

import "errors"

// TODO: Keep shared domain errors here.
// TODO: Add more errors only when they support real control flow.

var (
	ErrVideoNotFound = errors.New("video not found")
)
