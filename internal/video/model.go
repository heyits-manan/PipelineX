package video

import "time"

type Status string

const (
	StatusUploaded   Status = "uploaded"
	StatusProcessing Status = "processing"
	StatusReady      Status = "ready"
	StatusFailed     Status = "failed"
)

type Video struct {
	ID           string
	Filename     string
	ContentType  string
	Size         int64
	StorageKey   string
	OutputKey    string
	ThumbnailKey string
	Status       Status
	Error        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreateVideoInput struct {
	Filename    string
	ContentType string
	Size        int64
	StorageKey  string
}

type UpdateVideoStatusInput struct {
	VideoID      string
	Status       Status
	UpdatedAt    time.Time
	Error        string
	OutputKey    string
	ThumbnailKey string
}
