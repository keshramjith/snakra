package server

import (
	"github.com/go-chi/cors"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/go-chi/chi/v5"
	s3config "github.com/keshramjith/snakra/internal/s3config"
)

type Server struct {
	router      *chi.Mux
	driver      *http.Server
	s3client    *s3.Client
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func NewServer() *Server {
	mux := newMux()
	s3Client := s3config.NewS3Client()
	addr := ":3001"
	drvSrv := &http.Server{Addr: addr, Handler: mux}
	srv := &Server{
		router:   mux,
		s3client: s3Client,
		driver:   drvSrv,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) ListenAndServe() error {
	return s.driver.ListenAndServe()
}

func newMux() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
	}))
	return mux
}
