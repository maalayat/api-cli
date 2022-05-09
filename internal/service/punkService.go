package service

import (
	"github.com/mayalaat/api-cli/internal/models"
	"github.com/mayalaat/api-cli/internal/repository"
	"sync"
)

type punkService struct {
	repository repository.PunkRepository
}

func PunkServiceImp(repository repository.PunkRepository) PunkService {
	return punkService{repository: repository}
}

func (service punkService) FetchPunkBeers(wg *sync.WaitGroup) (p []models.Punk, e error) {
	p, e = service.repository.GetPunks()
	wg.Done()
	return
}
