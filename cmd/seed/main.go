package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/garrettladley/roundest/internal/services/pokeapi"
	"github.com/garrettladley/roundest/internal/settings"
	"github.com/garrettladley/roundest/internal/storage/postgres"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	settings, err := settings.Load()
	if err != nil {
		log.Fatalf("Failed to load settings: %v", err)
	}

	db, err := postgres.New(postgres.From(settings.Database))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	ctx := context.Background()

	if err := db.Schema(ctx); err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	pokemon, err := pokeapi.GetAllPokemon(ctx)
	if err != nil {
		log.Fatalf("Failed to get all pokemon: %v", err)
	}

	if err := db.Seed(ctx, pokemon); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	slog.Info("Database seeded")
}
