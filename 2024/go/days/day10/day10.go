package day10

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/wajones98/advent-of-code/common"
	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 10

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

func LoadInput(s *bufio.Scanner) (*common.TwoDMap[int], error) {
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	width, height := len(lines[0]), len(lines)
	twoDMap := common.NewTwoDMap[int](width, height)

	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			value, err := strconv.Atoi(char)
			if err != nil {
				return nil, err
			}
			err = twoDMap.Put(x, y, value)
			if err != nil {
				return nil, err
			}
		}
	}

	return twoDMap, nil
}
