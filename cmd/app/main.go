package main

import "github.com/vadimpk/go-gym-manager-api/internal/app"

const configPath = "configs/main"

// @title Gym Manager API
// @version 1.0
// @description REST API for Gym Management

// @host localhost:8000
// @BasePath /
func main() {
	app.Run(configPath)
}
