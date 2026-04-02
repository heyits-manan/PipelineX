package video

import "context"

// TODO: Define the persistence contract for videos.
// TODO: Keep this interface small and tied to actual use cases.
// TODO: Implement an in-memory version first, then Postgres later.

type Store interface {
	Create(ctx context.Context, video Video) error
	GetByID(ctx context.Context, id string) (Video, error)
	UpdateStatus(ctx context.Context, input UpdateVideoStatusInput) error
	List(ctx context.Context) ([]Video, error)
}
