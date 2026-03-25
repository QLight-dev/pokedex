# Pokedex


![Go Version](https://img.shields.io/badge/Go-1.26.1-blue?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge&logo=opensourceinitiative&logoColor=white)
![Repo Size](https://img.shields.io/github/repo-size/QLight-dev/pokedex?style=for-the-badge)
![Github Repo](https://img.shields.io/badge/github-repo-blue?style=for-the-badge&logo=github)

[demo](assets/pokedex_demo.png)

A pokedex CLI written in Go that uses the amazing [PokéAPI](https://pokeapi.co/).
## Installation
> [!NOTE]
> requires `go 1.26.1`
```zsh
go install github.com/QLight-dev/pokedex@latest
```
## Usage
### Show info for a specific Pokémon.
```zsh
pokedex pokemon <pokemon>
```
> **Example:**
> Show info for Arbok:
> ```zsh
> pokedex pokemon Arbok
> ```

> [!TIP]
> You can also pass in a Pokémon's ID.
### Show info for a random Pokémon.
```zsh
pokedex random
```
### Show help
```zsh
pokedex help
```

## License
MIT © Muhammed Uzair
