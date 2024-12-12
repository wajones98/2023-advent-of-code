package day12

import (
	"bufio"
	"reflect"
	"strings"
	"testing"

	"github.com/wajones98/advent-of-code/common"
)

const Input = `AAAA
BBCD
BBCC
EEEC`

var Data = common.TwoDMap[string]{
	Width:  4,
	Height: 4,
	Map:    []string{"A", "A", "A", "A", "B", "B", "C", "D", "B", "B", "C", "C", "E", "E", "E", "C"},
}

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(Data, *actual) {
		t.Errorf("Expected: %v, Actual: %v\n", Data, actual)
	}
}

func TestFindPlantGroups(t *testing.T) {
	expected := map[string][][]Coords{
		"A": [][]Coords{{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
		"B": [][]Coords{{{0, 1}, {1, 1}, {0, 2}, {1, 2}}},
		"C": [][]Coords{{{2, 1}, {2, 2}, {3, 2}, {3, 3}}},
		"D": [][]Coords{{{1, 3}}},
		"E": [][]Coords{{{3, 0}, {1, 3}, {2, 3}}},
	}

	actual := FindPlantGroups(&Data)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v\n", Data, actual)
	}
}
