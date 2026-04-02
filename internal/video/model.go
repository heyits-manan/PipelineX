package video

import "time"

// TODO: Keep domain models in this file.
// TODO: Start with only the fields required for create/get video flows.
// TODO: Add processing-specific fields later when you implement workers and FFmpeg.

type Status string

const (
	// TODO: Add or rename statuses as the project evolves.
	StatusUploaded   Status = "uploaded"
	StatusProcessing Status = "processing"
	StatusReady      Status = "ready"
	StatusFailed     Status = "failed"
)

type Video struct {
	ID string

	// TODO: Add metadata fields such as filename, content type, size, storage key.
	Status Status
	Filename string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateVideoInput struct {
	// TODO: Add fields required to create/register a new video.
	Filename string

}

type UpdateVideoStatusInput struct {
	VideoID string
	Status  Status

	// TODO: Add optional error details, output locations, thumbnail paths, etc.
}
