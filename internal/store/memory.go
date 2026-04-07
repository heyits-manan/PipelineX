package store

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/heyits-manan/PipelineX.git/internal/video"
)

type MemoryVideoStore struct {
	mu     sync.RWMutex
	videos map[string]video.Video
}

var _ video.Store = (*MemoryVideoStore)(nil)

func NewMemoryVideoStore() *MemoryVideoStore {
	return &MemoryVideoStore{
		videos: make(map[string]video.Video),
	}
}

func (s *MemoryVideoStore) Create(ctx context.Context, v video.Video) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.videos[v.ID]; exists {
		return video.ErrConflict
	}

	if v.UpdatedAt.IsZero() {
		v.UpdatedAt = v.CreatedAt
	}

	s.videos[v.ID] = v
	return nil
}

func (s *MemoryVideoStore) GetByID(ctx context.Context, id string) (video.Video, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	v, ok := s.videos[id]
	if !ok {
		return video.Video{}, video.ErrNotFound
	}

	return v, nil
}

func (s *MemoryVideoStore) UpdateStatus(ctx context.Context, input video.UpdateVideoStatusInput) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	v, ok := s.videos[input.VideoID]
	if !ok {
		return video.ErrNotFound
	}

	v.Status = input.Status
	v.Error = input.Error
	v.OutputKey = input.OutputKey
	v.ThumbnailKey = input.ThumbnailKey
	if input.UpdatedAt.IsZero() {
		v.UpdatedAt = time.Now()
	} else {
		v.UpdatedAt = input.UpdatedAt
	}

	s.videos[input.VideoID] = v
	return nil
}

func (s *MemoryVideoStore) List(ctx context.Context) ([]video.Video, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]video.Video, 0, len(s.videos))
	for _, v := range s.videos {
		out = append(out, v)
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].CreatedAt.Equal(out[j].CreatedAt) {
			return out[i].ID < out[j].ID
		}
		return out[i].CreatedAt.Before(out[j].CreatedAt)
	})

	return out, nil
}
