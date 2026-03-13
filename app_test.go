package main

import (
	"errors"
	"fmt"
	"slices"
	"testing"
)

func TestGetPokemonId(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		want          int
		wantErr       error
		errorExpected bool
	}{
		{name: "Normal case", input: "pichu", want: 172, wantErr: nil, errorExpected: false},
		{name: "Non-existent Pokémon", input: "invalidpokemon", want: 0, wantErr: fmt.Errorf("failed to retrieve Pokémon info: 404 Not Found"), errorExpected: true},
		{name: "Empty Pokémon name", input: "", want: 0, wantErr: fmt.Errorf("failed to retrieve Pokémon info: 404 Not Found"), errorExpected: true},
		{name: "Pokémon name with multiple words", input: "mr-mime", want: 122, wantErr: nil, errorExpected: false},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPokemonInfo(tt.input)
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
	tests := []struct {
		name          string
		input         string
		want          int
		wantErr       error
		errorExpected bool
	}{
		{name: "Normal case", input: "charizard", want: 17, wantErr: nil, errorExpected: false},
		{name: "Non-existent Pokémon", input: "invalidpokemon", want: 0, wantErr: fmt.Errorf("failed to retrieve Pokémon info: 404 Not Found"), errorExpected: true},
		{name: "Empty Pokémon name", input: "", want: 0, wantErr: fmt.Errorf("failed to retrieve Pokémon info: 404 Not Found"), errorExpected: true},
		{name: "Pokémon name with multiple words", input: "dragonite", want: 22, wantErr: nil, errorExpected: false},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPokemonInfo(tt.input)
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
	tests := []struct {
		name          string
		input         string
		want          int
		wantErr       error
		errorExpected bool
	}{
		{name: "pikachu", input: "pikachu", want: 60, wantErr: nil, errorExpected: false},
		{name: "bulbasaur", input: "bulbasaur", want: 69, wantErr: nil, errorExpected: false},
		{name: "charmander", input: "charmander", want: 85, wantErr: nil, errorExpected: false},
		{name: "invalid pokemon", input: "non-existent-pokemon", want: 0, wantErr: errors.New("pokemon not found"), errorExpected: true},
		{name: "empty input", input: "", want: 0, wantErr: errors.New("invalid input"), errorExpected: true},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPokemonInfo(tt.input)
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

func TestGetPokemonTypes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{name: "normal case", input: "pichu", want: []string{"electric"}},
		{name: "normal case 2", input: "squirtle", want: []string{"water"}},
		{name: "dual types", input: "Swampert", want: []string{"water", "ground"}},
		{name: "dual types 2", input: "Starmie", want: []string{"water", "psychic"}},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := GetPokemonInfo(tt.input)
			if err != nil {
				panic(err)
			}

			got := GetPokemonTypes(input)

			correctTypes := 0
			for _, g := range got {
				if slices.Contains(tt.want, g) {
					correctTypes++
				}
			}
			if correctTypes != len(tt.want) {
				t.Errorf("input: %v - got: %v - want: %v", tt.input, got, tt.want)
			}
		})
	}
}
