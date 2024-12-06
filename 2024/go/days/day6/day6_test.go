package day6

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

const Input = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestMap(t *testing.T) {
	var width, height uint = 4, 4
	twoDMap := NewTwoDMap(width, height)

	t.Run("Check that the underlying slice is the correct length", func(t *testing.T) {
		var expectedLength uint = uint(len(twoDMap.Map))
		actualLength := width * height
		if expectedLength < actualLength {
			t.Errorf("Expected: %d, Got: %d\n", expectedLength, actualLength)
		}
	})

	t.Run("Check that Put works as expected", func(t *testing.T) {
		err := twoDMap.Put(0, 1, "d")
		if err != nil {
			t.Error(err)
		}

		if twoDMap.Map[4] != "d" {
			t.Errorf("Expected %s, Got %s\n", "d", twoDMap.Map[4])
		}
	})

	t.Run("Check that Get works as expected", func(t *testing.T) {
		s, err := twoDMap.Get(0, 1)
		if err != nil {
			t.Error(err)
		}

		if s != "d" {
			t.Errorf("Expected %s, Got %s\n", "d", s)
		}
	})

	t.Run("Check FindPosition", func(t *testing.T) {
		x, y := twoDMap.FindPosition(4)

		if x != 0 {
			t.Errorf("Expected %d, Got %d\n", 0, x)
		} else if y != 1 {
			t.Errorf("Expected %d, Got %d\n", 1, y)
		}
	})
}

func TestLoadInput(t *testing.T) {
	const smallerInput = `..#
#.#
..^`
	s := bufio.NewScanner(strings.NewReader(smallerInput))

	expectedSlice := []string{".", ".", "#", "#", ".", "#", ".", ".", "^"}
	actualTwoDMap, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	if !slices.Equal(expectedSlice, actualTwoDMap.Map) {
		t.Errorf("Expected %v, Got %v\n", expectedSlice, actualTwoDMap.Map)
	}
}

func TestFindGuard(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	twoDMap, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	x, y, err := FindGuard(twoDMap)
	if err != nil {
		t.Error(err)
	}

	if x != 4 {
		t.Errorf("Expected %d, Got %d\n", 0, x)
	} else if y != 6 {
		t.Errorf("Expected %d, Got %d\n", 1, y)
	}
}
