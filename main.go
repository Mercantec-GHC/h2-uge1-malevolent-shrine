package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

var pokemonList []Pokemon

func main() {
	// Сначала загружаем покемонов
	loadPokemon()

	// Потом запускаем сервер
	http.HandleFunc("/", handleGame)
	http.HandleFunc("/check", handleCheck)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
func loadPokemon() {
	fmt.Println("Loading a Pokemon...")

	response, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=10")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var pokeAPIResponse PokeAPIResponse
	err = json.Unmarshal(body, &pokeAPIResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	pokemonList = pokeAPIResponse.Results

	for _, pokemon := range pokemonList {
		fmt.Println(pokemon.Name)
	}
}

func handleGame(w http.ResponseWriter, r *http.Request) {
	//randomly select a Pokemon from the list
	randomIndex := rand.Intn(len(pokemonList))
	correctPokemon := pokemonList[randomIndex]

	///load the details of the selected Pokemon
	detailsURL := correctPokemon.URL
	response, err := http.Get(detailsURL)
	if err != nil {
		http.Error(w, "Error fetching Pokemon details", http.StatusInternalServerError)
		return
	}

	defer response.Body.Close()

	var details PokemonDetails
	body, err := io.ReadAll(response.Body)
	json.Unmarshal(body, &details)
	imageURL := details.Sprites.FrontDefault

	// Generate two extra random names
	var options []string
	options = append(options, correctPokemon.Name)

	for _, p := range pokemonList {
		if p.Name != correctPokemon.Name {
			options = append(options, p.Name)
		}
		if len(options) == 3 {
			break
		}
	}

	//Shuffle the options
	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	// HTML-страница с аниме-стилем
	html := `<html>
	<head>
		<title>Gæt Pokémon!</title>
		<style>
			@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@700&family=Pacifico&display=swap');
			body {
				background: linear-gradient(135deg, #ffb6e6 0%, #7fffd4 100%);
				min-height: 100vh;
				margin: 0;
				font-family: 'Montserrat', Arial, sans-serif;
				display: flex;
				flex-direction: column;
				align-items: center;
				justify-content: center;
			}
			h1 {
				font-family: 'Pacifico', cursive;
				color: #ff4fa3;
				font-size: 2.8rem;
				text-shadow: 2px 4px 12px #fff, 0 0 10px #7fffd4;
				margin-bottom: 1.5rem;
			}
			form {
				background: rgba(255,255,255,0.85);
				border-radius: 24px;
				box-shadow: 0 8px 32px 0 rgba(127,255,212,0.25), 0 1.5px 8px 0 #ffb6e6;
				padding: 2.5rem 2rem;
				display: flex;
				flex-direction: column;
				align-items: center;
			}
			button {
				background: linear-gradient(90deg, #ffb6e6 0%, #7fffd4 100%);
				color: #fff;
				font-size: 1.5rem;
				font-family: 'Montserrat', Arial, sans-serif;
				font-weight: bold;
				letter-spacing: 1px;
				border: 3px solid #ff4fa3;
				border-radius: 24px;
				padding: 1rem 2.8rem;
				margin: 0.7rem 0;
				box-shadow: 0 4px 16px #ffb6e6, 0 2px 8px #7fffd4, 0 0 0 4px #fff8, 0 0 12px 2px #ffe066;
				cursor: pointer;
				transition: transform 0.12s, box-shadow 0.22s, background 0.22s;
				text-shadow: 1px 2px 8px #ffb6e6, 0 0 4px #7fffd4, 0 0 2px #fff;
				position: relative;
				overflow: hidden;
			}
			button:before {
				content: '';
				display: block;
				position: absolute;
				top: 0; left: 0; right: 0; bottom: 0;
				background: url('https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/items/poke-ball.png') no-repeat center right 18px/32px, linear-gradient(90deg, #fff0 60%, #ffe06644 100%);
				pointer-events: none;
				opacity: 0.7;
			}
			button:hover {
				transform: scale(1.09) rotate(-2deg);
				box-shadow: 0 8px 32px #ff4fa3, 0 4px 16px #7fffd4, 0 0 0 6px #fff8, 0 0 24px 4px #ffe066;
				background: linear-gradient(90deg, #ff4fa3 0%, #7fffd4 100%);
				border-color: #ffe066;
			}
		</style>
	</head>
	<body>
		<h1>Hvem er denne Pokémon?</h1>
<img src="` + imageURL + `" alt="Pokemon" style="width:150px; image-rendering: pixelated;"><br/><br/>

		<form method="POST" action="/check">
			<input type="hidden" name="correct" value="` + correctPokemon.Name + `">
			<input type="hidden" name="index" value="` + fmt.Sprint(randomIndex) + `">
			<button name="answer" value="` + options[0] + `">` + options[0] + `</button><br/>
			<button name="answer" value="` + options[1] + `">` + options[1] + `</button><br/>
			<button name="answer" value="` + options[2] + `">` + options[2] + `</button>
		</form>
	</body>
</html>`

	fmt.Fprint(w, html)
}

func handleCheck(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	correct := r.FormValue("correct")
	answer := r.FormValue("answer")

	result := "Forkert svar. Prøv igen!"

	if answer == correct {
		result = "Korrekt! Du gættede rigtigt: " + correct
	}
	html := `<html>
	<head><title>Resultat</title></head>
	<body>
		<h1>` + result + `</h1>
		<p>Det rigtige svar var: ` + correct + `</p>
		<a href="/">Prøv igen</a>
	</body>
</html>`

	fmt.Fprint(w, html)
}
