package video

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeIDGenerator struct {
	n int
}

func (g *fakeIDGenerator) NewID() string {
	g.n++
	return fmt.Sprintf("id-%d", g.n)
}

func setupHTTPHandler(t *testing.T) *http.ServeMux {
	t.Helper()

	store := NewMemoryVideoStore()
	ids := &fakeIDGenerator{}
	service := NewService(store, ids, nil)
	handler := NewHandler(service)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)
	return mux
}

func TestCreateVideo_Success(t *testing.T) {
	mux := setupHTTPHandler(t)

	body := `{
		"filename":"clip.mp4",
		"content_type":"video/mp4",
		"size":123,
		"storage_key":"uploads/clip.mp4"
	}`

	req := httptest.NewRequest(http.MethodPost, "/videos", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d, body=%s", http.StatusCreated, rec.Code, rec.Body.String())
	}

	var got VideoResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if got.ID == "" {
		t.Fatalf("expected non-empty ID")
	}
	if got.Filename != "clip.mp4" {
		t.Fatalf("expected filename clip.mp4, got %q", got.Filename)
	}
	if got.Status != string(StatusUploaded) {
		t.Fatalf("expected status %q, got %q", StatusUploaded, got.Status)
	}
}

func TestCreateVideo_InvalidJSON(t *testing.T) {
	mux := setupHTTPHandler(t)

	req := httptest.NewRequest(http.MethodPost, "/videos", bytes.NewBufferString(`{"filename":}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d, body=%s", http.StatusBadRequest, rec.Code, rec.Body.String())
	}
}

func TestGetVideo_NotFound(t *testing.T) {
	mux := setupHTTPHandler(t)

	req := httptest.NewRequest(http.MethodGet, "/videos/does-not-exist", nil)
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d, body=%s", http.StatusNotFound, rec.Code, rec.Body.String())
	}
}

func TestListVideos_Success(t *testing.T) {
	mux := setupHTTPHandler(t)

	// Seed by hitting the create endpoint twice.
	create := func(filename string) {
		payload, _ := json.Marshal(CreateVideoRequest{
			Filename: filename,
		})
		req := httptest.NewRequest(http.MethodPost, "/videos", bytes.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		if rec.Code != http.StatusCreated {
			t.Fatalf("create failed: status=%d body=%s", rec.Code, rec.Body.String())
		}
	}

	create("a.mp4")
	create("b.mp4")

	req := httptest.NewRequest(http.MethodGet, "/videos", nil).WithContext(context.Background())
	rec := httptest.NewRecorder()

	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d, body=%s", http.StatusOK, rec.Code, rec.Body.String())
	}

	var got []VideoResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("failed to unmarshal list response: %v", err)
	}

	if len(got) != 2 {
		t.Fatalf("expected 2 videos, got %d", len(got))
	}
}
