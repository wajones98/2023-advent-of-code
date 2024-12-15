package day14

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/wajones98/advent-of-code/common"
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

var Data = common.TwoDMap[[]Robot]{
	Width:  11,
	Height: 7,
	Map: [][]Robot{
		{{1, 3}}, {}, {{2, -1}}, {{-2, -2}, {-1, -2}}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {{-1, -3}}, {{-1, 2}}, {}, {{2, 3}}, {{-1, 2}},
		{{3, -3}}, {}, {{2, -3}}, {}, {}, {}, {}, {}, {}, {}, {},
		{}, {}, {}, {}, {}, {}, {}, {}, {}, {{-3, -3}}, {},
		{}, {}, {}, {}, {}, {}, {}, {{-1, -3}}, {}, {}, {},
	},
}

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, err := LoadInput(s, 11, 7)
	if err != nil {
		t.Error(err)
	}

	if Data.Width != actual.Width {
		t.Errorf("\nExpected: %v\nActual: %v\n", Data.Width, actual.Width)
	}

	if Data.Height != actual.Height {
		t.Errorf("\nExpected: %v\nActual: %v\n", Data.Height, actual.Height)
	}

	if len(Data.Map) != len(actual.Map) {
		t.Errorf("\nExpected: %v\nActual: %v\n", len(Data.Map), len(actual.Map))
	}

	for i, v := range actual.Map {
		if len(v) != len(Data.Map[i]) && !reflect.DeepEqual(Data.Map[i], v) {
			t.Errorf("\nExpected: %v\nActual: %v\n", Data.Map[i], v)
		}
	}
}

func TestMoveRobot(t *testing.T) {
	tests := []struct {
		Robot                           Robot
		PX, EPX, PY, EPY, Width, Height int
	}{
		{
			Robot:  Robot{2, -3},
			PX:     2,
			PY:     4,
			EPX:    4,
			EPY:    1,
			Width:  11,
			Height: 7,
		},
		{
			Robot:  Robot{2, -3},
			PX:     4,
			PY:     1,
			EPX:    6,
			EPY:    5,
			Width:  11,
			Height: 7,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v - Starting -> X: %d, Y: %d", test.Robot, test.PX, test.PY), func(t *testing.T) {
			apx, apy := MoveRobot(test.Robot, test.PX, test.PY, test.Width, test.Height)
			if apx != test.EPX || apy != test.EPY {
				t.Errorf("Expected X: %d, Actual X: %d\nExpected Y: %d, Actual Y: %d", test.EPX, apx, test.EPY, apy)
			}
		})
	}
}
