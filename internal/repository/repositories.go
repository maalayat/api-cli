package repository

import "github.com/maalayat/api-cli/internal/models"

type PokemonRepository interface {
	GetPokemons() (models.PokemonResult, error)
}

type StarWarsRepository interface {
	GetStarWars() (models.StarWarsResult, error)
}

type BreweryRepository interface {
	GetBreweries() ([]models.Brewery, error)
}
