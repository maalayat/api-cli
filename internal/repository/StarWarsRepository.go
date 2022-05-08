package repository

import (
	"encoding/json"
	"fmt"
	"github.com/mayalaat/api-cli/internal/models"
	"io/ioutil"
	"net/http"
)

const (
	warUrl = "https://swapi.dev/api/films/"
)

type warRepo struct {
	url string
}

func StarWarsRepositoryImp() StarWarsRepository {
	return warRepo{url: warUrl}
}

func (w warRepo) GetStarWars() (data models.StarWarsResult, err error) {
	response, err := http.Get(w.url)
	if err != nil {
		fmt.Print(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}

	return
}
