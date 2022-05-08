package repository

import "github.com/mayalaat/api-cli/internal/models"

type PokemonRepository interface {
	GetPokemons() (models.Data, error)
}
