package video

import "context"
import "sync"
import "github.com/heyits-manan/PipelineX.git/internal/video/model"

// TODO: Define the persistence contract for videos.
// TODO: Keep this interface small and tied to actual use cases.
// TODO: Implement an in-memory version first, then Postgres later.

type Store interface {
	Create(ctx context.Context, video model.Video) error
	GetByID(ctx context.Context, id string) (model.Video, error)
	UpdateStatus(ctx context.Context, input model.UpdateVideoStatusInput) error
	List(ctx context.Context) ([]model.Video, error)
}


type MemoryVideoStore struct {
	mu sync.RWMutex
	videos map[string]model.Video
}

var _ Store = (*MemoryVideoStore)(nil)

func NewMemoryVideoStore() *MemoryVideoStore {
	return &MemoryVideoStore{
		mu: sync.RWMutex{},
		videos: make(map[string]model.Video),
	}
}

func (s *MemoryVideoStore) Create(ctx context.Context, video model.Video) error{
	s.mu.Lock()
	defer s.mu.Unlock()
	s.videos[video.ID] = video
	return nil
}