package main

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func TestShouldMapOverSliceOfStrings(t *testing.T) {
	strings := []string{"foo", "bar", "buzz"}

	result := mapper[string, string](strings, func(s string) string {
		return s + ":new"
	})

	snaps.MatchJSON(t, result)
}

func TestShouldMapOverSliceOfIntegers(t *testing.T) {
	ints := []int{1, 2, 3}

	result := mapper[int, int](ints, func(s int) int {
		return s * 2
	})

	snaps.MatchJSON(t, result)
}
