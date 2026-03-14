package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func PrintFormattedOutput(p pokemon) {
	// replace - to space to look more professional
	name := strings.ReplaceAll(p.Name, "-", " ")

	// upercase each word to make more professional
	caser := cases.Title(language.English)
	formattedName := caser.String(name)

	fmt.Printf("#%d %s\n", p.Id, formattedName)

	fmt.Printf("Type(s): \n")
	for _, t := range GetPokemonTypes(p) {
		fmt.Printf("	- %s\n", t)
	}

	// convert for human-friendly measurements
	// height = utils.DecimetersToMeters(float32(p.Height))
	// weight = utils.HectogramsToKg(float32(p.Weight))
}
