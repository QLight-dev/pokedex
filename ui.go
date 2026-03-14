package main

import (
	"fmt"

	"github.com/QLight-dev/pokedex/utils"
)

func PrintFormattedOutput(p pokemon) {
	fmt.Printf("#%d %s\n", p.Id, p.Name)

	fmt.Printf("Type(s): \n")
	for _, t := range GetPokemonTypes(p) {
		fmt.Printf("	- %s\n", t)
	}

	// convert for human-friendly measurements
	height = utils.DecimetersToMeters(float32(p.Height))
	weight = utils.HectogramsToKg(float32(p.Weight))
}
