package day12

import (
	"bufio"
	"slices"
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

type Direction = Coords

var Directions = []Direction{Up, Down, Left, Right}

var (
	Up    Direction = Coords{0, -1}
	Down            = Coords{0, 1}
	Left            = Coords{-1, 0}
	Right           = Coords{1, 0}
)

func FindPlantGroups(m *common.TwoDMap[string]) map[string][][]Coords {
	found := map[Coords]bool{}
	groups := map[string][][]Coords{}

	var traverseMap func(x, y, groupIndex int, value string)
	traverseMap = func(x, y, groupIndex int, value string) {
		for _, d := range Directions {
			coords := TraverseMap(m, x, y, value, d)
			if coords == nil {
				continue
			} else if _, ok := found[*coords]; ok {
				continue
			}
			found[*coords] = true
			groups[value][groupIndex] = append(groups[value][groupIndex], *coords)
			nextValue, _ := m.Get(coords.X, coords.Y)
			traverseMap(coords.X, coords.Y, groupIndex, nextValue)
		}
	}

	for i, v := range m.Map {
		x, y := m.FindPosition(i)
		coords := Coords{x, y}
		if _, ok := found[coords]; ok {
			continue
		} else if _, ok := groups[v]; !ok {
			groups[v] = [][]Coords{}
		}
		found[Coords{x, y}] = true
		groups[v] = append(groups[v], []Coords{coords})
		traverseMap(x, y, len(groups[v])-1, v)
	}
	return groups
}

func TraverseMap(m *common.TwoDMap[string], x, y int, value string, direction Direction) *Coords {
	nextX, nextY := x+direction.X, y+direction.Y
	nextValue, err := m.Get(nextX, nextY)
	if err != nil || nextValue != value {
		return nil
	}

	return &Coords{nextX, nextY}
}

func CalculatePerimeter(group []Coords) int {
	sides := 0
	for _, g := range group {
		for _, d := range Directions {
			x, y := g.X+d.X, g.Y+d.Y
			if !slices.Contains(group, Coords{x, y}) {
				sides += 1
			}
		}
	}
	return sides
}
