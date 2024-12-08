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

var Frequencies map[string][]Coords = map[string][]Coords{
	"0": {{8, 1}, {5, 2}, {7, 3}, {4, 4}},
	"A": {{6, 5}, {8, 8}, {9, 9}},
}

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

	expected := Frequencies
	actual := FindFrequencies(twoDMap)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Got: %v\n", expected, actual)
	}
}

func TestFindAntinodes(t *testing.T) {
	pOne := Coords{4, 3}
	pTwo := Coords{5, 5}

	expected := []Coords{{3, 1}, {6, 7}}
	actual := FindAntinodes(pOne, pTwo)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Got: %v\n", expected, actual)
	}
}
