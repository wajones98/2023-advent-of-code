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
	Up, Down, Left, Right Direction = "^", "v", "<", ">"
)

type Guard struct {
	X         uint
	Y         uint
	Direction Direction
}

var Directions []Direction = []Direction{Up, Down, Left, Right}

func Part1() int {
	_, closeFile, err := input.GetInput(Day)
	if err != nil {
		panic(err)
	}
	defer closeFile()

	return 0
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
			}, nil
		}
	}

	return nil, errors.New("Could not find guard")
}
