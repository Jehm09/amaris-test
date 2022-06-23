package pokemon

import (
	"testing"
)

func TestGetPokemonByID(t *testing.T) {
	pokemon := GetPokemonByID(1)
	if pokemon != "bulbasaur" {
		t.Fatal(pokemon)
	}

	pokemon = GetPokemonByID(100)
	if pokemon != "voltorb" {
		t.Fatal(pokemon)
	}

	pokemon = GetPokemonByID(20)
	if pokemon != "raticate" {
		t.Fatal(pokemon)
	}

	pokemon = GetPokemonByID(200)
	if pokemon != "misdreavus" {
		t.Fatal(pokemon)
	}
}

func TestInvalidURL(t *testing.T) {
	_, err := getPokemonByID("invalid url")
	if err == nil {
		t.Fatal("Expected error")
	}
}
