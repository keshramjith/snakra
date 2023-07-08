package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/keshramjith/snakra/internal/handlers"
)

func New() *chi.Mux {
	mux := chi.NewRouter()

	mux.Get("/vn/{id}", http.HandlerFunc(handlers.HandleVoicenoteRead))
	mux.Post("/vn", http.HandlerFunc(handlers.HandleVoicenoteCreate))
	return mux
}
