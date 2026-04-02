package video

import "context"

// TODO: Service should coordinate domain logic between handlers, store, and queue.
// TODO: Keep business rules here, not in HTTP handlers.
// TODO: Later, enqueue processing jobs after video creation.

type IDGenerator interface {
	NewID() string
}

type JobPublisher interface {
	PublishVideoUploaded(ctx context.Context, video Video) error
}

type Service struct {
	store Store
	ids   IDGenerator
	jobs  JobPublisher
}

// Suggested declarations to implement later:
// func NewService(store Store, ids IDGenerator, jobs JobPublisher) *Service
// func (s *Service) CreateVideo(ctx context.Context, input CreateVideoInput) (Video, error)
// func (s *Service) GetVideo(ctx context.Context, id string) (Video, error)
// func (s *Service) ListVideos(ctx context.Context) ([]Video, error)
// func (s *Service) MarkProcessing(ctx context.Context, input UpdateVideoStatusInput) error
// func (s *Service) MarkReady(ctx context.Context, input UpdateVideoStatusInput) error
// func (s *Service) MarkFailed(ctx context.Context, input UpdateVideoStatusInput) error
