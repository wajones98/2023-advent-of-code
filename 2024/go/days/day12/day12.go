package day12

import (
	"bufio"
	"strings"

	"github.com/wajones98/advent-of-code/common"
	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 12

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

func LoadInput(s *bufio.Scanner) (*common.TwoDMap[string], error) {
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	width, height := len(lines[0]), len(lines)
	twoDMap := common.NewTwoDMap[string](width, height)

	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			err := twoDMap.Put(x, y, char)
			if err != nil {
				return nil, err
			}
		}
	}

	return twoDMap, nil
}

type Coords struct {
	X, Y int
}

func FindGroups(m *common.TwoDMap[string]) map[string][]Coords {
	found := map[Coords]bool{}
	groups := map[string][]Coords{}
	for i, _ := range m.Map {
		x, y := m.FindPosition(i)
		if _, ok := found[Coords{x, y}]; ok {
			continue
		}
	}
	return groups
}
