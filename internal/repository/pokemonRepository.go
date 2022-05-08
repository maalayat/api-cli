package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mayalaat/api-cli/internal/models"
)

const (
	pokeUrl = "https://pokeapi.co/api/v2/pokemon"
)

type repository struct {
	url string
}

func PokemonRepositoryImp() PokemonRepository {
	return &repository{url: pokeUrl}
}

func (repo *repository) GetPokemons() (data models.Data, err error) {
	response, err := http.Get(repo.url)
	if err != nil {
		fmt.Print(err)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(contents, &data)
	if err != nil {
		fmt.Println(err)
	}

	return
}
