package repository

import (
	"encoding/json"
	"fmt"
	"github.com/maalayat/api-cli/internal/models"
	"io/ioutil"
	"net/http"
)

const (
	warsUrl = "https://swapi.dev/api/films/"
)

type warsRepo struct {
	url string
}

func StarWarsRepositoryImp() StarWarsRepository {
	return warsRepo{url: warsUrl}
}

func (repository warsRepo) GetStarWars() (data models.StarWarsResult, err error) {
	response, err := http.Get(repository.url)
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
