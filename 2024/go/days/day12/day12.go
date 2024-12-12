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
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	twoDMap, err := LoadInput(s)
	if err != nil {
		return 0, err
	}

	plantGroups := FindPlantGroups(twoDMap)

	return CalculatePrice(plantGroups, false), nil
}

func Part2() (int, error) {
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	twoDMap, err := LoadInput(s)
	if err != nil {
		return 0, err
	}

	plantGroups := FindPlantGroups(twoDMap)

	return CalculatePrice(plantGroups, true), nil
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

	UpLeft    = Coords{-1, -1}
	UpRight   = Coords{1, -1}
	DownLeft  = Coords{-1, 1}
	DownRight = Coords{1, 1}
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
	edges := 0
	for _, g := range group {
		for _, d := range Directions {
			x, y := g.X+d.X, g.Y+d.Y
			if !slices.Contains(group, Coords{x, y}) {
				edges += 1
			}
		}
	}
	return edges
}

func CalculateSides(group []Coords) int {
	sides := 0
	foundCorners := map[Coords][]Coords{}
	for _, g := range group {

		surroundingCoords := map[Direction]Coords{
			Up:        Coords{g.X + Up.X, g.Y + Up.Y},
			Down:      Coords{g.X + Down.X, g.Y + Down.Y},
			Left:      Coords{g.X + Left.X, g.Y + Left.Y},
			Right:     Coords{g.X + Right.X, g.Y + Right.Y},
			UpLeft:    Coords{g.X + UpLeft.X, g.Y + UpLeft.Y},
			UpRight:   Coords{g.X + UpRight.X, g.Y + UpRight.Y},
			DownLeft:  Coords{g.X + DownLeft.X, g.Y + DownLeft.Y},
			DownRight: Coords{g.X + DownRight.X, g.Y + DownRight.Y},
		}

		foundCorners[g] = make([]Coords, 0)

		if !slices.Contains(group, surroundingCoords[UpLeft]) && !slices.Contains(group, surroundingCoords[Up]) && !slices.Contains(group, surroundingCoords[Left]) {
			sides += 1
			foundCorners[g] = append(foundCorners[g], surroundingCoords[UpLeft])
		}
		if !slices.Contains(group, surroundingCoords[UpRight]) && !slices.Contains(group, surroundingCoords[Up]) && !slices.Contains(group, surroundingCoords[Right]) {
			sides += 1
			foundCorners[g] = append(foundCorners[g], surroundingCoords[UpRight])
		}
		if !slices.Contains(group, surroundingCoords[DownLeft]) && !slices.Contains(group, surroundingCoords[Down]) && !slices.Contains(group, surroundingCoords[Left]) {
			sides += 1
			foundCorners[g] = append(foundCorners[g], surroundingCoords[DownLeft])
		}
		if !slices.Contains(group, surroundingCoords[DownRight]) && !slices.Contains(group, surroundingCoords[Down]) && !slices.Contains(group, surroundingCoords[Right]) {
			sides += 1
			foundCorners[g] = append(foundCorners[g], surroundingCoords[DownRight])
		}
		if !slices.Contains(group, surroundingCoords[UpRight]) && (slices.Contains(group, surroundingCoords[Up]) && slices.Contains(group, surroundingCoords[Right])) {
			sides += 1
			foundCorners[g] = append(foundCorners[g], surroundingCoords[UpRight])
		}
		if !slices.Contains(group, surroundingCoords[UpLeft]) && (slices.Contains(group, surroundingCoords[Up]) && slices.Contains(group, surroundingCoords[Left])) {
			sides += 1
			foundCorners[g] = append(foundCorners[g], surroundingCoords[UpLeft])
		}
		if !slices.Contains(group, surroundingCoords[DownLeft]) && slices.Contains(group, surroundingCoords[Down]) && slices.Contains(group, surroundingCoords[Left]) {
			sides += 1
			foundCorners[g] = append(foundCorners[g], surroundingCoords[DownLeft])
		}
		if !slices.Contains(group, surroundingCoords[DownRight]) && slices.Contains(group, surroundingCoords[Down]) && slices.Contains(group, surroundingCoords[Right]) {
			sides += 1
			foundCorners[g] = append(foundCorners[g], surroundingCoords[DownRight])
		}
		if !slices.Contains(group, surroundingCoords[Up]) && slices.Contains(group, surroundingCoords[UpLeft]) && slices.Contains(group, surroundingCoords[UpRight]) {
			sides += 1
			foundCorners[g] = append(foundCorners[g], surroundingCoords[DownRight])
		}
		if !slices.Contains(group, surroundingCoords[Down]) && slices.Contains(group, surroundingCoords[DownLeft]) && slices.Contains(group, surroundingCoords[DownRight]) {
			sides += 1
			foundCorners[g] = append(foundCorners[g], surroundingCoords[DownRight])
		}
	}
	return sides
}

func CalculatePrice(plants map[string][][]Coords, withDiscount bool) int {
	price := 0

	for _, plant := range plants {
		for _, groups := range plant {
			var sides int
			if withDiscount {
				sides = CalculateSides(groups)
			} else {
				sides = CalculatePerimeter(groups)
			}
			price += sides * len(groups)
		}
	}

	return price
}
