package day7

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

const Input string = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

var Data []Equation = []Equation{
	{
		Result: 3267,
		Values: []int{81, 40, 27},
	},
	{
		Result: 83,
		Values: []int{17, 5},
	},
	{
		Result: 156,
		Values: []int{15, 6},
	},
	{
		Result: 7290,
		Values: []int{68615},
	},
	{
		Result: 161011,
		Values: []int{16, 10, 13},
	},
	{
		Result: 192,
		Values: []int{17, 8, 14},
	},
	{
		Result: 21037,
		Values: []int{9, 7, 18, 13},
	},
	{
		Result: 292,
		Values: []int{11, 6, 16, 20},
	},
}

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(Data, actual) {
		t.Errorf("Expected: %v\nGot: %v\n", Data, actual)
	}
}
