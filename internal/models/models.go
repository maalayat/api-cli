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

type StarWarsResult struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Results  *[]StarWars `json:"results"`
}

type StarWars struct {
	Title        string `json:"title"`
	EpisodeId    int    `json:"episode_id"`
	OpeningCrawl string `json:"opening_crawl"`
	Director     string `json:"director"`
	Producer     string `json:"producer"`
	ReleaseDate  string `json:"release_date"`
}

type Brewery struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	BreweryType string `json:"brewery_type"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
	Website     string `json:"website_url"`
}
