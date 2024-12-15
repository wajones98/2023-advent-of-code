package day14

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

const Input = `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`

var Data = []Robot{
	{0, 4, 3, -3}, {6, 3, -1, -3}, {10, 3, -1, 2}, {2, 0, 2, -1}, {0, 0, 1, 3}, {3, 0, -2, -2},
	{7, 6, -1, -3}, {3, 0, -1, -2}, {9, 3, 2, 3}, {7, 3, -1, 2}, {2, 4, 2, -3}, {9, 5, -3, -3},
}

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, err := LoadInput(s, 11, 7)
	if err != nil {
		t.Error(err)
	}

	if len(Data) != len(actual) {
		t.Errorf("\nExpected: %v\nActual: %v\n", len(Data), len(actual))
	}

	for i, v := range actual {
		if !reflect.DeepEqual(Data[i], v) {
			t.Errorf("\nExpected: %v\nActual: %v\n", Data[i], v)
		}
	}
}

func TestMoveRobot(t *testing.T) {
	tests := []struct {
		Robot                   Robot
		EPX, EPY, Width, Height int
	}{
		{
			Robot:  Robot{2, 4, 2, -3},
			EPX:    4,
			EPY:    1,
			Width:  11,
			Height: 7,
		},
		{
			Robot:  Robot{4, 1, 2, -3},
			EPX:    6,
			EPY:    5,
			Width:  11,
			Height: 7,
		},
		{
			Robot:  Robot{6, 5, 2, -3},
			EPX:    8,
			EPY:    2,
			Width:  11,
			Height: 7,
		},
		{
			Robot:  Robot{8, 2, 2, -3},
			EPX:    10,
			EPY:    6,
			Width:  11,
			Height: 7,
		},
		{
			Robot:  Robot{10, 6, 2, -3},
			EPX:    1,
			EPY:    3,
			Width:  11,
			Height: 7,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.Robot), func(t *testing.T) {
			test.Robot.MoveRobot(test.Width, test.Height)
			if test.Robot.PX != test.EPX || test.Robot.PY != test.EPY {
				t.Errorf("Expected X: %d, Actual X: %d\nExpected Y: %d, Actual Y: %d", test.EPX, test.Robot.PX, test.EPY, test.Robot.PY)
			}
		})
	}
}

func TestMoveRobots(t *testing.T) {
	data := make([]Robot, len(Data))
	copy(data, Data)
	PrintRobots(data, 11, 7)
	for range 100 {
		MoveRobots(data, 11, 7)
	}
	PrintRobots(data, 11, 7)
}

func TestFindMiddle(t *testing.T) {
	tests := []struct {
		Input    int
		Expected []int
	}{
		{
			Input:    7,
			Expected: []int{3},
		},
		{
			Input:    11,
			Expected: []int{5},
		},
		{
			Input:    10,
			Expected: []int{4, 5},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.Input), func(t *testing.T) {
			actual := FindMiddle(test.Input)
			if !reflect.DeepEqual(test.Expected, actual) {
				t.Errorf("Expected: %v, Actual: %v\n", test.Expected, actual)
			}
		})
	}
}
