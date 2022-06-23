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

	names = ""
	_, size = SplitNames(names)
	if size != 0 {
		t.Fatal("Size will be 0")
	}

	names = "joe"
	_, size = SplitNames(names)
	if size != 1 {
		t.Fatal("Size will be 1")
	}
}

func TestFriendlyStrings(t *testing.T) {

	X := "tokyo"
	friendlyStrings := []string{"otoky", "yotok", "okyot", "kyoto"}
	for _, Y := range friendlyStrings {
		isFriendlyStrings := FriendsStrigns(X, Y)
		if !isFriendlyStrings {
			t.Fatal("Must be strings friends")
		}
	}

	isFriendlyStrings := FriendsStrigns(X, "toyoka")
	if isFriendlyStrings {
		t.Fatal("Different lengths")
	}

	isFriendlyStrings = FriendsStrigns(X, "toyok")
	if isFriendlyStrings {
		t.Fatal("Must be not strings friends")
	}
}
