package video

import "net/http"
import "encoding/json"
import "strings"
import "errors"

type Handler struct {
	service *Service
}

type CreateVideoRequest struct {
	// TODO: Add JSON fields for client input.
}

type VideoResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Filename string `json: "filename"`
	// TODO: Add fields that should be returned to clients.
}

type CreateVideoRequest struct {
	Filename string `json:"filename"`
}

func NewHandler(service *Service) *Handler{
	return &Handler{service: service}
}



// Suggested declarations to implement later:
// func NewHandler(service *Service) *Handler
// func (h *Handler) RegisterRoutes(mux *http.ServeMux)
// func (h *Handler) CreateVideo(w http.ResponseWriter, r *http.Request)
// func (h *Handler) GetVideo(w http.ResponseWriter, r *http.Request)
// func (h *Handler) ListVideos(w http.ResponseWriter, r *http.Request)
