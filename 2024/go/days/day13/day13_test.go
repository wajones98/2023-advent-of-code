package day13

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

const Input = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

var Data = []Prize{
	{
		Location: Coords{
			X: 8400,
			Y: 5400,
		},
		ButtonA: Coords{
			X: 94,
			Y: 34,
		},
		ButtonB: Coords{
			X: 22,
			Y: 67,
		},
	},
	{
		Location: Coords{
			X: 12748,
			Y: 12176,
		},
		ButtonA: Coords{
			X: 26,
			Y: 66,
		},
		ButtonB: Coords{
			X: 67,
			Y: 21,
		},
	},
	{
		Location: Coords{
			X: 7870,
			Y: 6450,
		},
		ButtonA: Coords{
			X: 17,
			Y: 86,
		},
		ButtonB: Coords{
			X: 84,
			Y: 37,
		},
	},
	{
		Location: Coords{
			X: 18641,
			Y: 10279,
		},
		ButtonA: Coords{
			X: 69,
			Y: 23,
		},
		ButtonB: Coords{
			X: 27,
			Y: 71,
		},
	},
}

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(Data, actual) {
		t.Errorf("Expected: %v, Actual: %v\\n", Data, actual)
	}
}

type PossibleCombinationsInput struct {
	Location, A, B int
}

func TestPossibleCombinations(t *testing.T) {
	tests := []struct {
		Input    PossibleCombinationsInput
		Expected map[int]int
	}{
		{
			Input: PossibleCombinationsInput{
				Location: 8400,
				A:        94,
				B:        22,
			},
			Expected: map[int]int{
				3:  369,
				14: 322,
				25: 275,
				36: 228,
				47: 181,
				58: 134,
				69: 87,
				80: 40,
			},
		},
		{
			Input: PossibleCombinationsInput{
				Location: 5400,
				A:        34,
				B:        67,
			},
			Expected: map[int]int{
				13:  74,
				80:  40,
				147: 6,
			},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.Input), func(t *testing.T) {
			actual, _ := PossibleCombinations(test.Input.Location, test.Input.A, test.Input.B)
			if !reflect.DeepEqual(test.Expected, actual) {
				t.Errorf("Expected: %v, Actual: %v\n", test.Expected, actual)
			}
		})
	}
}

func TestFindCheapestCombinations(t *testing.T) {
	tests := []struct {
		Input    []map[int]int
		Expected int
	}{
		{
			Input: []map[int]int{
				{
					3:  369,
					14: 322,
					25: 275,
					36: 228,
					47: 181,
					58: 134,
					69: 87,
					80: 40,
				},
				{
					13:  74,
					80:  40,
					147: 6,
				},
			},
			Expected: 280,
		},
	}

	for _, test := range tests {
		actual := FindCheapestCombination(test.Input[0], test.Input[0])
		if test.Expected != actual {
			t.Errorf("Expected: %d, Actual: %d\n", test.Expected, actual)
		}
	}
}

func TestFindTokenCost(t *testing.T) {
	tests := []struct {
		Input    Prize
		Expected int
		Offset   int
	}{
		{
			Input:    Data[0],
			Expected: 280,
			Offset:   0,
		},
		{
			Input:    Data[1],
			Expected: 0,
			Offset:   0,
		},
		{
			Input:    Data[2],
			Expected: 200,
			Offset:   0,
		},
		{
			Input:    Data[3],
			Expected: 0,
			Offset:   0,
		},
		{
			Input:    Data[0],
			Expected: 0,
			Offset:   10000000000000,
		},
		{
			Input:    Data[1],
			Expected: 459236326669,
			Offset:   10000000000000,
		},
		{
			Input:    Data[2],
			Expected: 0,
			Offset:   10000000000000,
		},
		{
			Input:    Data[3],
			Expected: 416082282239,
			Offset:   10000000000000,
		},
	}

	for _, test := range tests {
		actual := FindTokenCost(test.Input, test.Offset)
		if test.Expected != actual {
			t.Errorf("Expected: %d, Actual: %d\n", test.Expected, actual)
		}
	}
}

func TestTotalTokens(t *testing.T) {
	expected := 480
	actual := TotalTokens(Data, 0)
	if expected != actual {
		t.Errorf("Expected: %d, Actual: %d\n", expected, actual)
	}
}
