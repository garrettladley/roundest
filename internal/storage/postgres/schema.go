package postgres

import (
	"context"
	"fmt"
)

func (db DB) Schema(ctx context.Context) error {
	const schema string = `
    CREATE TABLE IF NOT EXISTS pokemon (
        id BIGINT PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        dex_id INTEGER NOT NULL,
        up_votes INTEGER DEFAULT 0,
        down_votes INTEGER DEFAULT 0,
        inserted_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

    -- Create an index on dex_id since it's likely to be queried
    CREATE INDEX IF NOT EXISTS idx_pokemon_dex_id ON pokemon(dex_id);`

	if _, err := db.db.ExecContext(ctx, schema); err != nil {
		return fmt.Errorf("failed to create pokemon schema: %w", err)
	}

	return nil
}
