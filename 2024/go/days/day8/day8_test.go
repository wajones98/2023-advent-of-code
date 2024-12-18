package day8

import (
	"bufio"
	"fmt"
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
	tests := []struct {
		POne     Coords
		PTwo     Coords
		Expected []Coords
	}{
		{
			POne:     Coords{4, 3},
			PTwo:     Coords{5, 5},
			Expected: []Coords{{3, 1}, {6, 7}},
		},
		{
			POne:     Coords{4, 3},
			PTwo:     Coords{8, 4},
			Expected: []Coords{{0, 2}, {12, 5}},
		},
		{
			POne:     Coords{8, 4},
			PTwo:     Coords{5, 5},
			Expected: []Coords{{11, 3}, {2, 6}},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("POne: %v, PTwo: %v\n", test.POne, test.PTwo), func(t *testing.T) {
			actual := FindAntinodes(test.POne, test.PTwo, 1)
			if !reflect.DeepEqual(test.Expected, actual) {
				t.Errorf("Expected: %v, Got: %v\n", test.Expected, actual)
			}
		})
	}
}

func TestIsValidAntinode(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	twoDMap, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		Coords   Coords
		Expected bool
	}{
		{
			Coords:   Coords{3, 1},
			Expected: true,
		},
		{
			Coords:   Coords{6, 7},
			Expected: true,
		},
		{
			Coords:   Coords{0, 2},
			Expected: true,
		},
		{
			Coords:   Coords{12, 5},
			Expected: false,
		},
		{
			Coords:   Coords{11, 3},
			Expected: true,
		},
		{
			Coords:   Coords{2, 6},
			Expected: true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v\n", test.Coords), func(t *testing.T) {
			actual := isValidAntinode(twoDMap.Width, twoDMap.Height, test.Coords)
			if test.Expected != actual {
				t.Errorf("Expected: %t, Got: %t\n", test.Expected, actual)
			}
		})
	}
}

func TestFindAllUniqueAntinodes(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	twoDMap, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	f := FindFrequencies(twoDMap)
	expected := 14
	actual := FindAllUniqueAntinodes(twoDMap, f, false)
	if expected != actual {
		t.Errorf("Expected: %d, Got: %d\n", expected, actual)
	}
}
