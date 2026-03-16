package main

import (
	"fmt"

	"github.com/QLight-dev/pokedex/utils"
)

func PrintFormattedOutput(p pokemon) {
	formattedName := utils.FormatName(p.Name)
	fmt.Printf("#%d %s\n", p.Id, formattedName)
	fmt.Println("")

	fmt.Printf("Type(s): \n")
	for _, t := range GetPokemonTypes(p) {
		fmt.Printf("	- %s\n", utils.FormatName(t))
	}
	fmt.Println("")

	// convert for human-friendly measurements
	height := utils.DecimetersToMeters(float32(p.Height))
	weight := utils.HectogramsToKg(float32(p.Weight))

	fmt.Printf("Height: %.1f m\n", height)
	fmt.Printf("Weight: %.1f kg\n", weight)
	fmt.Println("")

	fmt.Println("Abilities: ")
	for _, ability := range GetPokemonAbilities(p) {
		fmt.Printf("	- %s\n", utils.FormatName(ability))
	}
	fmt.Println("")

	stats := GetPokemonBaseStats(p)
	fmt.Println("Base Stats")
	fmt.Printf("HP: %v\n", stats.HP)
	fmt.Printf("Attack: %v\n", stats.Attack)
	fmt.Printf("Defense: %v\n", stats.Defense)
	fmt.Printf("Sp. Atk: %v\n", stats.SpAtk)
	fmt.Printf("Sp. Def: %v\n", stats.SpDef)
	fmt.Printf("Speed: %v\n", stats.Speed)
}
