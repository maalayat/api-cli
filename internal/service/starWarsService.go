package service

import (
	"github.com/maalayat/api-cli/internal/models"
	"github.com/maalayat/api-cli/internal/repository"
)

type warsService struct {
	repository repository.StarWarsRepository
}

func StarWarsServiceImp(repository repository.StarWarsRepository) StarWarsService {
	return warsService{repository: repository}
}

func (service warsService) FetchStarWars() (models.StarWarsResult, error) {
	return service.repository.GetStarWars()
}
