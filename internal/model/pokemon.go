package model

import "time"

type Pokemon struct {
	ID         int64     `db:"id" json:"id"`
	Name       string    `db:"name" json:"name"`
	DexID      int       `db:"dex_id" json:"dexID"`
	UpVotes    int       `db:"up_votes" json:"upVotes"`
	DownVotes  int       `db:"down_votes" json:"downVotes"`
	InsertedAt time.Time `db:"inserted_at" json:"insertedAt"`
	UpdatedAt  time.Time `db:"updated_at" json:"updatedAt"`
}

type Result struct {
	Name           string  `db:"name" json:"name"`
	ID             int64   `db:"id" json:"id"`
	DexID          int     `db:"dex_id" json:"dexID"`
	UpVotes        int     `db:"up_votes" json:"upVotes"`
	DownVotes      int     `db:"down_votes" json:"downVotes"`
	TotalVotes     int     `db:"total_votes" json:"totalVotes"`
	WinPercentage  float64 `db:"win_percentage" json:"winPercentage"`
	LossPercentage float64 `db:"loss_percentage" json:"lossPercentage"`
}
