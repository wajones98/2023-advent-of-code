package day15

import (
	"bufio"
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
	_, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}
}
