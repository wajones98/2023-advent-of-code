package day10

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

const Input = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestLoadInput(t *testing.T) {
	input := `0123
1234
8765
9876`

	s := bufio.NewScanner(strings.NewReader(input))
	twoDMap, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}
	if strings.ReplaceAll(twoDMap.String(), "\n", "") != strings.ReplaceAll(input, "\n", "") {
		t.Errorf("\nExpected: \n%s\nGot: \n%s\n", input, twoDMap.String())
	}
}

func TestTraverseTrail(t *testing.T) {
	input := `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`

	s := bufio.NewScanner(strings.NewReader(input))
	twoDMap, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		Expected  bool
		Coords    Coords
		Direction Direction
		Value     int
	}{
		{
			Expected:  false,
			Coords:    Coords{3, 0},
			Direction: Up,
			Value:     0,
		},
		{
			Expected:  true,
			Coords:    Coords{3, 0},
			Direction: Down,
			Value:     0,
		},
		{
			Expected:  true,
			Coords:    Coords{3, 3},
			Direction: Left,
			Value:     3,
		},
		{
			Expected:  true,
			Coords:    Coords{3, 3},
			Direction: Right,
			Value:     3,
		},
		{
			Expected:  false,
			Coords:    Coords{3, 3},
			Direction: Up,
			Value:     3,
		},
		{
			Expected:  false,
			Coords:    Coords{0, 6},
			Direction: Left,
			Value:     6,
		},
		{
			Expected:  false,
			Coords:    Coords{6, 5},
			Direction: Right,
			Value:     8,
		},
	}

	for _, test := range tests {
		directionName := ""
		switch test.Direction {
		case Up:
			directionName = "Up"
		case Down:
			directionName = "Down"
		case Left:
			directionName = "Left"
		case Right:
			directionName = "Right"
		}

		t.Run(fmt.Sprintf("X: %d, Y: %d, Direction: %s, Value: %d", test.Coords.X, test.Coords.Y, directionName, test.Value), func(t *testing.T) {
			_, ok := TraverseTrail(twoDMap, test.Coords.X, test.Coords.Y, test.Value, test.Direction)
			if ok != test.Expected {
				t.Errorf("Expected %t, Got: %t\n", test.Expected, ok)
			}
		})
	}

}

func TestFindTrails(t *testing.T) {

	tests := []struct {
		Input    string
		Expected map[Coords]int
		X, Y     int
	}{
		{
			Input: `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`,
			Expected: map[Coords]int{{0, 6}: 1, {6, 6}: 1},
			X:        3,
			Y:        0,
		},
		{
			Input: `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`,
			Expected: map[Coords]int{{6, 0}: 1, {5, 1}: 1, {4, 4}: 1, {0, 6}: 10},
			X:        3,
			Y:        0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			s := bufio.NewScanner(strings.NewReader(test.Input))
			twoDMap, err := LoadInput(s)
			if err != nil {
				t.Error(err)
			}

			actual := PossiblePaths(twoDMap, test.X, test.Y, 0)

			if !reflect.DeepEqual(test.Expected, actual) {
				t.Errorf("Expected: %v, Got: %v\n", test.Expected, actual)
			}
		})
	}

}
