package video

// import "net/http"

// TODO: This file should contain HTTP transport code only.
// TODO: Parse requests, call service methods, and write JSON responses.
// TODO: Do not put domain logic directly in handlers.

type Handler struct {
	service *Service
}

type CreateVideoRequest struct {
	// TODO: Add JSON fields for client input.
}

type VideoResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`

	// TODO: Add fields that should be returned to clients.
}

// Suggested declarations to implement later:
// func NewHandler(service *Service) *Handler
// func (h *Handler) RegisterRoutes(mux *http.ServeMux)
// func (h *Handler) CreateVideo(w http.ResponseWriter, r *http.Request)
// func (h *Handler) GetVideo(w http.ResponseWriter, r *http.Request)
// func (h *Handler) ListVideos(w http.ResponseWriter, r *http.Request)
