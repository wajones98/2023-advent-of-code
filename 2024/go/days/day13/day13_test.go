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

func TestPossibleCombinations(t *testing.T) {
	tests := []struct {
		Input    Prize
		Expected map[int]int
	}{
		{
			Input: Data[0],
			Expected: map[int]int{
				80: 40,
			},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.Input), func(t *testing.T) {
			actual, err := PossibleCombinations(test.Input)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(test.Expected, actual) {
				t.Errorf("Expected: %v\n, Actual: %v\n", test.Expected, actual)
			}
		})
	}
}
