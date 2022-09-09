package app

import (
	_ "github.com/lib/pq"
	"github.com/vadimpk/go-gym-manager-api/internal/config"
	"github.com/vadimpk/go-gym-manager-api/internal/delivery/http"
	"github.com/vadimpk/go-gym-manager-api/internal/repository"
	"github.com/vadimpk/go-gym-manager-api/internal/repository/postgres"
	"github.com/vadimpk/go-gym-manager-api/internal/server"
	"github.com/vadimpk/go-gym-manager-api/internal/service"
	"github.com/vadimpk/go-gym-manager-api/pkg/auth"
	"log"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("Error when parsing config: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Error when connecting to db: %s", err.Error())
	}

	tokenManager, err := auth.NewManager(cfg.Auth.SigningKey)
	if err != nil {
		log.Fatalf("Error when creating token manager: %s", err.Error())
	}

	repo := repository.NewRepositories(db)

	services := service.NewServices(cfg, tokenManager, repo)

	handler := http.NewHandler(services, tokenManager)

	srv := server.NewServer(cfg, handler.Init())

	if err := srv.Run(); err != nil {
		log.Fatalf("Error while running server: %s", err.Error())
	}
}
