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

var TwoDMap = &common.TwoDMap[Tile]{
	Width:  8,
	Height: 8,
	Map:    []Tile{Wall},
}

const Moves = ""

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, moves, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Moves: %s\n", moves)
	fmt.Printf("Map: %s\n", actual)

	if !reflect.DeepEqual(Data, *actual) {
		t.Errorf("Expected: %v, Actual: %v\\n", Data, actual)
	}
}
