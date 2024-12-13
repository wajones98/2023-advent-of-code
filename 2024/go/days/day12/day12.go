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

type Coords[T int | float64] struct {
	X, Y T
}

type Direction = Coords[int]
type Corner = Coords[float64]

var Directions = []Direction{Up, Down, Left, Right}

var (
	Up        Direction = Coords[int]{0, -1}
	Down                = Coords[int]{0, 1}
	Left                = Coords[int]{-1, 0}
	Right               = Coords[int]{1, 0}
	UpLeft              = Coords[int]{-1, -1}
	UpRight             = Coords[int]{1, -1}
	DownLeft            = Coords[int]{-1, 1}
	DownRight           = Coords[int]{1, 1}
)

var (
	TopLeft     Corner = Coords[float64]{-0.5, -0.5}
	TopRight           = Coords[float64]{0.5, -0.5}
	BottomLeft         = Coords[float64]{-0.5, 0.5}
	BottomRight        = Coords[float64]{0.5, 0.5}
)

func FindPlantGroups(m *common.TwoDMap[string]) map[string][][]Coords[int] {
	found := map[Coords[int]]bool{}
	region := map[string][][]Coords[int]{}

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
			region[value][groupIndex] = append(region[value][groupIndex], *coords)
			nextValue, _ := m.Get(coords.X, coords.Y)
			traverseMap(coords.X, coords.Y, groupIndex, nextValue)
		}
	}

	for i, v := range m.Map {
		x, y := m.FindPosition(i)
		coords := Coords[int]{x, y}
		if _, ok := found[coords]; ok {
			continue
		} else if _, ok := region[v]; !ok {
			region[v] = [][]Coords[int]{}
		}
		found[Coords[int]{x, y}] = true
		region[v] = append(region[v], []Coords[int]{coords})
		traverseMap(x, y, len(region[v])-1, v)
	}
	return region
}

func TraverseMap(m *common.TwoDMap[string], x, y int, value string, direction Direction) *Coords[int] {
	nextX, nextY := x+direction.X, y+direction.Y
	nextValue, err := m.Get(nextX, nextY)
	if err != nil || nextValue != value {
		return nil
	}

	return &Coords[int]{nextX, nextY}
}

func CalculatePerimeter(region []Coords[int]) int {
	edges := 0
	for _, g := range region {
		for _, d := range Directions {
			x, y := g.X+d.X, g.Y+d.Y
			if !slices.Contains(region, Coords[int]{x, y}) {
				edges += 1
			}
		}
	}
	return edges
}

func CalculateSides(region []Coords[int]) int {
	sides := 0

	corners := []Coords[float64]{}
	cornerDirections := []Corner{TopLeft, TopRight, BottomLeft, BottomRight}
	for _, p := range region {
		for _, cd := range cornerDirections {
			corner := Coords[float64]{
				X: float64(p.X) + cd.X,
				Y: float64(p.Y) + cd.Y,
			}
			if !slices.Contains(corners, corner) {
				corners = append(corners, corner)
			}
		}
	}

	for _, c := range corners {
		inRegion := []int{}
		for _, cd := range cornerDirections {
			coord := Coords[int]{
				X: int(c.X + cd.X),
				Y: int(c.Y + cd.Y),
			}
			if slices.Contains(region, coord) {
				inRegion = append(inRegion, 1)
			} else {
				inRegion = append(inRegion, 0)
			}
		}
		sum := 0
		for _, ir := range inRegion {
			sum += ir
		}

		sumDiagLeftRight := inRegion[0] + inRegion[3]
		sumDiagRightLeft := inRegion[1] + inRegion[2]

		if sum == 1 || sum == 3 {
			sides += 1
		} else if sumDiagLeftRight == 2 && sumDiagRightLeft == 0 {
			sides += 2
		} else if sumDiagLeftRight == 0 && sumDiagRightLeft == 2 {
			sides += 2
		}
	}

	return sides
}

func CalculatePrice(plants map[string][][]Coords[int], withDiscount bool) int {
	price := 0

	for _, plant := range plants {
		for _, region := range plant {
			var sides int
			if withDiscount {
				sides = CalculateSides(region)
			} else {
				sides = CalculatePerimeter(region)
			}
			price += sides * len(region)
		}
	}

	return price
}
