package day8

import (
	"bufio"
	"strings"

	"github.com/wajones98/advent-of-code/common"
	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 8

type Coords struct {
	X int
	Y int
}

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
		panic(err)
	}
	defer closeFile()

	return 0, err
}

func LoadInput(s *bufio.Scanner) (*common.TwoDMap, error) {
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	width, height := len(lines[0]), len(lines)
	twoDMap := common.NewTwoDMap(width, height)

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

func FindFrequencies(m *common.TwoDMap) map[string][]Coords {
	frequencies := map[string][]Coords{}

	for i, c := range m.Map {
		if c == "." {
			continue
		}
		coords, ok := frequencies[c]
		if !ok {
			coords = []Coords{}
		}
		x, y := m.FindPosition(i)
		coords = append(coords, Coords{X: x, Y: y})
		frequencies[c] = coords
	}

	return frequencies
}

func FindAntinodes(pOne, pTwo Coords) []Coords {
	antinodes := []Coords{}

	diffY := pTwo.Y - pOne.Y
	diffX := pTwo.X - pOne.X
	slope := float64(diffY) / float64(diffX)
	if slope < 1 {
		delta := int(1.0 / slope)
		antinodes = append(antinodes, Coords{X: pOne.X - delta, Y: pOne.Y - 1}, Coords{X: pTwo.X + delta, Y: pTwo.Y - 1})
	} else {
		delta := int(slope)
		antinodes = append(antinodes, Coords{X: pOne.X - 1, Y: pOne.Y - delta}, Coords{X: pTwo.X + 1, Y: pTwo.Y + delta})
	}

	return antinodes
}
