package main

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokeAPIResponse struct {
	Results []Pokemon `json:"results"`
}
