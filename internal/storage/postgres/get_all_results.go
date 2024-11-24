package postgres

import (
	"context"
	"math"
	"sort"

	"github.com/garrettladley/roundest/internal/model"
)

func (db DB) GetAllResults(ctx context.Context) ([]model.Result, error) {
	pokemon, err := db.GetAllPokemon(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]model.Result, len(pokemon))

	for i, pokemon := range pokemon {
		var (
			totalVotes     = pokemon.UpVotes + pokemon.DownVotes
			winPercentage  float64
			lossPercentage float64
		)

		if totalVotes > 0 {
			winPercentage = math.Round((float64(pokemon.UpVotes)/float64(totalVotes)*100)*100) / 100
			lossPercentage = math.Round((float64(pokemon.DownVotes)/float64(totalVotes)*100)*100) / 100
		}

		results[i] = model.Result{
			Name:           pokemon.Name,
			ID:             pokemon.ID,
			DexID:          pokemon.DexID,
			UpVotes:        pokemon.UpVotes,
			DownVotes:      pokemon.DownVotes,
			TotalVotes:     totalVotes,
			WinPercentage:  winPercentage,
			LossPercentage: lossPercentage,
		}
	}

	sort.Slice(results, func(i, j int) bool {
		// if win percentages are equal
		if results[i].WinPercentage == results[j].WinPercentage {
			return results[i].UpVotes > results[j].UpVotes
		}

		// sort by win percentage (higher percentage first)
		return results[i].WinPercentage > results[j].WinPercentage
	})

	return results, nil
}
