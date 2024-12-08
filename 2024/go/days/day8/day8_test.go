package day8

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

const Input = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	twoDMap, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}
	if strings.ReplaceAll(twoDMap.String(), "\n", "") != strings.ReplaceAll(Input, "\n", "") {
		t.Errorf("\nExpected: \n%s\nGot: \n%s\n", Input, twoDMap.String())
	}
}

func TestFindFrequencies(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	twoDMap, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	expected := map[string][]Coords{
		"0": {{8, 1}, {5, 2}, {7, 3}, {4, 4}},
		"A": {{6, 5}, {8, 8}, {9, 9}},
	}

	actual := FindFrequencies(twoDMap)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Got: %v\n", expected, actual)
	}
}
