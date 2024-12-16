package day15

import (
	"bufio"

	"github.com/wajones98/advent-of-code/common"
	"github.com/wajones98/advent-of-code/days"
	"github.com/wajones98/advent-of-code/input"
)

const Day int = 15

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

type Tile rune
type Move rune

const (
	Wall  Tile = '#'
	Robot      = '@'
	Box        = 'O'
	Empty      = '.'

	Left  Move = '<'
	Right      = '>'
	Up         = '^'
	Down       = 'v'
)

type Coords struct {
	X, Y int
}

type Data struct {
	Robot   *Coords
	Moves   []Move
	TwoDMap *common.TwoDMap[Tile]
}

func LoadInput(s *bufio.Scanner) (*Data, error) {
	data := &Data{
		Moves: []Move{},
	}
	lines := []string{}
	for s.Scan() {
		line := s.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	moves := lines[len(lines)-1]
	for _, m := range moves {
		data.Moves = append(data.Moves, Move(m))
	}

	lines = lines[:len(lines)-1]

	width, height := len(lines[0]), len(lines)
	twoDMap := common.NewTwoDMap[Tile](width, height)

	for y, line := range lines {
		for x, char := range line {
			tile := Tile(char)
			err := twoDMap.Put(x, y, tile)
			if err != nil {
				return nil, err
			}
			if tile == Robot {
				data.Robot = &Coords{x, y}
			}
		}
	}

	data.TwoDMap = twoDMap

	return data, nil
}
