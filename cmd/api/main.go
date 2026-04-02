package main

import "fmt"
import "context"
import "log"
import "net/http"
import "github.com/heyits-manan/PipelineX.git/internal/config"
import "github.com/heyits-manan/PipelineX.git/internal/video"

func main(){
	config := config.MustLoad()	
	if err := config.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", config.HTTPPort),
		Handler: mux,
	}

	if err := runHTTP(context.Background(), config, mux); err != nil {
		log.Fatalf("Failed to start HTTP server: % v", err)
	}

	log.Printf("HTTP server started on port %d", config.HTTPPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

	log.Printf("HTTP server stopped")
	return nil

}


func runHTTP(ctx context.Context, config config.Config, mux *http.ServeMux) error {
	handler := video.NewHandler(video.NewService(config.StorageBackend, config.QueueBackend, config.DatabaseDSN))
	handler.RegisterRoutes(mux)
	return nil
}
