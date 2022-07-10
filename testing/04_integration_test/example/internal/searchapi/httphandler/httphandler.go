package httphandler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/steveanlorn/learning-go/testing/04_integration_test/example/internal/searchapi"
)

// Music ...
type Music struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

// Handler ...
type Handler struct {
	service searchapi.Service
}

// NewHandler ...
func NewHandler(service searchapi.Service) *Handler {
	handler := Handler{
		service: service,
	}

	return &handler
}

// ServeHTTP ...
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")

	musics, err := h.service.GetMusicByTitle(r.Context(), title)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	results := make([]Music, len(musics))
	for i, music := range musics {
		results[i] = Music{
			Id:    music.ID,
			Title: music.Title,
		}
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&results)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
}
