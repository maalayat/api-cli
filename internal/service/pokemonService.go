package service

import (
	"github.com/mayalaat/api-cli/internal/models"
	"github.com/mayalaat/api-cli/internal/repository"
	"sync"
)

type pokeService struct {
	repository repository.PokemonRepository
}

func PokemonServiceImp(repository repository.PokemonRepository) PokemonService {
	return pokeService{repository: repository}
}

func (service pokeService) FetchPokemons(wg *sync.WaitGroup) (p models.PokemonResult, e error) {
	p, e = service.repository.GetPokemons()
	wg.Done()
	return
}
