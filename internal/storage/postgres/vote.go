package postgres

import (
	"context"
	"fmt"
	"log/slog"
)

type PokemonVotes struct {
	ID        int
	UpVotes   int
	DownVotes int
}

func (db DB) Vote(ctx context.Context, upvoteID int, downvoteID int) error {
	tx, err := db.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if r := recover(); r != nil {
			if err := tx.Rollback(); err != nil {
				slog.Error("rollback failed", slog.Any("error", err), slog.Any("recover", r))
			}
		}
	}()

	execQuery := func(query string, id int) error {
		if _, err := tx.Exec(query, id); err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				return fmt.Errorf("rollback failed: %v (original error: %v)", rbErr, err)
			}
			return fmt.Errorf("query failed: %w", err)
		}
		return nil
	}

	const upvoteQuery string = `UPDATE pokemon SET up_votes = up_votes + 1 WHERE id = $1`
	if err := execQuery(upvoteQuery, upvoteID); err != nil {
		return fmt.Errorf("upvote failed: %w", err)
	}

	const downvoteQuery string = `UPDATE pokemon SET down_votes = down_votes + 1 WHERE id = $1`
	if err := execQuery(downvoteQuery, downvoteID); err != nil {
		return fmt.Errorf("downvote failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
