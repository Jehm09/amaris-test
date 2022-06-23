package pokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	API_POKEMON = "https://pokeapi.co/api/v2/pokemon-form"
)

type formData struct {
	Pokemon pokemon `json:"pokemon"`
}

type pokemon struct {
	Name string `json:"name"`
}

func close(closable io.Closer) {
	if err := closable.Close(); err != nil {
		panic(err)
	}
}

func getPokemonByID(url string) (*formData, error) {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer close(res.Body)

	formData := &formData{}
	err = json.NewDecoder(res.Body).Decode(&formData)
	if err != nil {
		return nil, err
	}

	return formData, err
}

func GetPokemonByID(id int) string {
	fullURL := fmt.Sprintf("%s/%d", API_POKEMON, id)

	formData, err := getPokemonByID(fullURL)
	if err != nil {
		panic(err)
	}

	return formData.Pokemon.Name
}
