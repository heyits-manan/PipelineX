package video

import (
	"context"
	"time"
)

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

func NewService(store Store, ids IDGenerator, jobs JobPublisher) *Service {
	return &Service{
		store: store,
		ids:   ids,
		jobs:  jobs,
	}
}

func (s *Service) CreateVideo(ctx context.Context, input CreateVideoInput) (Video, error) {
	if input.Filename == "" {
		return Video{}, ErrInvalidInput
	}

	now := time.Now()
	video := Video{
		ID:          s.ids.NewID(),
		Filename:    input.Filename,
		ContentType: input.ContentType,
		Size:        input.Size,
		StorageKey:  input.StorageKey,
		Status:      StatusUploaded,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.store.Create(ctx, video); err != nil {
		return Video{}, err
	}

	if s.jobs != nil {
		if err := s.jobs.PublishVideoUploaded(ctx, video); err != nil {
			// Keep the stored video as uploaded and let the caller decide how to retry publishing.
			return Video{}, err
		}
	}

	return video, nil
}

func (s *Service) GetVideo(ctx context.Context, id string) (Video, error) {
	if id == "" {
		return Video{}, ErrInvalidInput
	}

	video, err := s.store.GetByID(ctx, id)
	if err != nil {
		return Video{}, err
	}

	return video, nil
}

func (s *Service) ListVideos(ctx context.Context) ([]Video, error) {
	videos, err := s.store.List(ctx)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (s *Service) MarkProcessing(ctx context.Context, videoID string) error {
	return s.updateStatus(ctx, UpdateVideoStatusInput{
		VideoID:   videoID,
		Status:    StatusProcessing,
		UpdatedAt: time.Now(),
	})
}

func (s *Service) MarkReady(ctx context.Context, input UpdateVideoStatusInput) error {
	input.Status = StatusReady
	input.UpdatedAt = time.Now()

	return s.updateStatus(ctx, input)
}

func (s *Service) MarkFailed(ctx context.Context, videoID string, processErr string) error {
	return s.updateStatus(ctx, UpdateVideoStatusInput{
		VideoID:   videoID,
		Status:    StatusFailed,
		UpdatedAt: time.Now(),
		Error:     processErr,
	})
}

func (s *Service) updateStatus(ctx context.Context, input UpdateVideoStatusInput) error {
	if input.VideoID == "" {
		return ErrInvalidInput
	}

	if input.Status == "" {
		return ErrInvalidInput
	}

	return s.store.UpdateStatus(ctx, input)
}
