package operations

import (
	"sort"
	"strings"
)

// SplitNames This function creates an array of alphabetically ordered names obtained from a string separated by commas.
func SplitNames(str string) ([]string, int) {
	if strings.Trim(str, " ") == "" {
		return []string{}, 0
	}

	names := strings.Split(str, ",")

	sort.SliceStable(names, func(i, j int) bool {
		return strings.ToLower(names[i]) < strings.ToLower(names[j])
	})

	return names, len(names)
}

// FriendsStrigns This function checks if two strings are friend strings, only if X=uv and Y=vu.
func FriendsStrigns(X, Y string) bool {
	if len(X) != len(Y) {
		return false
	}

	for i := 1; i < len(X); i++ {
		u := X[0:i]
		v := X[i:]

		if Y == (v + u) {
			return true
		}
	}
	return false
}
