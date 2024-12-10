package day9

import (
	"bufio"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

const Input = `2333133121414131402`

func TestLoadInput(t *testing.T) {
	input := "12345"
	s := bufio.NewScanner(strings.NewReader(input))

	actual := LoadInput(s)
	expected := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Got: %v\n", expected, actual)
	}
}

func TestLoadInputPartTwo(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))

	actual := LoadInputPartTwo(s)
	expected := []Block{{0, 2}, {-1, 3}, {1, 3}, {-1, 3}, {2, 1}, {-1, 3}, {3, 3}, {-1, 1}, {4, 2}, {-1, 1}, {5, 4}, {-1, 1}, {6, 4}, {-1, 1}, {7, 3}, {-1, 1}, {8, 4}, {9, 2}}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nExpected: %v\nGot:      %v\n", expected, actual)
	}
}

func TestCompress(t *testing.T) {
	data := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	expected := []int{0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1}
	expectedChecksum := 60
	checksum := Compress(data)

	if !reflect.DeepEqual(expected, data) {
		t.Errorf("Expected: %v, Got: %v\n", expected, data)
	}

	if expectedChecksum != checksum {
		t.Errorf("Expected: %d, Got: %d\n", expectedChecksum, checksum)
	}

}

func TestExampleInputPartOne(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	data := LoadInput(s)
	expected := 1928
	checksum := Compress(data)

	if expected != checksum {
		t.Errorf("Expected: %d, Got: %d\n", expected, checksum)
	}
}

func TestExampleInputPartTwo(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(Input))
	data := LoadInput(s)
	expectedData := []int{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1}

	// expectedChecksum := 2858
	CompressPartTwo(data)

	// if expectedChecksum != checksum {
	// 	t.Errorf("Expected: %d, Got: %d\n", expected, checksum)
	// }

	if len(expectedData) != len(data) {
		t.Errorf("Expected: %d, Got: %d\n", len(expectedData), len(data))

	}

	for i, e := range expectedData {
		if !reflect.DeepEqual(e, data[i]) {
			t.Errorf("Expected: %v, Got: %v\n", e, data[i])
		}
	}

	fmt.Printf("\nExpected: %v\nGot:      %v\n", expectedData, data)
}
