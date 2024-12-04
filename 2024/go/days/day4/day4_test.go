package day4

import (
	"slices"
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
		combinations := GenerateLineCombinations(Data)
		actual := len(combinations)
		if expected != actual {
			t.Errorf("Incorrect number of combinations. Expected: %d, Got: %d", expected, actual)
		}
	})
}

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

func TestGenerateDiagonalLines(t *testing.T) {

	t.Run("Produces correct combinations left to right", func(t *testing.T) {
		// Smaller subset of data for sake of time
		data := []string{
			"MMMS",
			"MSAM",
			"AMXS",
			"MSAM",
		}
		expected := []string{
			"M",
			"MM",
			"ASM",
			"MMAS",
			"SXM",
			"AS",
			"M",
		}
		expectedLen := len(expected)

		actual := GenerateDiagonalLines(data)
		actualLen := len(actual)
		if expectedLen != actualLen {
			t.Errorf("Incorrect number of combinations. Expected: %d, Got: %d", expectedLen, actualLen)
		}

		for i := 0; i < expectedLen; i++ {
			if expected[i] != actual[i] {
			}
		}
		for _, line := range actual {
			if !slices.Contains(expected, line) {
				t.Errorf("Expected %s to be present in %v\n", line, expected)
			}
		}
	})

	t.Run("Produces correct combinations right to left", func(t *testing.T) {
		// Smaller subset of data for sake of time
		data := []string{
			"XMA",
			"ASM",
			"MMM",
		}
		expected := []string{
			"X",
			"AM",
			"MSA",
			"MM",
			"M",
		}
		expectedLen := len(expected)

		actual := GenerateDiagonalLines(data)
		actualLen := len(actual)
		if expectedLen != actualLen {
			t.Errorf("Incorrect number of combinations. Expected: %d, Got: %d", expectedLen, actualLen)
		}

		for i := 0; i < expectedLen; i++ {
			if expected[i] != actual[i] {
			}
		}
		for _, line := range actual {
			if !slices.Contains(expected, line) {
				t.Errorf("Expected %s to be present in %v\n", line, expected)
			}
		}
	})
}
