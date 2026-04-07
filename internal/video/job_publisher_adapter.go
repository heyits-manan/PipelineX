package video

import (
	"context"

	"github.com/heyits-manan/PipelineX.git/internal/queue"
)

type QueueJobPublisher struct {
	publisher queue.Publisher
}

func NewQueueJobPublisher(publisher queue.Publisher) *QueueJobPublisher {
	return &QueueJobPublisher{publisher: publisher}
}

func (p *QueueJobPublisher) PublishVideoUploaded(ctx context.Context, v Video) error {
	return p.publisher.PublishVideoUploaded(ctx, queue.VideoUploadedJob{
		VideoID: v.ID,
	})
}
