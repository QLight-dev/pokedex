package main

import (
	"errors"
	"fmt"
	"testing"
)

// define tableTests struct to enforce DRY
type tableTests struct {
	name          string
	input         string
	want          int
	wantErr       error
	errorExpected bool
}

func TestGetPokemonId(t *testing.T) {
	tests := []tableTests{
		{name: "Normal case", input: "pichu", want: 172, wantErr: nil, errorExpected: false},
		{name: "Non-existent Pokémon", input: "invalidpokemon", want: 0, wantErr: fmt.Errorf("failed to retrieve Pokémon info: 404 Not Found"), errorExpected: true},
		{name: "Empty Pokémon name", input: "", want: 0, wantErr: fmt.Errorf("failed to retrieve Pokémon info: 404 Not Found"), errorExpected: true},
		{name: "Pokémon name with multiple words", input: "mr-mime", want: 122, wantErr: nil, errorExpected: false},
	}

	t.Parallel()
	for _, tt := range tests {
		// capture the variable to avoid modifying all test's tt variable when modifying it
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got, err := RequestPokemonInfo(tt.input)
			if !tt.errorExpected && err != tt.wantErr {
				t.Errorf("input: %v - got: %v, %v - want: %v, %v", tt.input, got, err, tt.want, tt.wantErr)
			} else if err == tt.wantErr && tt.wantErr != nil {
				t.SkipNow()
			}

			if got.Id != tt.want {
				t.Errorf("input: %v - got: %v, %v - want: %v, %v", tt.input, got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestGetPokemonHeight(t *testing.T) {
	tests := []tableTests{
		{name: "Normal case", input: "charizard", want: 17, wantErr: nil, errorExpected: false},
		{name: "Non-existent Pokémon", input: "invalidpokemon", want: 0, wantErr: fmt.Errorf("failed to retrieve Pokémon info: 404 Not Found"), errorExpected: true},
		{name: "Empty Pokémon name", input: "", want: 0, wantErr: fmt.Errorf("failed to retrieve Pokémon info: 404 Not Found"), errorExpected: true},
		{name: "Pokémon name with multiple words", input: "dragonite", want: 22, wantErr: nil, errorExpected: false},
	}

	t.Parallel()
	for _, tt := range tests {
		// capture the variable to avoid modifying all test's tt variable when modifying it
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got, err := RequestPokemonInfo(tt.input)
			if !tt.errorExpected && err != tt.wantErr {
				t.Errorf("input: %v - got: %v, %v - want: %v, %v", tt.input, got, err, tt.want, tt.wantErr)
			} else if err == tt.wantErr && tt.wantErr != nil {
				t.SkipNow()
			}

			if got.Height != tt.want {
				t.Errorf("input: %v - got: %v, %v - want: %v, %v", tt.input, got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestGetPokemonWeight(t *testing.T) {
	tests := []tableTests{
		{name: "pikachu", input: "pikachu", want: 60, wantErr: nil, errorExpected: false},
		{name: "bulbasaur", input: "bulbasaur", want: 69, wantErr: nil, errorExpected: false},
		{name: "charmander", input: "charmander", want: 85, wantErr: nil, errorExpected: false},
		{name: "invalid pokemon", input: "non-existent-pokemon", want: 0, wantErr: errors.New("pokemon not found"), errorExpected: true},
		{name: "empty input", input: "", want: 0, wantErr: errors.New("invalid input"), errorExpected: true},
	}

	t.Parallel()
	for _, tt := range tests {
		// capture the variable to avoid modifying all test's tt variable when modifying it
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			got, err := RequestPokemonInfo(tt.input)
			if !tt.errorExpected && err != tt.wantErr {
				t.Errorf("input: %v - got: %v, %v - want: %v, %v", tt.input, got, err, tt.want, tt.wantErr)
			} else if err == tt.wantErr && tt.wantErr != nil {
				t.SkipNow()
			}

			if got.Weight != tt.want {
				t.Errorf("input: %v - got: %v, %v - want: %v, %v", tt.input, got, err, tt.want, tt.wantErr)
			}
		})
	}
}
