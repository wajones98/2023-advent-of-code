package day15

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/wajones98/advent-of-code/common"
)

const Input = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

var LoadedData = Data{
	TwoDMap: &common.TwoDMap[Tile]{
		Width:  8,
		Height: 8,
		Map: []Tile{
			TileWall, TileWall, TileWall, TileWall, TileWall, TileWall, TileWall, TileWall,
			TileWall, TileEmpty, TileEmpty, TileBox, TileEmpty, TileBox, TileEmpty, TileWall,
			TileWall, TileWall, TileRobot, TileEmpty, TileBox, TileEmpty, TileEmpty, TileWall,
			TileWall, TileEmpty, TileEmpty, TileEmpty, TileBox, TileEmpty, TileEmpty, TileWall,
			TileWall, TileEmpty, TileWall, TileEmpty, TileBox, TileEmpty, TileEmpty, TileWall,
			TileWall, TileEmpty, TileEmpty, TileEmpty, TileBox, TileEmpty, TileEmpty, TileWall,
			TileWall, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileEmpty, TileWall,
			TileWall, TileWall, TileWall, TileWall, TileWall, TileWall, TileWall, TileWall,
		},
	},
	Robot: &Robot{2, 2},
	Moves: []Move{MoveLeft, MoveUp, MoveUp, MoveRight, MoveRight, MoveRight, MoveDown, MoveDown, MoveLeft, MoveDown, MoveRight, MoveRight, MoveDown, MoveLeft, MoveLeft},
}

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(*LoadedData.TwoDMap, *actual.TwoDMap) {
		t.Errorf("\nExpected:\n%v\nActual:\n%v\n", *LoadedData.TwoDMap, *actual.TwoDMap)
	}

	if !reflect.DeepEqual(*LoadedData.Robot, *actual.Robot) {
		t.Errorf("\nExpected:\n%v\nActual:\n%v\n", *LoadedData.Robot, *actual.Robot)
	}

	if !reflect.DeepEqual(LoadedData.Moves, actual.Moves) {
		t.Errorf("\nExpected:\n%v\nActual:\n%v\n", LoadedData.Moves, actual.Moves)
	}
}

func TestMoveRobot(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	data, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	for _, m := range data.Moves {
		data.MoveRobot(m)
		fmt.Printf("%s", data.String())
	}

	fmt.Printf("Sum: %d\n", data.Sum())
}

func TestLargerInput(t *testing.T) {
	largerInput := `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

	s := bufio.NewScanner(strings.NewReader(largerInput))
	data, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	for _, m := range data.Moves {
		data.MoveRobot(m)
	}

	fmt.Printf("%s", data.String())
	fmt.Printf("Sum: %d\n", data.Sum())
}
