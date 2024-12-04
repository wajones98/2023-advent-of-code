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

func LoadLines() ([]string, error) {
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
	_, err := LoadLines()
	if err != nil {
		panic(err)
	}
	return 0
}

func Part2() int {
	return 0
}
