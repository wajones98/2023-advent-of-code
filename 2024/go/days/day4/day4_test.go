package day4

import (
	"reflect"
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

func TestGenerateLineCombinations(t *testing.T) {
	t.Run("GenerateLineCombinations returns the correct number of possible lines", func(t *testing.T) {
		expected := 58
		combinations, err := GenerateLineCombinations(Data)
		if err != nil {
			t.Error(err)
		}
		actual := len(combinations)
		if expected != actual {
			t.Errorf("Incorrect number of combinations. Expected: %d, Got: %d", expected, actual)
		}
	})
}

func TestGenerateHorizontalLines(t *testing.T) {
	t.Run("GenerateHorizontalLines returns the correct horizontal lines", func(t *testing.T) {
		expected := [][]string{
			{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
			{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
			{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
			{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
			{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
			{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
			{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
			{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
			{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
			{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
		}
		expectedLen := len(expected)

		actual, err := GenerateHorizontalLines(Data)
		if err != nil {
			t.Error(err)
		}
		actualLen := len(actual)

		if expectedLen != actualLen {
			t.Errorf("Incorrect number of combinations. Expected: %d, Got: %d", expectedLen, actualLen)
		}

		for i := 0; i < expectedLen; i++ {
			if !reflect.DeepEqual(expected[i], actual[i]) {
				t.Errorf("Elements do not match!\n Expected: %v\nGot: %v\n", expected[i], actual[i])
			}
		}
	})
}
