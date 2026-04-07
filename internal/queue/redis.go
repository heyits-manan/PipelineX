package queue

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/redis/go-redis/v9"
)

const videoUploadedStream = "video_uploaded"

type RedisQueue struct {
	client *redis.Client
}

func NewRedisQueue(client *redis.Client) *RedisQueue {
	return &RedisQueue{client: client}
}

func (q *RedisQueue) PublishVideoUploaded(ctx context.Context, job VideoUploadedJob) error {
	b, err := json.Marshal(job)
	if err != nil {
		return err
	}

	return q.client.XAdd(ctx, &redis.XAddArgs{
		Stream: videoUploadedStream,
		Values: map[string]any{"payload": string(b)},
	}).Err()
}

func (q *RedisQueue) ConsumeVideoUploaded(ctx context.Context) (<-chan VideoUploadedJob, error) {
	out := make(chan VideoUploadedJob)

	go func() {
		defer close(out)

		lastID := "$"
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			res, err := q.client.XRead(ctx, &redis.XReadArgs{
				Streams: []string{videoUploadedStream, lastID},
				Count:   1,
				Block:   0,
			}).Result()

			if err != nil {
				if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
					return
				}
				continue
			}

			for _, stream := range res {
				for _, msg := range stream.Messages {
					lastID = msg.ID

					raw, ok := msg.Values["payload"]
					if !ok {
						continue
					}
					payload, ok := raw.(string)
					if !ok {
						continue
					}

					var job VideoUploadedJob
					if err := json.Unmarshal([]byte(payload), &job); err != nil {
						continue
					}

					select {
					case out <- job:
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()

	return out, nil
}
