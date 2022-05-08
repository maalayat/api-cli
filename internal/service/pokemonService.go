package service

import (
	"github.com/mayalaat/api-cli/internal/models"
	"github.com/mayalaat/api-cli/internal/repository"
)

type service struct {
	pokemonRepository repository.PokemonRepository
}

func PokemonServiceImp(repo repository.PokemonRepository) PokemonService {
	return service{pokemonRepository: repo}
}

func (serv service) FetchPokemons() (models.Data, error) {
	return serv.pokemonRepository.GetPokemons()
}
