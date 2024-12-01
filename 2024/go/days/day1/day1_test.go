package day1

import (
	"log"
	"testing"
)

func TestPart2(t *testing.T) {
	leftList := []int{3, 4, 2, 1, 3, 3}
	rightList := []int{4, 3, 5, 3, 9, 3}

	t.Run("Total is correct", func(t *testing.T) {
		expected := 31
		actual := Part2(leftList, rightList)

		if expected != actual {
			log.Fatalf("Got %d, expected %d", actual, expected)
		}
	})
}
