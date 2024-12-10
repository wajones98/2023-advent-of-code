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

	expected := []Coords{{0, 6}, {6, 6}}
	actual := PossiblePaths(twoDMap, 0, 3, 0)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Got: %v\n", expected, actual)
	}

}
