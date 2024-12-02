package day2

import (
	"errors"
	"log"
	"testing"
)

func TestDetermineDirection(t *testing.T) {
	tests := []struct {
		Title         string
		Left          uint64
		Right         uint64
		ExpectedValue Direction
		ExpectedErr   error
	}{
		{
			Title:         "Left < Right returns Increasing",
			Left:          1,
			Right:         2,
			ExpectedValue: Increasing,
			ExpectedErr:   nil,
		},
		{

			Title:         "Left > Right returns Decreasing",
			Left:          2,
			Right:         1,
			ExpectedValue: Decreasing,
			ExpectedErr:   nil,
		},
		{

			Title:         "Left == Right returns error",
			Left:          1,
			Right:         1,
			ExpectedValue: false,
			ExpectedErr:   DirectionError,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			actual, err := determineDirection(test.Left, test.Right)
			if actual != test.ExpectedValue {
				log.Fatalf("Value Error -> Expected: %v, Actual: %v", test.ExpectedValue, actual)
			} else if !errors.Is(err, test.ExpectedErr) {
				log.Fatalf("Error Error -> Expected: %v, Actual: %v", test.ExpectedErr, err)
			}
		})
	}
}
