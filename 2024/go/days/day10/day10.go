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
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	twoDMap, err := LoadInput(s)
	if err != nil {
		return 0, err
	}

	unique, _ := FindTrails(twoDMap)

	return unique, nil
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

	_, rating := FindTrails(twoDMap)

	return rating, nil
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

type Direction = Coords

var (
	Up    Direction = Coords{0, -1}
	Down            = Coords{0, 1}
	Left            = Coords{-1, 0}
	Right           = Coords{1, 0}
)

func DirectionString(direction Direction) string {
	switch direction {
	case Up:
		return "Up"
	case Right:
		return "Right"
	case Left:
		return "Left"
	case Down:
		return "Down"
	default:
		return "Unknown"
	}
}

type Coords struct {
	X, Y int
}

func FindTrails(m *common.TwoDMap[int]) (int, int) {
	trailPaths := []map[Coords]int{}
	for i, value := range m.Map {
		if value == 0 {
			x, y := m.FindPosition(i)
			trailPaths = append(trailPaths, PossiblePaths(m, x, y, value))
		}
	}

	unique := 0
	rating := 0

	for _, paths := range trailPaths {
		keys := make([]Coords, 0, len(paths))
		for c := range paths {
			keys = append(keys, c)
			rating += paths[c]
		}
		unique += len(keys)
	}

	return unique, rating
}

func PossiblePaths(m *common.TwoDMap[int], x, y, value int) map[Coords]int {
	directions := []Direction{Up, Down, Left, Right}
	coords := map[Coords]int{}

	fmt.Printf("Finding possible paths for x: %d, y: %d\n-----------------------------------\n", x, y)

	var helper func(x, y, value int)
	helper = func(x, y, value int) {
		fmt.Printf("\nError: %v for %d, %d\n", m.CheckBounds(x, y), x, y)
		for _, d := range directions {
			fmt.Printf("Checking direction %s for %d, %d\n", DirectionString(d), x, y)
			found, ok := TraverseTrail(m, x, y, value, d)
			fmt.Printf("	Found: %v, Ok: %t\n", found, ok)
			if !ok {
				continue
			} else if found != nil {
				count, ok := coords[*found]
				if !ok {
					count = 0
				}
				coords[*found] = count + 1
				continue
			}
			newValue, _ := m.Get(x+d.X, y+d.Y)
			helper(x+d.X, y+d.Y, newValue)

			fmt.Printf("\n")
		}

		fmt.Printf("--------------------------------\n")

	}
	helper(x, y, value)

	return coords
}

func TraverseTrail(m *common.TwoDMap[int], x, y, value int, direction Direction) (*Coords, bool) {
	nextX, nextY := x+direction.X, y+direction.Y
	nextValue, err := m.Get(nextX, nextY)
	if err != nil || nextValue-value != 1 {
		return nil, false
	}

	if nextValue == 9 {
		return &Coords{nextX, nextY}, true
	}

	return nil, true
}

// func FindTrail(x, y, value int, m *common.TwoDMap[int], initial []Coords, index int) []Coords {
// 	if index == len(m.Map) {
// 		return initial
// 	}
// 	coords := initial
// Loop:
// 	for i := range 4 {
// 		for {
// 			newX, newY, newValue, ok := TraverseTrail(x, y, value, i, m)
//
// 			fmt.Printf("X: %d, Y: %d, Value: %d, Ok: %t\n", newX, newY, newValue, ok)
//
// 			if !ok {
// 				continue Loop
// 			} else if newValue == 9 {
// 				coords = append(coords, Coords{newX, newY})
// 				continue Loop
// 			}
//
// 			x = newX
// 			y = newY
// 			value = newValue
// 		}
// 	}
// 	return FindTrail(x, y, value, m, coords, index+1)
// }
//
// func TraverseTrail(x, y, currentPointValue int, direction Direction, m *common.TwoDMap[int]) (int, int, int, bool) {
// 	var nextPointValue int
// 	var diff int
// 	var newX, newY int = x, y
//
// 	switch direction {
// 	case Up:
// 		newY = y - 1
// 		if newY < 0 {
// 			return -1, -1, -1, false
// 		}
// 	case Down:
// 		newY = y + 1
// 		if newY >= m.Height {
// 			return -1, -1, -1, false
// 		}
// 	case Left:
// 		newX = x - 1
// 		if newX < 0 {
// 			return -1, -1, -1, false
// 		}
// 	case Right:
// 		fmt.Printf("CHECKING RIGHT FOR %d, %d of VALUE %d\n", x, y, currentPointValue)
// 		newX = x + 1
// 		if newX >= m.Width {
// 			return -1, -1, -1, false
// 		}
// 	}
//
// 	nextPointValue, _ = m.Get(newX, newY)
//
// 	diff = nextPointValue - currentPointValue
//
// 	if diff != 1 {
// 		return -1, -1, -1, false
// 	}
//
// 	return newX, newY, nextPointValue, true
// }
