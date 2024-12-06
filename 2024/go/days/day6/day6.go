package day6

import (
	"bufio"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/wajones98/advent-of-code/input"

	"github.com/wajones98/advent-of-code/days"
)

const Day int = 6

func Run() (*days.Result[int, int], error) {

	return &days.Result[int, int]{
		Part1: Part1(),
		Part2: Part2(),
	}, nil
}

type Direction = string

const (
	Up, Right, Down, Left Direction = "^", ">", "V", "<"
)

type Visited struct {
	X, Y  uint
	Count uint
}

type Guard struct {
	X         uint
	Y         uint
	Direction Direction
	Visited   map[Direction][]Visited
}

func (g *Guard) ChangeDirection(d Direction) {
	switch g.Direction {
	case Up:
		g.Direction = Right
	case Right:
		g.Direction = Down
	case Down:
		g.Direction = Left
	default:
		g.Direction = Up
	}
}

var Directions []Direction = []Direction{Up, Down, Left, Right}

func Part1() int {
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		panic(err)
	}
	defer closeFile()

	twoDMap, err := LoadInput(s)
	if err != nil {
		panic(err)
	}

	total, err := Patrol(twoDMap)
	if err != nil {
		panic(err)
	}
	return total
}

func Part2() int {
	_, closeFile, err := input.GetInput(Day)
	if err != nil {
		panic(err)
	}
	defer closeFile()

	return 0
}

type TwoDMap struct {
	Map    []string
	Width  uint
	Height uint
}

func NewTwoDMap(width, height uint) *TwoDMap {
	return &TwoDMap{
		Map:    make([]string, width*height),
		Width:  width,
		Height: height,
	}
}

func (m *TwoDMap) Put(x, y uint, r string) error {
	err := m.checkBounds(x, y)
	if err != nil {
		return err
	}
	m.Map[m.getIndex(x, y)] = r
	return nil
}

func (m *TwoDMap) Get(x, y uint) (string, error) {
	err := m.checkBounds(x, y)
	if err != nil {
		return "", err
	}
	return m.Map[m.getIndex(x, y)], nil
}

func (m *TwoDMap) getIndex(x, y uint) uint {
	return y*m.Width + x
}

func (m *TwoDMap) checkBounds(x, y uint) error {
	if x > m.Width {
		return fmt.Errorf("%d is out of bounds %d", x, m.Width)
	} else if y > m.Height {
		return fmt.Errorf("%d is out of bounds %d", y, m.Height)
	}

	return nil
}

func (m *TwoDMap) FindPosition(i uint) (uint, uint) {
	y := i / m.Width
	x := i % m.Width
	return x, y
}

func (m *TwoDMap) String() string {
	result := ""
	for i, c := range m.Map {
		result += c
		x := (i + 1) % int(m.Width)
		if x == 0 {
			result += "\n"
		}
	}
	result += "\n"
	return result
}

func LoadInput(s *bufio.Scanner) (*TwoDMap, error) {
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	width, height := uint(len(lines[0])), uint(len(lines))
	twoDMap := NewTwoDMap(width, height)

	for y, line := range lines {
		chars := strings.Split(line, "")
		for x, char := range chars {
			err := twoDMap.Put(uint(x), uint(y), char)
			if err != nil {
				return nil, err
			}
		}
	}

	return twoDMap, nil
}

func FindGuard(m *TwoDMap) (*Guard, error) {
	for i, p := range m.Map {
		if slices.Contains(Directions, p) {
			x, y := m.FindPosition(uint(i))
			return &Guard{
				X:         x,
				Y:         y,
				Direction: p,
				Visited:   map[Direction][]Visited{},
			}, nil
		}
	}

	return nil, errors.New("Could not find guard")
}

func Patrol(m *TwoDMap) (int, error) {
	total := 0
	guard, err := FindGuard(m)
	if err != nil {
		return total, err
	}

	hasExited := false
Loop:
	for !hasExited {
		exited, unique := false, false
		switch guard.Direction {
		case Up:
			exited, unique, err = patrol(m, guard, guard.X, guard.Y-1)
		case Right:
			exited, unique, err = patrol(m, guard, guard.X+1, guard.Y)
		case Down:
			exited, unique, err = patrol(m, guard, guard.X, guard.Y+1)
		case Left:
			exited, unique, err = patrol(m, guard, guard.X-1, guard.Y)
		default:
			break Loop
		}

		if err != nil {
			return total, err
		}
		hasExited = exited
		if unique {
			total += 1
		}
		fmt.Print(m.String())
	}

	return total, nil
}

func patrol(m *TwoDMap, guard *Guard, x, y uint) (bool, bool, error) {
	if guard.Y == 0 || guard.Y == m.Height-1 || guard.X == 0 || guard.X == m.Width-1 {
		return true, true, nil
	}

	c, err := m.Get(x, y)
	if err != nil {
		return false, false, err
	}

	m.Put(guard.X, guard.Y, "X")
	if err != nil {
		return false, false, err
	}

	if c == "#" {
		guard.ChangeDirection(guard.Direction)
	} else {
		guard.X = x
		guard.Y = y

	}
	err = m.Put(guard.X, guard.Y, guard.Direction)
	if err != nil {
		return false, false, err
	}

	return false, c == ".", nil
}
