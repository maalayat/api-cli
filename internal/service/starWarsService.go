package service

import (
	"github.com/mayalaat/api-cli/internal/models"
	"github.com/mayalaat/api-cli/internal/repository"
	"sync"
)

type warsService struct {
	repository repository.StarWarsRepository
}

func StarWarsServiceImp(repository repository.StarWarsRepository) StarWarsService {
	return warsService{repository: repository}
}

func (service warsService) FetchStarWars(wg *sync.WaitGroup) (s models.StarWarsResult, e error) {
	s, e = service.repository.GetStarWars()
	wg.Done()
	return
}
