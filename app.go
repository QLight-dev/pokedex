package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type pokemon struct {
	Id     int `json:"id"`
	Height int `json:"height"`
	Weight int `json:"weight"`

	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`

	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		}
	} `json:"abilities"`
}

// GetPokemonInfo sends a GET request to the Pokémon API to retrieve information for a given Pokémon.
// It returns the retrieved Pokémon and an error, if any.
func GetPokemonInfo(pokemonName string) (pokemon, error) {
	res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName))
	if err != nil {
		return pokemon{}, err
	}

	if res.StatusCode != http.StatusOK {
		return pokemon{}, fmt.Errorf("failed to retrieve Pokémon info: %s", res.Status)
	}

	defer res.Body.Close()

	var result pokemon
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&result); err != nil {
		return pokemon{}, err
	}

	return result, nil
}

// Extracts and returns a slice of types from the given struct
func GetPokemonTypes(p pokemon) []string {
	var result []string
	for _, t := range p.Types {
		result = append(result, t.Type.Name)
	}
	return result
}

// Extracts and returns a slice of abilities from the given struct
func GetPokemonAbilities(p pokemon) []string {
	var result []string
	for _, t := range p.Abilities {
		result = append(result, t.Ability.Name)
	}
	return result
}
