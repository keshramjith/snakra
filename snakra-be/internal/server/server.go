package server

import (
	"net/http"

	"github.com/keshramjith/snakra/internal/routes"
)

func NewServer() *http.Server {
	mux := routes.New()
	addr := ":3001"
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return srv
}
