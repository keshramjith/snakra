package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/vn:id", http.HandlerFunc(handleVoicenoteRead))
	mux.Get("/vn:user_id", http.HandlerFunc(handleVoicenotesByUserRead))
	mux.Post("/vn", http.HandlerFunc(handleVoicenoteCreate))
	return mux
}
