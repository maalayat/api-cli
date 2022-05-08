package service

import (
	"github.com/mayalaat/api-cli/internal/models"
	"github.com/mayalaat/api-cli/internal/repository"
)

type pokeService struct {
	pokemonRepository repository.PokemonRepository
}

func PokemonServiceImp(pokeRepo repository.PokemonRepository) PokemonService {
	return pokeService{pokemonRepository: pokeRepo}
}

func (service pokeService) FetchPokemons() (models.PokemonResult, error) {
	return service.pokemonRepository.GetPokemons()
}
