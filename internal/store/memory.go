package store

import "github.com/heyits-manan/PipelineX.git/internal/video"

// TODO: Implement the first concrete video.Store here using an in-memory map.
// TODO: Add synchronization with sync.RWMutex when you implement it.
// TODO: This package should satisfy the video.Store interface.

type MemoryVideoStore struct {
	// TODO: Add mutex and map fields.
}

var _ video.Store = (*MemoryVideoStore)(nil)

// Suggested declarations to implement later:
// func NewMemoryVideoStore() *MemoryVideoStore
// func (s *MemoryVideoStore) Create(ctx context.Context, v video.Video) error
// func (s *MemoryVideoStore) GetByID(ctx context.Context, id string) (video.Video, error)
// func (s *MemoryVideoStore) UpdateStatus(ctx context.Context, input video.UpdateVideoStatusInput) error
// func (s *MemoryVideoStore) List(ctx context.Context) ([]video.Video, error)
