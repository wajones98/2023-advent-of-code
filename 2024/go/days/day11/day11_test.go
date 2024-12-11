package day11

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

const Input = `125 17`

var Data = []int{125, 17}

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(Data, actual) {
		t.Errorf("Expected: %v, Actual: %v\n", Data, actual)
	}
}

func TestTransformStone(t *testing.T) {
	tests := []struct {
		Stone    int
		Expected []int
	}{
		{0, []int{1}},
		{1, []int{2024}},
		{10, []int{1, 0}},
		{99, []int{9, 9}},
		{999, []int{2021976}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.Stone), func(t *testing.T) {
			actual := TransformStone(test.Stone)
			if !reflect.DeepEqual(test.Expected, actual) {
				t.Errorf("Expected: %v, Actual: %v\n", test.Expected, actual)
			}
		})
	}
}

func TestBlink(t *testing.T) {
	tests := []struct {
		Input    []int
		Expected []int
	}{
		{Input: []int{125, 17}, Expected: []int{253000, 1, 7}},
		{Input: []int{253000, 1, 7}, Expected: []int{253, 0, 2024, 14168}},
		{Input: []int{253, 0, 2024, 14168}, Expected: []int{512072, 1, 20, 24, 28676032}},
		{Input: []int{512072, 1, 20, 24, 28676032}, Expected: []int{512, 72, 2024, 2, 0, 2, 4, 2867, 6032}},
		{Input: []int{512, 72, 2024, 2, 0, 2, 4, 2867, 6032}, Expected: []int{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32}},
		{Input: []int{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32}, Expected: []int{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.Input), func(t *testing.T) {
			actual := Blink(test.Input)
			if !reflect.DeepEqual(test.Expected, actual) {
				t.Errorf("Expected: %v, Actual: %v\n", test.Expected, actual)
			}
		})
	}
}

func TestGetStoneCount(t *testing.T) {
	tests := []struct {
		BlinkCount int
		Expected   int
	}{
		{6, 22}, {25, 58159},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.BlinkCount), func(t *testing.T) {
			actual := GetStoneCount(test.BlinkCount, Data)
			if test.Expected != actual {
				t.Errorf("Expected: %d, Actual: %d\n", test.Expected, actual)
			}
		})
	}

}
