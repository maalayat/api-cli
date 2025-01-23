package service

import (
	"github.com/maalayat/api-cli/internal/models"
	"github.com/maalayat/api-cli/internal/repository"
)

type pokeService struct {
	repository repository.PokemonRepository
}

func PokemonServiceImp(repository repository.PokemonRepository) PokemonService {
	return pokeService{repository: repository}
}

func (service pokeService) FetchPokemons() (models.PokemonResult, error) {
	return service.repository.GetPokemons()
}
