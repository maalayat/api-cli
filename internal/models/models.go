package models

type PokemonResult struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  *[]Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
