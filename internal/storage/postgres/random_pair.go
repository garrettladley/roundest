package postgres

import (
	"context"
	"fmt"

	"github.com/garrettladley/roundest/internal/model"
	"github.com/garrettladley/roundest/internal/types"
)

func (db DB) RandomPair(ctx context.Context) (types.Pair[model.Pokemon], error) {
	const query string = `SELECT name, id, dex_id, up_votes, down_votes, inserted_at, updated_at FROM pokemon WHERE id < 1025 ORDER BY RANDOM() LIMIT 2`
	var pokemon []model.Pokemon
	if err := db.db.SelectContext(ctx, &pokemon, query); err != nil {
		return types.Pair[model.Pokemon]{}, err
	}

	if len(pokemon) != 2 {
		return types.Pair[model.Pokemon]{}, fmt.Errorf("expected 2 pokemon, got %d", len(pokemon))
	}

	return types.Pair[model.Pokemon]{
		A: pokemon[0],
		B: pokemon[1],
	}, nil
}
