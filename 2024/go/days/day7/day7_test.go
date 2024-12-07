package day7

import (
	"bufio"
	"fmt"
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
		Result: 190,
		Values: []int{10, 19},
	},
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
		Values: []int{6, 8, 6, 15},
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

	if len(Data) != len(actual) {
		t.Errorf("Expected: %d\nGot: %d\n", len(Data), len(actual))
	}

	for i, a := range actual {
		if !reflect.DeepEqual(a, Data[i]) {
			t.Errorf("Expected: %v\nGot: %v\n", Data[i], a)
		}
	}
}

func TestGenerateCombinations(t *testing.T) {
	length := 2
	actual := GenerateCombinations(length)
	expected := [][]int{
		{Add, Add},
		{Add, Multiply},
		{Multiply, Add},
		{Multiply, Multiply},
	}

	if len(expected) != len(actual) {
		t.Errorf("Expected: %d\nGot: %d\n", len(expected), len(actual))
	}

	for i, a := range actual {
		if !reflect.DeepEqual(a, expected[i]) {
			t.Errorf("Expected: %v\nGot: %v\n", expected[i], a)
		}
	}
}

func TestEquationIsValid(t *testing.T) {
	tests := []struct {
		Combinations [][]int
		Equation     Equation
		Expected     []bool
	}{
		{
			Combinations: [][]int{
				{Add},
				{Multiply},
			},
			Equation: Equation{
				Result: 190,
				Values: []int{10, 19},
			},
			Expected: []bool{false, true},
		},
		{
			Combinations: [][]int{
				{Add, Add},
				{Add, Multiply},
				{Multiply, Add},
				{Multiply, Multiply},
			},
			Equation: Equation{
				Result: 3267,
				Values: []int{81, 40, 27},
			},
			Expected: []bool{false, true, true, false},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.Equation), func(t *testing.T) {
			for i, c := range test.Combinations {
				actual := test.Equation.IsValid(c)
				if actual != test.Expected[i] {
					t.Errorf("\nExpected: %t\nGot: %t\n", test.Expected[i], actual)
				}
			}
		})
	}

}
