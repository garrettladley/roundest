package pokeapi

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	go_json "github.com/goccy/go-json"
)

type apiClient struct {
	client  *http.Client
	baseURL string
}

func (c *apiClient) getAllPokemon(ctx context.Context) ([]PokemonData, error) {
	jsonBody, err := go_json.Marshal(map[string]string{"query": getAllPokemonQuery})
	if err != nil {
		return nil, fmt.Errorf("error marshalling JSON: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	var graphQLResp getAllPokemonGraphQLResponse
	if err := go_json.NewDecoder(resp.Body).Decode(&graphQLResp); err != nil {
		return nil, fmt.Errorf("error parsing JSON response: %w", err)
	}

	pokemonList := make([]PokemonData, len(graphQLResp.Data.Pokemon))
	for i, p := range graphQLResp.Data.Pokemon {
		pokemonList[i] = PokemonData{
			ID:   p.ID,
			Name: p.PokemonSpecy.Name,
		}
	}

	return pokemonList, nil
}

const getAllPokemonQuery string = `
		query GetAllPokemon {
			pokemon_v2_pokemon {
				id
				pokemon_v2_pokemonspecy {
					name
				}
			}
		}`

type getAllPokemonGraphQLResponse struct {
	Data struct {
		Pokemon []struct {
			ID           int `json:"id"`
			PokemonSpecy struct {
				Name string `json:"name"`
			} `json:"pokemon_v2_pokemonspecy"`
		} `json:"pokemon_v2_pokemon"`
	} `json:"data"`
}
