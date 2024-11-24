package postgres

import (
	"context"

	"github.com/garrettladley/roundest/internal/model"
)

func (db DB) GetAllPokemon(ctx context.Context) ([]model.Pokemon, error) {
	const query string = `SELECT id, name, dex_id, up_votes, down_votes, inserted_at, updated_at FROM pokemon WHERE id < 1025 ORDER BY up_votes DESC`
	var pokemon []model.Pokemon
	if err := db.db.SelectContext(ctx, &pokemon, query); err != nil {
		return nil, err
	}

	return pokemon, nil
}
