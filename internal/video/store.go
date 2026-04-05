package video

import (
	"context"
	"sort"
	"sync"
)

type Store interface {
	Create(ctx context.Context, video Video) error
	GetByID(ctx context.Context, id string) (Video, error)
	UpdateStatus(ctx context.Context, input UpdateVideoStatusInput) error
	List(ctx context.Context) ([]Video, error)
}

type MemoryVideoStore struct {
	mu     sync.RWMutex
	videos map[string]Video
}

var _ Store = (*MemoryVideoStore)(nil)

func NewMemoryVideoStore() *MemoryVideoStore {
	return &MemoryVideoStore{
		mu:     sync.RWMutex{},
		videos: make(map[string]Video),
	}
}

func (s *MemoryVideoStore) Create(ctx context.Context, video Video) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.videos[video.ID]; exists {
		return ErrConflict
	}

	if video.UpdatedAt.IsZero() {
		video.UpdatedAt = video.CreatedAt
	}

	s.videos[video.ID] = video
	return nil
}

func (s *MemoryVideoStore) GetByID(ctx context.Context, id string) (Video, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	video, ok := s.videos[id]
	if !ok {
		return Video{}, ErrNotFound
	}

	return video, nil
}

func (s *MemoryVideoStore) UpdateStatus(ctx context.Context, input UpdateVideoStatusInput) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	video, ok := s.videos[input.VideoID]
	if !ok {
		return ErrNotFound
	}

	video.Status = input.Status
	video.UpdatedAt = input.UpdatedAt
	video.Error = input.Error
	video.OutputKey = input.OutputKey
	video.ThumbnailKey = input.ThumbnailKey
	s.videos[input.VideoID] = video

	return nil
}

func (s *MemoryVideoStore) List(ctx context.Context) ([]Video, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	videos := make([]Video, 0, len(s.videos))
	for _, video := range s.videos {
		videos = append(videos, video)
	}

	sort.Slice(videos, func(i, j int) bool {
		if videos[i].CreatedAt.Equal(videos[j].CreatedAt) {
			return videos[i].ID < videos[j].ID
		}
		return videos[i].CreatedAt.Before(videos[j].CreatedAt)
	})

	return videos, nil
}

func (s *MemoryVideoStore) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.videos = nil
	return nil
}
