package day4

import (
	"github.com/wajones98/advent-of-code/input"

	"github.com/wajones98/advent-of-code/days"
)

const Day int = 4

func Run() (*days.Result[int, int], error) {

	return &days.Result[int, int]{
		Part1: Part1(),
		Part2: Part2(),
	}, nil
}

func LoadInput() ([]string, error) {
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return nil, err
	}
	defer closeFile()
	lines := []string{}
	for s.Scan() {
		line := s.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func Part1() int {
	_, err := LoadInput()
	if err != nil {
		panic(err)
	}
	return 0
}

func Part2() int {
	return 0
}

// Given a grid, generate all possible lines vertically, horizontally and diagonally.
func GenerateLineCombinations(lines []string) ([]string, error) {
	return nil, nil
}

func GenerateVerticalLines(lines []string) []string {
	combinations := []string{}
	for i, _ := range lines {
		result := ""
		for _, lineInner := range lines {
			result += string(lineInner[i])
		}
		combinations = append(combinations, result)
	}
	return combinations
}

func GenerateDiagonalLines(lines []string) []string {
	combinations := []string{}
	return combinations
}
