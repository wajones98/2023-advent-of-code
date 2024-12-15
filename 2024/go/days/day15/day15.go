package day15

import (
	"bufio"

	"github.com/wajones98/advent-of-code/common"
	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 15

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

type Tile rune

const (
	Wall  Tile = '#'
	Robot      = '@'
	Box        = 'O'
	Empty      = '.'
)

func LoadInput(s *bufio.Scanner) (*common.TwoDMap[Tile], string, error) {
	lines := []string{}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	moves := lines[len(lines)-1]
	lines = lines[:len(lines)-1]

	width, height := len(lines[0]), len(lines)
	twoDMap := common.NewTwoDMap[Tile](width, height)

	for y, line := range lines {
		for x, char := range line {
			err := twoDMap.Put(x, y, Tile(char))
			if err != nil {
				return nil, "", err
			}
		}
	}

	return twoDMap, moves, nil
}
