package main

import (
	"context"

	"github.com/CharFranR/Hackaton2026/infrastructure/config"
	"github.com/CharFranR/Hackaton2026/infrastructure/database"
)

func main() {
	cfg := config.Load()
	database.MakeMigrations(context.Background(), cfg.DatabaseURL)
}
