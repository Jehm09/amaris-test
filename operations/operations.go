package operations

import (
	"sort"
	"strings"
)

func SplitNames(str string) ([]string, int) {
	names := strings.Split(str, ",")

	sort.SliceStable(names, func(i, j int) bool {
		return strings.ToLower(names[i]) < strings.ToLower(names[j])
	})

	return names, len(names)
}

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
