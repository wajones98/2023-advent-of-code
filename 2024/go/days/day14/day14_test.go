package day14

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

const Input = ""
const Data = ""

func TestLoadInput(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	actual, err := LoadInput(s)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(Data, *actual) {
		t.Errorf("Expected: %v, Actual: %v\\n", Data, actual)
	}
}
