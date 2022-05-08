package service

import "github.com/mayalaat/api-cli/internal/models"

type PokemonService interface {
	FetchPokemons() (models.Data, error)
}
