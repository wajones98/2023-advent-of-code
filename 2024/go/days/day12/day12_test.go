package day12

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/wajones98/advent-of-code/common"
)

const Input = `AAAA
BBCD
BBCC
EEEC`

var Data = common.TwoDMap[string]{
	Width:  4,
	Height: 4,
	Map:    []string{"A", "A", "A", "A", "B", "B", "C", "D", "B", "B", "C", "C", "E", "E", "E", "C"},
}

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(Data, *actual) {
		t.Errorf("Expected: %v, Actual: %v\n", Data, actual)
	}
}
