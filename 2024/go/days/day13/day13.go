package day13

import (
	"bufio"
	"fmt"

	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 13

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

type Coords struct {
	X, Y int
}

type Prize struct {
	Location,
	ButtonA,
	ButtonB Coords
}

func LoadInput(s *bufio.Scanner) ([]Prize, error) {
	prizes := []Prize{}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}

		buttonALine := line
		s.Scan()
		buttonBLine := s.Text()
		s.Scan()
		prizeLine := s.Text()

		fmt.Printf("%s\n%s\n%s\n", buttonALine, buttonBLine, prizeLine)
	}
	return prizes, nil
}
