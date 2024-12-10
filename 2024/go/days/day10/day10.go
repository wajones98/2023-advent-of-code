package day10

import (
	"bufio"
	"fmt"
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
				value = -1
			}
			err = twoDMap.Put(x, y, value)
			if err != nil {
				return nil, err
			}
		}
	}

	return twoDMap, nil
}

func FindTrails(m *common.TwoDMap[int]) {
	for i, h := range m.Map {
		if h == 0 {
			x, y := m.FindPosition(i)
			fmt.Printf("FOUND START: X: %d, Y: %d\n", x, y)
			valid := FindTrail(x, y, h, m, []Coords{})
			fmt.Printf("%v\n", valid)
		}
	}
}

type Direction = int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Coords struct {
	X, Y int
}

func FindTrail(x, y, value int, m *common.TwoDMap[int], coords []Coords) []Coords {
	for i := range 3 {
		for {
			newX, newY, newValue, ok := TraverseTrail(x, y, value, i, m)
			if !ok {
				break
			} else if newValue == 9 {
				coords = append(coords, Coords{newX, newY})
				break
			}

			return FindTrail(newX, newY, newValue, m, coords)
		}
	}

	return coords
}

func TraverseTrail(x, y, currentPointValue int, direction Direction, m *common.TwoDMap[int]) (int, int, int, bool) {
	var nextPointValue int
	var diff int
	var newX, newY int

	switch direction {
	case Up:
		newY = y - 1
		if newY < 0 {
			return -1, -1, -1, false
		}
	case Down:
		newY = y + 1
		if newY >= m.Height {
			return -1, -1, -1, false
		}
	case Left:
		newX = x - 1
		if newX < 0 {
			return -1, -1, -1, false
		}
	case Right:
		newX = x + 1
		if newX >= m.Width {
			return -1, -1, -1, false
		}
	}

	nextPointValue, _ = m.Get(newX, newY)

	diff = nextPointValue - currentPointValue

	if diff != 1 {
		return -1, -1, -1, false
	}

	return newX, newY, nextPointValue, true
}
