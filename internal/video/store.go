package video

import (
	"context"
)

type Store interface {
	Create(ctx context.Context, video Video) error
	GetByID(ctx context.Context, id string) (Video, error)
	UpdateStatus(ctx context.Context, input UpdateVideoStatusInput) error
	List(ctx context.Context) ([]Video, error)
}
