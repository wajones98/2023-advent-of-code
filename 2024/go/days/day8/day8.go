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
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	twoDMap, err := LoadInput(s)
	if err != nil {
		return 0, err
	}
	f := FindFrequencies(twoDMap)

	return FindAllUniqueAntinodes(twoDMap, f), nil
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

func FindAllUniqueAntinodes(m *common.TwoDMap, frequencies map[string][]Coords) int {
	unique := map[Coords]bool{}

	for _, coords := range frequencies {
		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				nodes := FindAntinodes(coords[i], coords[j])
				for _, n := range nodes {
					if isValidAntinode(m.Width, m.Height, n) {
						_, ok := unique[n]
						if !ok {
							unique[n] = true
						}
					}
				}
			}
		}
	}

	return len(unique)
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
	dx := pTwo.X - pOne.X
	dy := pTwo.Y - pOne.Y

	antiNodeOne := Coords{X: pOne.X - dx, Y: pOne.Y - dy}
	antiNodeTwo := Coords{X: pTwo.X + dx, Y: pTwo.Y + dy}

	return []Coords{antiNodeOne, antiNodeTwo}
}

func isValidAntinode(width, height int, node Coords) bool {
	return node.X >= 0 && node.X < width && node.Y >= 0 && node.Y < height
}
