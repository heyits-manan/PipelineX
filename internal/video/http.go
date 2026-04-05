package video

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	service *Service
}

type CreateVideoRequest struct {
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	Size        int64  `json:"size"`
	StorageKey  string `json:"storage_key"`
}

type VideoResponse struct {
	ID           string `json:"id"`
	Filename     string `json:"filename"`
	ContentType  string `json:"content_type"`
	Size         int64  `json:"size"`
	StorageKey   string `json:"storage_key"`
	OutputKey    string `json:"output_key"`
	ThumbnailKey string `json:"thumbnail_key"`
	Status       string `json:"status"`
	Error        string `json:"error,omitempty"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /videos", h.CreateVideo)
	mux.HandleFunc("GET /videos", h.ListVideos)
	mux.HandleFunc("GET /videos/{id}", h.GetVideo)
}

func (h *Handler) CreateVideo(w http.ResponseWriter, r *http.Request) {
	var req CreateVideoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	video, err := h.service.CreateVideo(r.Context(), CreateVideoInput{
		Filename:    req.Filename,
		ContentType: req.ContentType,
		Size:        req.Size,
		StorageKey:  req.StorageKey,
	})
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.writeJSON(w, http.StatusCreated, toVideoResponse(video))
}

func (h *Handler) GetVideo(w http.ResponseWriter, r *http.Request) {
	video, err := h.service.GetVideo(r.Context(), r.PathValue("id"))
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	h.writeJSON(w, http.StatusOK, toVideoResponse(video))
}

func (h *Handler) ListVideos(w http.ResponseWriter, r *http.Request) {
	videos, err := h.service.ListVideos(r.Context())
	if err != nil {
		h.handleServiceError(w, err)
		return
	}

	response := make([]VideoResponse, 0, len(videos))
	for _, video := range videos {
		response = append(response, toVideoResponse(video))
	}

	h.writeJSON(w, http.StatusOK, response)
}

func (h *Handler) handleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, ErrInvalidInput):
		h.writeError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, ErrNotFound):
		h.writeError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, ErrConflict):
		h.writeError(w, http.StatusConflict, err.Error())
	default:
		h.writeError(w, http.StatusInternalServerError, "internal server error")
	}
}

func (h *Handler) writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func (h *Handler) writeError(w http.ResponseWriter, status int, message string) {
	h.writeJSON(w, status, errorResponse{Error: message})
}

func toVideoResponse(video Video) VideoResponse {
	return VideoResponse{
		ID:           video.ID,
		Filename:     video.Filename,
		ContentType:  video.ContentType,
		Size:         video.Size,
		StorageKey:   video.StorageKey,
		OutputKey:    video.OutputKey,
		ThumbnailKey: video.ThumbnailKey,
		Status:       string(video.Status),
		Error:        video.Error,
		CreatedAt:    video.CreatedAt.UTC().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:    video.UpdatedAt.UTC().Format("2006-01-02T15:04:05Z07:00"),
	}
}
