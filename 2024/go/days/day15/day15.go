package day15

import (
	"bufio"
	"fmt"
	"time"

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
	s, closeFile, err := input.GetInput(Day)
	if err != nil {
		return 0, err
	}
	defer closeFile()

	data, err := LoadInput(s)
	if err != nil {
		return 0, err
	}

	for _, m := range data.Moves {
		data.MoveRobot(m)
		fmt.Print("\033[H\033[2J")
		fmt.Print(data)
		time.Sleep(time.Millisecond * 100)
	}

	return data.Sum(), nil
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
type Direction struct {
	X, Y int
}

const (
	TileWall  Tile = '#'
	TileRobot      = '@'
	TileBox        = 'O'
	TileEmpty      = '.'

	MoveLeft  Move = '<'
	MoveRight      = '>'
	MoveUp         = '^'
	MoveDown       = 'v'
)

var Directions = map[Move]Direction{
	MoveUp:    {0, -1},
	MoveDown:  {0, 1},
	MoveLeft:  {-1, 0},
	MoveRight: {1, 0},
}

type Robot struct {
	X, Y int
}

type Data struct {
	Robot   *Robot
	Moves   []Move
	TwoDMap *common.TwoDMap[Tile]
}

func (d *Data) MoveRobot(m Move) {
	dir := Directions[m]

	newX, newY := d.Robot.X+dir.X, d.Robot.Y+dir.Y
	t, _ := d.TwoDMap.Get(newX, newY)
	switch t {
	case TileWall:
		return
	case TileEmpty:
		d.TwoDMap.Put(d.Robot.X, d.Robot.Y, TileEmpty)
		d.TwoDMap.Put(newX, newY, TileRobot)
		d.Robot.X = newX
		d.Robot.Y = newY
	case TileBox:
		endCoords := Direction{}
		canMove := false

		x, y := newX, newY
	BoxLoop:
		for {
			x += dir.X
			y += dir.Y

			t, _ := d.TwoDMap.Get(x, y)
			switch t {
			case TileWall:
				canMove = false
				break BoxLoop
			case TileEmpty:
				canMove = true
				endCoords.X = x
				endCoords.Y = y
				break BoxLoop
			}
		}

		if !canMove {
			return
		}

		d.TwoDMap.Put(d.Robot.X, d.Robot.Y, TileEmpty)
		d.TwoDMap.Put(newX, newY, TileRobot)
		d.Robot.X = newX
		d.Robot.Y = newY

		d.TwoDMap.Put(endCoords.X, endCoords.Y, TileBox)
	}
}

func (d *Data) Sum() int {
	sum := 0
	for i, t := range d.TwoDMap.Map {
		if t != TileBox {
			continue
		}
		x, y := d.TwoDMap.FindPosition(i)

		sum += 100*y + x
	}

	return sum
}

func (d *Data) String() string {
	result := ""

	for i, c := range d.TwoDMap.Map {
		result += fmt.Sprintf("%s", string(c))
		x := (i + 1) % int(d.TwoDMap.Width)
		if x == 0 {
			result += "\n"
		}
	}
	result += "\n"

	return result
}

func LoadInput(s *bufio.Scanner) (*Data, error) {
	data := &Data{
		Moves: []Move{},
	}
	lines := []string{}
	isMoves := false
	for s.Scan() {
		line := s.Text()
		if line == "" {
			isMoves = true
			continue
		}
		if isMoves {
			for _, m := range line {
				data.Moves = append(data.Moves, Move(m))
			}
		} else {

			lines = append(lines, line)
		}
	}

	width, height := len(lines[0]), len(lines)
	twoDMap := common.NewTwoDMap[Tile](width, height)

	for y, line := range lines {
		for x, char := range line {
			tile := Tile(char)
			err := twoDMap.Put(x, y, tile)
			if err != nil {
				return nil, err
			}
			if tile == TileRobot {
				data.Robot = &Robot{x, y}
			}
		}
	}

	data.TwoDMap = twoDMap

	return data, nil
}
