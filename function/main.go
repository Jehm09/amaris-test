package main

import (
	"amaris/operations"
	"amaris/pokemon"
	"fmt"
)

func main() {
	pokemonName := pokemon.GetPokemonByID(1)
	fmt.Println(pokemonName)

	names, size := operations.SplitNames("Joe,Carlos,Sebastian")
	fmt.Printf("Names :%v \nSize: %d\n", names, size)

	X, Y := "tokyo", "kyoto"
	isFriendlyStrings := operations.FriendsStrigns(X, Y)
	if isFriendlyStrings {
		fmt.Printf("%s %s They are friendly strings\n", X, Y)
	} else {
		fmt.Printf("%s %s They are not friendly strings\n", X, Y)
	}
}
