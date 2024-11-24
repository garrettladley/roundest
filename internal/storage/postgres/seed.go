package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/garrettladley/roundest/internal/services/pokeapi"
)

func (db DB) Seed(ctx context.Context, p []pokeapi.PokemonData) error {
	if _, err := db.db.ExecContext(ctx, "DELETE FROM pokemon"); err != nil {
		return fmt.Errorf("error clearing existing pokemon: %w", err)
	}

	if len(p) == 0 {
		return nil
	}

	var (
		valueArgs = make([]interface{}, 0, len(p)*5)
		builder   strings.Builder
	)

	builder.Grow(len(`INSERT INTO pokemon (name, dex_id, up_votes, down_votes, id, inserted_at, updated_at) VALUES `) + (len(p) * 50))

	builder.WriteString(`INSERT INTO pokemon (
        name, dex_id, up_votes, down_votes, id, inserted_at, updated_at
    ) VALUES `)

	for i, pokemon := range p {
		if i > 0 {
			builder.WriteByte(',')
		}
		fmt.Fprintf(&builder, "($%d, $%d, $%d, $%d, $%d, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)",
			i*5+1, i*5+2, i*5+3, i*5+4, i*5+5)

		valueArgs = append(valueArgs,
			pokemon.Name,
			pokemon.ID,
			0, // up_votes
			0, // down_votes
			pokemon.ID)
	}

	if _, err := db.db.ExecContext(ctx, builder.String(), valueArgs...); err != nil {
		return fmt.Errorf("error performing bulk insert: %w", err)
	}

	return nil
}
