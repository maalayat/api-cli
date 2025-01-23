package repository

import (
	"encoding/json"
	"fmt"
	"github.com/maalayat/api-cli/internal/models"
	"io/ioutil"
	"net/http"
)

const (
	punkUrl = "https://api.punkapi.com/v2/beers"
)

type punkRepo struct {
	url string
}

func PunkRepositoryImp() PunkRepository {
	return punkRepo{url: punkUrl}
}

func (repository punkRepo) GetPunks() (data []models.Punk, err error) {
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
