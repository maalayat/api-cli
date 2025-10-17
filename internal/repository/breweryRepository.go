package repository

import (
	"encoding/json"
	"fmt"
	"github.com/maalayat/api-cli/internal/models"
	"io"
	"net/http"
)

const (
	breweryUrl = "https://api.openbrewerydb.org/v1/breweries"
)

type breweryRepo struct {
	url string
}

func BreweryRepositoryImp() BreweryRepository {
	return breweryRepo{url: breweryUrl}
}

func (repository breweryRepo) GetBreweries() (data []models.Brewery, err error) {
	response, err := http.Get(repository.url)
	if err != nil {
		fmt.Print(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}

	return
}
