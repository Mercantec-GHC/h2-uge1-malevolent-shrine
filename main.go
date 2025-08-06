package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	//load pokeAPIResponse from JSON file
	fmt.Println("Loading a Pokemon...")

	response, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=10")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	defer response.Body.Close()
	//read the response body

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	//unmarshal the JSON data into PokeAPIResponse struct

	var pokeAPIResponse PokeAPIResponse
	err = json.Unmarshal(body, &pokeAPIResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for _, pokemon := range pokeAPIResponse.Results {
		fmt.Println(pokemon.Name)
	}

}
