package day2

import (
	"errors"
	"fmt"
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

func TestIsSafe(t *testing.T) {
	tests := []struct {
		Title         string
		Left          uint64
		Right         uint64
		Direction     Direction
		ExpectedValue bool
	}{
		{
			Title:         "Decreasing with acceptable value",
			Left:          7,
			Right:         6,
			Direction:     Decreasing,
			ExpectedValue: true,
		},
		{
			Title:         "Increasing with acceptable value",
			Left:          6,
			Right:         7,
			Direction:     Increasing,
			ExpectedValue: true,
		},
		{
			Title:         "Value decreasing when Increasing",
			Left:          6,
			Right:         5,
			Direction:     Increasing,
			ExpectedValue: false,
		},
		{
			Title:         "Value increasing when Decreasing",
			Left:          5,
			Right:         6,
			Direction:     Decreasing,
			ExpectedValue: false,
		},
		{
			Title:         "Increasing but more than 3",
			Left:          1,
			Right:         5,
			Direction:     Increasing,
			ExpectedValue: false,
		},
		{
			Title:         "Decreasing but more than 3",
			Left:          5,
			Right:         1,
			Direction:     Decreasing,
			ExpectedValue: false,
		},
		{
			Title:         "Decreasing but same value",
			Left:          1,
			Right:         1,
			Direction:     Decreasing,
			ExpectedValue: false,
		},
		{
			Title:         "Increasing but same value",
			Left:          1,
			Right:         1,
			Direction:     Increasing,
			ExpectedValue: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			actual := isSafe(test.Left, test.Right, test.Direction)
			if actual != test.ExpectedValue {
				log.Fatalf("Value Error -> Expected: %v, Actual: %v", test.ExpectedValue, actual)
			}
		})
	}
}

func TestExampleData(t *testing.T) {
	tests := []struct {
		Data          []uint64
		ExpectedValue bool
	}{
		{
			Data:          []uint64{7, 6, 4, 2, 1},
			ExpectedValue: true,
		},
		{
			Data:          []uint64{1, 2, 7, 8, 9},
			ExpectedValue: false,
		},
		{
			Data:          []uint64{9, 7, 6, 2, 1},
			ExpectedValue: false,
		},
		{
			Data:          []uint64{1, 3, 2, 4, 5},
			ExpectedValue: false,
		},
		{
			Data:          []uint64{8, 6, 4, 4, 1},
			ExpectedValue: false,
		},
		{
			Data:          []uint64{1, 3, 6, 7, 9},
			ExpectedValue: true,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v", test.Data)
		t.Run(testName, func(t *testing.T) {
			actual := reportIsSafe(test.Data)
			if actual != test.ExpectedValue {
				log.Fatalf("Value Error -> Expected: %v, Actual: %v", test.ExpectedValue, actual)
			}
		})
	}
}
