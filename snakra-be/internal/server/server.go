package server

import (
	"log"
	"net/http"

	"github.com/go-chi/cors"
	"github.com/keshramjith/snakra/internal/dbservice"
	"github.com/keshramjith/snakra/internal/s3service"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router      *chi.Mux
	db          *dbservice.DbService
	s3client    *s3service.S3Client
	infoLogger  *log.Logger
	errorLogger *log.Logger
	driver      *http.Server
}

func NewServer(infoLogger, errorLogger *log.Logger, s3bn, addr string) *Server {
	mux := newMux()
	db := dbservice.NewDbConn()
	s3Client := s3service.NewS3Client(s3bn)
	drvSrv := &http.Server{Addr: addr, Handler: mux}
	srv := &Server{
		router:      mux,
		db:          db,
		s3client:    s3Client,
		driver:      drvSrv,
		infoLogger:  infoLogger,
		errorLogger: errorLogger,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) ListenAndServe() error {
	return s.driver.ListenAndServe()
}

func (s *Server) CloseDb() {
	s.db.CloseDb()
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
