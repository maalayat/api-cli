package service

import (
	"github.com/maalayat/api-cli/internal/models"
	"github.com/maalayat/api-cli/internal/repository"
)

type breweryService struct {
	repository repository.BreweryRepository
}

func BreweryServiceImp(repository repository.BreweryRepository) BreweryService {
	return breweryService{repository: repository}
}

func (service breweryService) FetchBreweries() ([]models.Brewery, error) {
	return service.repository.GetBreweries()
}
