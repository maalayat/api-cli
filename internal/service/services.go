package service

import (
	"github.com/mayalaat/api-cli/internal/models"
	"sync"
)

type PokemonService interface {
	FetchPokemons(*sync.WaitGroup) (models.PokemonResult, error)
}

type StarWarsService interface {
	FetchStarWars(*sync.WaitGroup) (models.StarWarsResult, error)
}

type PunkService interface {
	FetchPunkBeers(*sync.WaitGroup) ([]models.Punk, error)
}
