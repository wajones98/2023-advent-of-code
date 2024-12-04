package day4

import (
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

//
// func TestGenerateLineCombinations(t *testing.T) {
// 	t.Run("GenerateLineCombinations returns the correct number of possible lines", func(t *testing.T) {
// 		expected := 58
// 		combinations, err := GenerateLineCombinations(Data)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		actual := len(combinations)
// 		if expected != actual {
// 			t.Errorf("Incorrect number of combinations. Expected: %d, Got: %d", expected, actual)
// 		}
// 	})
// }

func TestGenerateVerticalLines(t *testing.T) {
	t.Run("GenerateVerticalLines returns the correct vertical lines", func(t *testing.T) {
		expected := []string{
			"MMAMXXSSMM",
			"MSMSMXMAAX",
			"MAXAAASXMM",
			"SMSMSMMAMX",
			"XXXAAMSMMA",
			"XMMSMXAAXX",
			"MSAMXXSSMM",
			"AMASAAXAMA",
			"SSMMMMSAMS",
			"MAMXMASAMX",
		}
		expectedLen := len(expected)

		actual := GenerateVerticalLines(Data)
		actualLen := len(actual)
		if expectedLen != actualLen {
			t.Errorf("Incorrect number of combinations. Expected: %d, Got: %d", expectedLen, actualLen)
		}

		for i := 0; i < expectedLen; i++ {
			if expected[i] != actual[i] {
				t.Errorf("Lines do not match!\n Expected: %v\nGot: %v\n", expected[i], actual[i])
			}
		}
	})
}
