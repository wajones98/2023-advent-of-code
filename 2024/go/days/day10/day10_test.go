package day10

import (
	"bufio"
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
	actual := FindTrails(twoDMap)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Got: %v\n", expected, actual)
	}

}
