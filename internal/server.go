package ingestor

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/ScMofeoluwa/ingestor/internal/config"
)

type Server struct {
	router *chi.Mux
	config config.Config
}

func NewServer(config config.Config) *Server {
	return &Server{
		router: chi.NewRouter(),
		config: config,
	}
}

func (s *Server) Start() error {
	if err := s.migrateDB(); err != nil {
		log.Fatalf("error: %s\n", err)
	}
	s.setupRoutes()

	log.Printf("listening on port: %s\n", s.config.Port)
	http.ListenAndServe(":"+s.config.Port, s.router)
	return nil
}

func (s *Server) migrateDB() error {
	m, err := migrate.New("file://internal/database/migrations", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	fmt.Println("migrations successfully applied")
	return nil
}

func (s *Server) setupRoutes() {
	logService := NewLogService(s.config)
	logHandler := NewLogHandler(logService)
	s.router.Use(middleware.Logger)
	s.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	s.router.Mount("/debug", middleware.Profiler())
	s.router.Post("/ingest", logHandler.InsertLog)
}
