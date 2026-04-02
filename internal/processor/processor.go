package processor

import "context"

// TODO: This package should hold the background processing use cases.
// TODO: Start with fake processing before integrating FFmpeg.
// TODO: Keep external dependencies behind interfaces.

type Processor struct {
	// TODO: Add dependencies such as video service, storage, transcoder, and logger.
}

// Suggested declarations to implement later:
// func NewProcessor(...) *Processor
// func (p *Processor) ProcessVideoUploaded(ctx context.Context, videoID string) error
