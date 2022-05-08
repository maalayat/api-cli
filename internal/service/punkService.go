package service

import (
	"github.com/mayalaat/api-cli/internal/models"
	"github.com/mayalaat/api-cli/internal/repository"
)

type punkService struct {
	repository repository.PunkRepository
}

func PunkServiceImp(repository repository.PunkRepository) PunkService {
	return punkService{repository: repository}
}

func (service punkService) FetchPunkBeers() ([]models.Punk, error) {
	return service.repository.GetPunks()
}
