package app

import (
	"github.com/vadimpk/go-gym-manager-api/internal/config"
	"log"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("Error when parsing config: %s", err.Error())
	}

	log.Println(cfg)
}
