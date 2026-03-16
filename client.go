package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type validParameters interface {
	~int | ~string
}
// GetPokemonInfo sends a GET request to the Pokémon API to retrieve information for a given Pokémon.
// It returns the retrieved Pokémon and an error, if any.
func RequestPokemonInfo[T validParameters](pokemonName T) (pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v/", pokemonName)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemon{}, err
	}

	client := http.Client{}
	res, err := client.Do(req)
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
