package pokeapi

import (
	"context"
	"net/http"
)

func GetAllPokemon(ctx context.Context) ([]PokemonData, error) {
	return client.getAllPokemon(ctx)
}

var client = &apiClient{
	client:  &http.Client{},
	baseURL: "https://beta.pokeapi.co/graphql/v1beta",
}
