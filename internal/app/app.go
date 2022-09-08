package app

import (
	"github.com/vadimpk/go-gym-manager-api/internal/config"
	"github.com/vadimpk/go-gym-manager-api/internal/delivery/http"
	"github.com/vadimpk/go-gym-manager-api/internal/server"
	"github.com/vadimpk/go-gym-manager-api/internal/service"
	"log"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("Error when parsing config: %s", err.Error())
	}

	services := service.NewServices()

	handler := http.NewHandler(services)

	srv := server.NewServer(cfg, handler.Init())

	if err := srv.Run(); err != nil {
		log.Fatalf("Error while running server: %s", err.Error())
	}
}
