package day11

import (
	"bufio"
	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 11

func Run() (*days.Result[int, int], error) {
	pOne, err := Part1()
	if err != nil {
		return nil, err
	}

	pTwo, err := Part2()
	if err != nil {
		return nil, err
	}

	return &days.Result[int, int]{
		Part1: pOne,
		Part2: pTwo,
	}, nil
}

func Part1() (int, error) {
	_, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	return 0, nil
}

func Part2() (int, error) {
	_, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	return 0, nil
}

func LoadInput(s *bufio.Scanner) ([]int, error) {
	for s.Scan() {
	}

	return []int{}, nil
}
