package storage

import (
	"context"

	"github.com/garrettladley/roundest/internal/model"
	"github.com/garrettladley/roundest/internal/services/pokeapi"
	"github.com/garrettladley/roundest/internal/types"
)

type Storage interface {
	GetAllPokemon(ctx context.Context) ([]model.Pokemon, error)
	RandomPair(ctx context.Context) (types.Pair[model.Pokemon], error)

	GetAllResults(ctx context.Context) ([]model.Result, error)
	Vote(ctx context.Context, upvoteID int, downvoteID int) error

	Schema(ctx context.Context) error
	Seed(ctx context.Context, p []pokeapi.PokemonData) error
}
