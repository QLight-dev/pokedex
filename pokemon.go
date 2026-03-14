package main

type pokemon struct {
	Id     int    `json:"id"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Name   string `json:"name"`

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

	Stats []struct {
		Stat     int `json:"base_stat"`
		StatInfo struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
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

type BaseStats struct {
	HP      int
	Attack  int
	Defense int
	SpAtk   int
	SpDef   int
	Speed   int
}

// Extracts and returns a BaseStats struct from the given struct
func GetPokemonBaseStats(p pokemon) BaseStats {
	var result BaseStats

	for _, stat := range p.Stats {
		switch stat.StatInfo.Name {
		case "hp":
			result.HP = stat.Stat
		case "attack":
			result.Attack = stat.Stat
		case "defense":
			result.Defense = stat.Stat
		case "special-attack":
			result.SpAtk = stat.Stat
		case "special-defense":
			result.SpDef = stat.Stat
		case "speed":
			result.Speed = stat.Stat
		}
	}
	return result
}
