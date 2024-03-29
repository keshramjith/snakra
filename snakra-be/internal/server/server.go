package server

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/go-chi/cors"
	"github.com/keshramjith/snakra/internal/dbservice"
	"github.com/keshramjith/snakra/internal/s3service"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router   *chi.Mux
	db       *dbservice.DbService
	s3client *s3service.S3Client
	logger   *zap.SugaredLogger
	driver   *http.Server
}

func NewServer(logger *zap.SugaredLogger, s3bn, addr string) *Server {
	mux := newMux()
	db := dbservice.NewDbConn()
	s3Client := s3service.NewS3Client(s3bn)
	drvSrv := &http.Server{Addr: addr, Handler: mux}
	srv := &Server{
		router:   mux,
		db:       db,
		s3client: s3Client,
		driver:   drvSrv,
		logger:   logger,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) ListenAndServe(env string) error {
	// if env == "dev" {
	// 	return s.driver.ListenAndServeTLS("../tls/cert.pem", "../tls/key.pem")
	// }
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
