package video

import "context"
import "sync"
import "slices"
import "fmt"

type Store interface {
	Create(ctx context.Context, video Video) error
	GetByID(ctx context.Context, id string) (Video, error)
	UpdateStatus(ctx context.Context, input UpdateVideoStatusInput) error
	List(ctx context.Context) ([]Video, error)
}


type MemoryVideoStore struct {
	mu sync.RWMutex
	videos map[string]Video
}

var _ Store = (*MemoryVideoStore)(nil)

func NewMemoryVideoStore() *MemoryVideoStore {
	return &MemoryVideoStore{
		mu: sync.RWMutex{},
		videos: make(map[string]Video),
	}
}

func (s *MemoryVideoStore) Create(ctx context.Context, video Video) error{
	s.mu.Lock()
	defer s.mu.Unlock()
	s.videos[video.ID] = video
	return nil
}


func (s *MemoryVideoStore) GetByID(ctx context.Context, id string) (Video, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	video, ok := s.videos[id]
	if !ok {
		return Video{}, fmt.Errorf("video not found")
	}
	return video, nil
}


func (s *MemoryVideoStore) UpdateStatus(ctx context.Context, input *UpdateVideoStatusInput) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.videos[input.VideoID].Status = input.Status
	return nil
}

func (s *MemoryVideoStore) List(ctx context.Context) ([]Video, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out: = make([]Video, 0, len(s.videos))
	for _, v := range s.videos {
		out = append(out, v)
	}
	return out, nil
}


func (s *MemoryVideoStore) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.videos = nil
	return nil
}

