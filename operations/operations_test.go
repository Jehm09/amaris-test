package operations

import (
	"testing"
)

func TestSplitNames(t *testing.T) {
	names := "joe, carlos, seabastian"
	namesArray, size := SplitNames(names)
	if size != len(namesArray) {
		t.Fatal("Error with size")
	}

	names = "joe, carlos, seabastian"
	namesArray, size = SplitNames(names)
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
