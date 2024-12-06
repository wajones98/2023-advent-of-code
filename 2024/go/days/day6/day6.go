package day6

import (
	"bufio"
	"fmt"

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

type Map struct {
	Map    []rune
	Width  uint
	Height uint
}

func NewMap(width, height uint) *Map {
	return &Map{
		Map:    make([]rune, width*height),
		Width:  width,
		Height: height,
	}
}

func (m *Map) Put(x, y uint, r rune) error {
	err := m.checkBounds(x, y)
	if err != nil {
		return err
	}
	m.Map[m.getIndex(x, y)] = r
	return nil
}

func (m *Map) Get(x, y uint) (rune, error) {
	err := m.checkBounds(x, y)
	if err != nil {
		return -1, err
	}
	return m.Map[m.getIndex(x, y)], nil
}

func (m *Map) getIndex(x, y uint) uint {
	return y*m.Width + x
}

func (m *Map) checkBounds(x, y uint) error {
	if x > m.Width {
		return fmt.Errorf("%d is out of bounds %d", x, m.Width)
	} else if y > m.Height {
		return fmt.Errorf("%d is out of bounds %d", y, m.Height)
	}

	return nil
}

func LoadInput(s *bufio.Scanner) (*Map, error) {
	return nil, nil
}
