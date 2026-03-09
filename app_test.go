package main

import (
	"fmt"
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
