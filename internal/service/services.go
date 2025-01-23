package service

import "github.com/maalayat/api-cli/internal/models"

type PokemonService interface {
	FetchPokemons() (models.PokemonResult, error)
}

type StarWarsService interface {
	FetchStarWars() (models.StarWarsResult, error)
}

type PunkService interface {
	FetchPunkBeers() ([]models.Punk, error)
}
