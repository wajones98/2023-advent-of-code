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

	tests := []struct {
		Length    int
		Operators []int
		Expected  [][]int
	}{
		{
			Length:    2,
			Operators: []int{Add, Multiply},
			Expected: [][]int{
				{Add, Add},
				{Add, Multiply},
				{Multiply, Add},
				{Multiply, Multiply},
			},
		},
		{
			Length:    2,
			Operators: []int{Add, Multiply, Combine},
			Expected: [][]int{
				{Add, Add},
				{Add, Multiply},
				{Add, Combine},
				{Multiply, Add},
				{Multiply, Multiply},
				{Multiply, Combine},
				{Combine, Add},
				{Combine, Multiply},
				{Combine, Combine},
			},
		},
		{
			Length:    3,
			Operators: []int{Add, Multiply},
			Expected: [][]int{
				{Add, Add, Add},
				{Add, Add, Multiply},
				{Add, Multiply, Add},
				{Add, Multiply, Multiply},
				{Multiply, Add, Add},
				{Multiply, Add, Multiply},
				{Multiply, Multiply, Add},
				{Multiply, Multiply, Multiply},
			},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("For length %d", test.Length), func(t *testing.T) {
			actual := GenerateCombinations(test.Length, test.Operators)
			if len(test.Expected) != len(actual) {
				t.Errorf("Expected: %d\nGot: %d\n", len(test.Expected), len(actual))
			}

			for i, a := range actual {
				if !reflect.DeepEqual(a, test.Expected[i]) {
					t.Errorf("Expected: %v\nGot: %v\n", test.Expected[i], a)
				}
			}
		})
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
				{Add},
				{Multiply},
				{Combine},
			},
			Equation: Equation{
				Result: 156,
				Values: []int{15, 6},
			},
			Expected: []bool{false, false, true},
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
		{
			Combinations: [][]int{
				{Add, Add, Add},
				{Add, Add, Multiply},
				{Add, Multiply, Add},
				{Add, Multiply, Multiply},
				{Multiply, Add, Add},
				{Multiply, Add, Multiply},
				{Multiply, Multiply, Add},
				{Multiply, Multiply, Multiply},
			},
			Equation: Equation{
				Result: 292,
				Values: []int{11, 6, 16, 20},
			},
			Expected: []bool{false, false, true, false, false, false, false, false},
		},
		{
			Combinations: [][]int{
				{Add, Add, Add},
				{Add, Add, Multiply},
				{Add, Add, Combine},
				{Add, Multiply, Add},
				{Add, Multiply, Multiply},
				{Add, Multiply, Combine},
				{Add, Combine, Add},
				{Add, Combine, Multiply},
				{Add, Combine, Combine},
				{Multiply, Add, Add},
				{Multiply, Add, Multiply},
				{Multiply, Add, Combine},
				{Multiply, Multiply, Add},
				{Multiply, Multiply, Multiply},
				{Multiply, Multiply, Combine},
				{Multiply, Combine, Add},
				{Multiply, Combine, Multiply},
				{Multiply, Combine, Combine},
				{Combine, Add, Add},
				{Combine, Add, Multiply},
				{Combine, Add, Combine},
				{Combine, Multiply, Add},
				{Combine, Multiply, Multiply},
				{Combine, Multiply, Combine},
				{Combine, Combine, Add},
				{Combine, Combine, Multiply},
				{Combine, Combine, Combine},
			},
			Equation: Equation{
				Result: 7290,
				Values: []int{6, 8, 6, 15},
			},
			Expected: []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, false},
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

func TestTotal(t *testing.T) {
	expected := 3749
	actual := GetPart1Total(Data)

	if expected != actual {
		t.Errorf("Expected: %d\nGot: %d\n", expected, actual)
	}
}

func TestTotalPartTwo(t *testing.T) {
	expected := 11387
	actual, err := GetPart2Total(Data)
	if err != nil {
		t.Error(err)
	}

	if expected != actual {
		t.Errorf("Expected: %d\nGot: %d\n", expected, actual)
	}
}
