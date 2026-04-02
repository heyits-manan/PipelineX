package processor

import "context"

// TODO: Use this interface to hide FFmpeg details from the rest of the application.
// TODO: Start with a fake implementation for learning and local development.

type Transcoder interface {
	Transcode(ctx context.Context, input TranscodeInput) (TranscodeOutput, error)
}

type TranscodeInput struct {
	VideoID string

	// TODO: Add source path, target profiles, output directory, thumbnail settings.
}

type TranscodeOutput struct {
	// TODO: Add output file paths, durations, thumbnails, manifests, and metadata.
}
