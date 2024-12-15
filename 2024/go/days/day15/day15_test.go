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

var TwoDMap = &common.TwoDMap[Tile]{
	Width:  8,
	Height: 8,
	Map: []Tile{
		Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall,
		Wall, Empty, Empty, Box, Empty, Box, Empty, Wall,
		Wall, Wall, Robot, Empty, Box, Empty, Empty, Wall,
		Wall, Empty, Empty, Empty, Box, Empty, Empty, Wall,
		Wall, Empty, Wall, Empty, Box, Empty, Empty, Wall,
		Wall, Empty, Empty, Empty, Box, Empty, Empty, Wall,
		Wall, Empty, Empty, Empty, Empty, Empty, Empty, Wall,
		Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall,
	},
}

const Moves = ""

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, _, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(*TwoDMap, *actual) {
		t.Errorf("\nExpected:\n%v\nActual:\n%v\n", TwoDMap, actual)
	}
}
