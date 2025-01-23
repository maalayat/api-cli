package service

import (
	"github.com/maalayat/api-cli/internal/models"
	"github.com/maalayat/api-cli/internal/repository"
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
