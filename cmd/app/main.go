package main

import "github.com/vadimpk/go-gym-manager-api/internal/app"

const configPath = "configs/main"

// @title Gym Manager API
// @version 1.0
// @description REST API for Gym Management

// @securityDefinitions.apikey ManagerAuth
// @in header
// @name Authorization

// @host localhost:8000
// @BasePath /api/v1
func main() {
	app.Run(configPath)
}
