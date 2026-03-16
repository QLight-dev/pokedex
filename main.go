package main

import (
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		printHelp()
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "pokemon":
		if len(os.Args) < 3 {
			fmt.Println("Error: missing Pokémon name")
			fmt.Println()
			printHelp()
			os.Exit(1)
		}
		p, err := RequestPokemonInfo(os.Args[2])
		if err != nil {
			fmt.Println("Error fetching Pokémon:", err)
			os.Exit(1)
		}
		PrintFormattedOutput(p)

	case "random":
		id := rand.IntN(1028-1) + 1
		p, err := RequestPokemonInfo(id)
		if err != nil {
			fmt.Println("Error fetching Pokémon:", err)
			os.Exit(1)
		}
		PrintFormattedOutput(p)

	case "help":
		printHelp()

	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println()
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  pokedex pokemon <name>  - Show a specific Pokémon")
	fmt.Println("  pokedex random           - Show a random Pokémon")
	fmt.Println("  pokedex help             - Show this help message")
}
