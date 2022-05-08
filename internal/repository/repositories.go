package repository

import "github.com/mayalaat/api-cli/internal/models"

type PokemonRepository interface {
	GetPokemons() (models.PokemonResult, error)
}

type StarWarsRepository interface {
	GetStarWars() (models.StarWarsResult, error)
}

type PunkRepository interface {
	GetPunks() ([]models.Punk, error)
}
