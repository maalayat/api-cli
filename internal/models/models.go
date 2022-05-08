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

type Punk struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Tagline     string  `json:"tagline"`
	FirstBrewed string  `json:"first_brewed"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Abv         float32 `json:"abv"`
	Ibu         float32 `json:"ibu"`
}
