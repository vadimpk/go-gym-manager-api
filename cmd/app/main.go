package main

import "github.com/vadimpk/go-gym-manager-api/internal/app"

const configPath = "configs/main"

func main() {
	app.Run(configPath)
}
