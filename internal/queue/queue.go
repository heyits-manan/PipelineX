package queue

import "context"

// TODO: Define queue abstractions here.
// TODO: Start with one job type for "video uploaded".
// TODO: Keep payloads explicit and serializable.

type VideoUploadedJob struct {
	VideoID string

	// TODO: Add fields workers need to process a video.
}

type Publisher interface {
	PublishVideoUploaded(ctx context.Context, job VideoUploadedJob) error
}

type Consumer interface {
	ConsumeVideoUploaded(ctx context.Context) (<-chan VideoUploadedJob, error)
}
