package service

import (
	"github.com/mayalaat/api-cli/internal/models"
	"github.com/mayalaat/api-cli/internal/repository"
)

type warService struct {
	starWarsRepository repository.StarWarsRepository
}

func StarWarsServiceImp(warsRepo repository.StarWarsRepository) StarWarsService {
	return warService{starWarsRepository: warsRepo}
}

func (service warService) FetchStarWars() (models.StarWarsResult, error) {
	return service.starWarsRepository.GetStarWars()
}
