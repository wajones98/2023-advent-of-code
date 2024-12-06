package day6

import (
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
	data := NewMap(width, height)

	t.Run("Check that the underlying slice is the correct length", func(t *testing.T) {
		var expectedLength uint = uint(len(data.Map))
		actualLength := width * height
		if expectedLength < actualLength {
			t.Errorf("Expected: %d, Got: %d\n", expectedLength, actualLength)
		}
	})

	t.Run("Check that Put works as expected", func(t *testing.T) {
		err := data.Put(0, 1, 'd')
		if err != nil {
			t.Error(err)
		}

		if data.Map[4] != 'd' {
			t.Errorf("Expected %s, Got %s\n", string('d'), string(data.Map[9]))
		}
	})
}
