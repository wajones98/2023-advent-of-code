package day14

import (
	"bufio"
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
