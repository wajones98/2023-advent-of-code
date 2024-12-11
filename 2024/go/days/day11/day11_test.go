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
