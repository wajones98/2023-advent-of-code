package day12

import (
	"bufio"
	"fmt"
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
		"B": [][]Coords{{{0, 1}, {0, 2}, {1, 2}, {1, 1}}},
		"C": [][]Coords{{{2, 1}, {2, 2}, {3, 2}, {3, 3}}},
		"D": [][]Coords{{{3, 1}}},
		"E": [][]Coords{{{0, 3}, {1, 3}, {2, 3}}},
	}

	actual := FindPlantGroups(&Data)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v\n                Actual: %v\n", expected, actual)
	}
}

func TestCalculatePerimeter(t *testing.T) {
	tests := []struct {
		Input    []Coords
		Expected int
	}{
		{
			Input:    []Coords{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
			Expected: 10,
		},
		{
			Input:    []Coords{{0, 1}, {0, 2}, {1, 2}, {1, 1}},
			Expected: 8,
		},
		{
			Input:    []Coords{{2, 1}, {2, 2}, {3, 2}, {3, 3}},
			Expected: 10,
		},
		{
			Input:    []Coords{{3, 1}},
			Expected: 4,
		},
		{
			Input:    []Coords{{0, 3}, {1, 3}, {2, 3}},
			Expected: 8,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.Input), func(t *testing.T) {
			actual := CalculatePerimeter(test.Input)
			if test.Expected != actual {
				t.Errorf("Expected %d, Actual: %d\n", test.Expected, actual)
			}
		})
	}
}

func TestCalculateSides(t *testing.T) {
	tests := []struct {
		Input    []Coords
		Expected int
	}{
		{
			Input:    []Coords{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
			Expected: 4,
		},
		{
			Input:    []Coords{{0, 1}, {0, 2}, {1, 2}, {1, 1}},
			Expected: 4,
		},
		{
			Input:    []Coords{{2, 1}, {2, 2}, {3, 2}, {3, 3}},
			Expected: 8,
		},
		{
			Input:    []Coords{{3, 1}},
			Expected: 4,
		},
		{
			Input:    []Coords{{0, 3}, {1, 3}, {2, 3}},
			Expected: 4,
		},
		{
			Input:    []Coords{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}, {0, 3}, {0, 4}, {1, 4}, {2, 4}, {3, 4}, {4, 4}},
			Expected: 12,
		},
		{
			Input: []Coords{
				{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0},
				{0, 1}, {1, 1}, {2, 1}, {5, 1},
				{0, 2}, {1, 2}, {2, 2}, {5, 2},
				{0, 3}, {3, 3}, {4, 3}, {5, 3},
				{0, 4}, {3, 4}, {4, 4}, {5, 4},
				{0, 5}, {1, 5}, {2, 5}, {3, 5}, {4, 5},
			},
			Expected: 12,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.Input), func(t *testing.T) {
			actual := CalculateSides(test.Input)
			if test.Expected != actual {
				t.Errorf("Expected %d, Actual: %d\n", test.Expected, actual)
			}
		})
	}
}

func TestCalculatePrice(t *testing.T) {
	tests := []struct {
		Input    map[string][][]Coords
		Expected int
	}{
		{
			Input: map[string][][]Coords{
				"E": [][]Coords{{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}, {0, 3}, {0, 4}, {1, 4}, {2, 4}, {3, 4}, {4, 4}}},
				"X": [][]Coords{{{1, 3}, {2, 3}, {3, 3}, {4, 3}}, {{1, 1}, {2, 1}, {3, 1}, {4, 1}}},
			},
			Expected: 236,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.Input), func(t *testing.T) {
			actual := CalculatePrice(test.Input, true)
			if test.Expected != actual {
				t.Errorf("Expected %d, Actual: %d\n", test.Expected, actual)
			}
		})
	}
}

func TestFullExample(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

	s := bufio.NewScanner(strings.NewReader(input))
	twoDMap, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	plantGroups := FindPlantGroups(twoDMap)

	actual := CalculatePrice(plantGroups, true)
	expected := 1206

	if expected != actual {
		t.Errorf("Expected %d, Actual: %d\n", expected, actual)
	}
}
