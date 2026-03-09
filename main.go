package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type pokemon struct {
	Id int `json:"id"`
}

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

func main() {

}
