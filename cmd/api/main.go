package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/heyits-manan/PipelineX.git/internal/config"
	"github.com/heyits-manan/PipelineX.git/internal/queue"
	"github.com/heyits-manan/PipelineX.git/internal/store"
	"github.com/heyits-manan/PipelineX.git/internal/video"
	"github.com/redis/go-redis/v9"
)

type randomIDGenerator struct{}

func (g *randomIDGenerator) NewID() string {
	buf := make([]byte, 12)
	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}

	return hex.EncodeToString(buf)
}

func main() {
	cfg := config.MustLoad()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("invalid configuration: %v", err)
	}

	mux := http.NewServeMux()

	if err := runHTTP(context.Background(), cfg, mux); err != nil {
		log.Fatalf("failed to wire HTTP server: %v", err)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler: mux,
	}

	log.Printf("HTTP server started on port %d", cfg.HTTPPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

func runHTTP(ctx context.Context, cfg config.Config, mux *http.ServeMux) error {
	store := store.NewMemoryVideoStore()
	ids := &randomIDGenerator{}
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // replace with cfg later
	})
	if err := rdb.Ping(ctx).Err(); err != nil {
		return err
	}
	redisQueue := queue.NewRedisQueue(rdb)
	jobs := video.NewQueueJobPublisher(redisQueue)
	service := video.NewService(store, ids, jobs)
	handler := video.NewHandler(service)
	handler.RegisterRoutes(mux)
	return nil
}
