package day4

import (
	"log"
	"testing"
)

var Data []string = []string{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func TestGenerateLines(t *testing.T) {
	t.Run("GenerateLines returns the correct number of possible lines", func(t *testing.T) {
		expected := 58
		combinations, err := GenerateLineCombinations(Data)
		if err != nil {
			panic(err)
		}
		actual := len(combinations)
		if expected != actual {
			log.Fatalf("Incorrect number of combinations. Expected: %d, Got: %d", expected, actual)
		}
	})
}
